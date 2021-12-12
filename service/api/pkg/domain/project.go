package domain

import (
	"context"
	"errors"
	"time"
)

type (
	ProjectID string
	OwnerType string

	Project struct {
		ID          ProjectID `json:"id"`
		Slug        string    `json:"slug"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		OwnerType   OwnerType `json:"owner_type"`
		IsValid     bool      `json:"is_valid"`
		IsDeleted   bool      `json:"is_deleted"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`

		// fk

		OwnerUser *User `json:"user"`
		OwnerOrg  *Org  `json:"org"`
	}

	ProjectInput struct {
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		OrgID       *OrgID `json:"org_id"`
	}

	IProjectInteractor interface {
		ListByUser(context.Context) ([]Project, error)
		ListByOrg(context.Context, OrgID) ([]Project, error)
		Detail(context.Context, ProjectID) (*Project, error)
		Create(context.Context, ProjectInput) (*ProjectID, error)
	}
)

const (
	// if owner is user
	OwnerTypeUser = OwnerType("user")
	// if owner is org
	OwnerTypeOrg = OwnerType("org")
)

var (
	ErrProjectSlugAlreadyExistsUser = errors.New("Slug duplicated: Project slug has already exists for user")         // 400
	ErrProjectSlugAlreadyExistsOrg  = errors.New("Slug duplicated: Project slug has already exists for organization") // 400
)
