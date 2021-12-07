package domain

import (
	"context"
	"time"
)

type (
	ProjectID string
	OwnerType string

	Project struct {
		ID        ProjectID `json:"id"`
		Slug      string    `json:"slug"`
		Name      string    `json:"name"`
		OwnerType OwnerType `json:"owner_type"`
		IsValid   bool      `json:"is_valid"`
		IsDeleted bool      `json:"is_deleted"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		// fk

		OwnerUser *User `json:"user"`
		OwnerOrg  *Org  `json:"org"`
	}

	ProjectInput struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
		Org  *Org   `json:"org"`
	}

	IProjectInteractor interface {
		List(context.Context) ([]Project, error)
		Detail(context.Context, string) (*Project, error)
		Create(context.Context, ProjectInput) (string, error)
	}
)

const (
	// if owner is user
	OwnerTypeUser = OwnerType("user")
	// if owner is org
	OwnerTypeOrg = OwnerType("org")
)
