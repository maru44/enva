package domain

import "context"

type (
	Settings struct {
		ProjectSlug string  `json:"project_slug"`
		EnvFileName string  `json:"env_file_name"`
		OrgSlug     *string `json:"org_slug,omitempty"`
	}

	ICommandInteractor interface {
		Run(context.Context, ...string) error
	}
)
