package cmd

import (
	"github.com/docker/compose/v2/pkg/api"
	"github.com/sanathkr/go-yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		backend := createDockerBackend()

		project := getDockerProject()

		config, _ := yaml.Marshal(project)

		ioutil.WriteFile("docker-compose.yml", config, os.ModePerm)

		upOptions := api.UpOptions{
			Create: api.CreateOptions{
				Services:             make([]string, 0),
				RemoveOrphans:        true,
				Recreate:             api.RecreateDiverged,
				RecreateDependencies: api.RecreateDiverged,
			},
			Start: api.StartOptions{
				Project: project,
				Attach:  nil,
			},
		}

		return backend.Up(cmd.Context(), project, upOptions)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
