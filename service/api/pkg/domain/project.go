package domain

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	ProjectID string
	OwnerType string

	Project struct {
		ID          ProjectID  `json:"id"`
		Slug        string     `json:"slug"`
		Name        string     `json:"name"`
		Description *string    `json:"description"`
		OwnerType   OwnerType  `json:"owner_type"`
		IsValid     bool       `json:"is_valid"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`

		// fk

		OwnerUser *User `json:"user"`
		OwnerOrg  *Org  `json:"org"`
	}

	ProjectInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		OrgID       *OrgID  `json:"org_id,omitempty"`
	}

	CliProjectInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		OrgSlug     *string `json:"org_slug"`
		Description *string `json:"description"`
	}

	IProjectInteractor interface {
		ListAll(context.Context) ([]Project, error)
		ListByUser(context.Context) ([]Project, error)
		ListByOrg(context.Context, OrgID) ([]Project, error)
		SlugListByUser(context.Context) ([]string, error)
		// SlugListByOrg(context.Context, OrgID) ([]string, error)

		GetBySlug(context.Context, string) (*Project, error)
		GetBySlugAndOrgID(context.Context, string, OrgID) (*Project, error)
		GetBySlugAndOrgSlug(context.Context, string, string) (*Project, error)
		GetByID(context.Context, ProjectID) (*Project, error)
		Create(context.Context, ProjectInput) (*string, error)
		Delete(context.Context, ProjectID) (int, error)

		CountValidByOrgID(context.Context, OrgID) (*int, *Subscription, error)
		CountValidByOrgSlug(context.Context, string) (*int, *Subscription, error)
		CountValidByUser(context.Context, UserID) (*int, *Subscription, error)
	}
)

func (p *ProjectInput) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Slug, validation.Required, validation.Length(1, 64), validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-_]+$`))),
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
	)
}

func (p *CliProjectInput) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Slug, validation.Required, validation.Length(1, 64), validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-_]+$`))),
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
	)
}

func (p *CliProjectInput) ToProjectInputWithoutOrg() ProjectInput {
	return ProjectInput{
		Slug:        p.Slug,
		Name:        p.Name,
		Description: p.Description,
	}
}

const (
	// if owner is user
	OwnerTypeUser = OwnerType("user")
	// if owner is org
	OwnerTypeOrg = OwnerType("org")
)

var (
	ErrProjectSlugAlreadyExistsUser = errors.New("Slug duplicated: Project slug has already exists for user")         // 400
	ErrProjectSlugAlreadyExistsOrg  = errors.New("Slug duplicated: Project slug has already exists for organization") // 400

	ErrProjectIsNotValid = errors.New("Project is not valid")
)

func (p *Project) Valid() error {
	if !p.IsValid || p.DeletedAt != nil {
		return ErrProjectIsNotValid
	}
	return nil
}

func CanCreateProject(s *Subscription, projectCount int, ownerType OwnerType) error {
	errStr := "Projects count is reached to max value(%d)"
	switch ownerType {
	case OwnerTypeOrg:
		if s == nil && projectCount > 4 {
			return errors.New(fmt.Sprintf(errStr, 5))
		}
		// if s.StripeProductID == a {}
		// if s.StripeProductID == b {}
	default:
		if s == nil && projectCount > 4 {
			return errors.New(fmt.Sprintf(errStr, 5))
		}
		// if s.StripeProductID == a {}
		// if s.StripeProductID == b {}
	}

	return nil
}
