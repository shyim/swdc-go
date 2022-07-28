package cmd

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

func getDockerProject() *types.Project {
	project := &types.Project{
		Name: "shopware-docker",
		Configs: map[string]types.ConfigObjConfig{
			"foo": types.ConfigObjConfig{
				Name: "bla",
			},
		},
		Services: []types.ServiceConfig{
			{
				Name:  "shit",
				Image: "nginx",
			},
		},
	}

	fixServices(project)

	return project
}

func fixServices(project *types.Project) {
	for key, service := range project.Services {
		service.CustomLabels = map[string]string{
			api.ProjectLabel:     "shopware-docker",
			api.ServiceLabel:     "shit",
			api.VersionLabel:     api.ComposeVersion,
			api.WorkingDirLabel:  "",
			api.ConfigFilesLabel: "",
			api.OneoffLabel:      "False", // default, will be overridden by `run` command
		}

		project.Services[key] = service
	}
}