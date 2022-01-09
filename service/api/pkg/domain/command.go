package domain

import "context"

type (
	Settings struct {
		ProjectSlug string  `json:"project_slug"`
		EnvFileName string  `json:"env_file_name"`
		OrgSlug     *string `json:"org_slug,omitempty"`
		PreSentence *string `json:"pre_sentence,omitempty"`
		SufSentence *string `json:"suf_sentence,omitempty"`
	}

	ICommandInteractor interface {
		Run(context.Context, ...string) error
		Explain() string
	}
)
