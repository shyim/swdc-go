package cmd

import (
	"github.com/docker/compose/v2/pkg/api"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stops all containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		backend := createDockerBackend()

		project := getDockerProject()

		downOptions := api.DownOptions{
			Project:       project,
			RemoveOrphans: true,
			Volumes:       true,
		}

		return backend.Down(cmd.Context(), project.Name, downOptions)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
