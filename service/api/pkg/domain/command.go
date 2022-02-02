package domain

import "context"

type (
	Settings struct {
		ProjectSlug string  `json:"project"`
		EnvFileName string  `json:"file"`
		OrgSlug     *string `json:"org,omitempty"`
		PreSentence *string `json:"pre,omitempty"`
		SufSentence *string `json:"suf,omitempty"`
	}

	ICommandInteractor interface {
		Run(context.Context, ...string) error
		Explain() string
	}
)
