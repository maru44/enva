package domain

import (
	"context"
	"errors"
	"time"

	"github.com/maru44/perr"
)

type (
	ProjectID string
	OwnerType string

	Project struct {
		ID          ProjectID `json:"id"`
		Slug        string    `json:"slug"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
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
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		OrgID       *OrgID  `json:"org_id"`
	}

	IProjectInteractor interface {
		ListByUser(context.Context) ([]Project, error)
		ListByOrg(context.Context, OrgID) ([]Project, error)
		SlugListByUser(context.Context) ([]string, error)
		// SlugListByOrg(context.Context, OrgID) ([]string, error)

		GetBySlug(context.Context, string) (*Project, error)
		GetByID(context.Context, ProjectID) (*Project, error)
		Create(context.Context, ProjectInput) (*string, error)
		Delete(context.Context, ProjectID) (int, error)

		// by org 系
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

func (p *Project) ValidateUserGet(u User) error {
	// user type
	if p.OwnerOrg == nil {
		if p.OwnerUser.ID == u.ID {
			return nil
		}
		return perr.New("Not owner of project", perr.Forbidden)
	}

	// org type
	if p.OwnerOrg.IsMember(u) {
		return nil
	}
	return perr.New("Not member of owner organization", perr.Forbidden)
}

// func (p *Project) ValidateUserPost()
