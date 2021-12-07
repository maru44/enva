package domain

import "context"

type (
	RelOrgMembersInput struct {
		OrgID    OrgID     `json:"org_id"`
		UserId   string    `json:"user_id"`
		UserType *UserType `json:"user_type"`
	}

	IRelOrgMembersInteractor interface {
		Add(context.Context, RelOrgMembersInput) (int, error)
	}
)
