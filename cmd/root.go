package cmd

import (
	"context"
	"fmt"
	"github.com/Shopify/go-lua"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swdc",
	Short: "Shopware Docker",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	l.Register("dockerCfg", func(state *lua.State) int {
		state.PushUserData(*getDockerProject())

		return 1
	})

	l.Register("registerCommand", func(state *lua.State) int {
		funcName := lua.CheckString(l, 1)

		l.SetGlobal(funcName)

		rootCmd.AddCommand(&cobra.Command{
			Use:   funcName,
			Short: lua.CheckString(l, 2),
			RunE: func(cmd *cobra.Command, args []string) error {
				l.Global(funcName)

				for _, arg := range args {
					l.PushString(arg)
				}

				l.Call(len(args), 1)

				if l.ToBoolean(1) {
					return nil
				}

				return fmt.Errorf("command failed")
			},
		})

		return 0
	})

	fmt.Println(lua.DoFile(l, "main.lua"))

}

func createDockerBackend() *api.ServiceProxy {
	cli, err := command.NewDockerCli()

	if err != nil {
		log.Fatalln(err)
	}

	if err := cli.Initialize(flags.NewClientOptions()); err != nil {
		log.Fatalln(err)
	}

	backend := api.NewServiceProxy()
	backend.WithService(compose.NewComposeService(cli))

	return backend
}
