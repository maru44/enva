package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	OrgID string

	Org struct {
		ID          OrgID     `json:"id"`
		Slug        string    `json:"slug"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
		IsValid     bool      `json:"is_valid"`
		CreatedBy   User      `json:"created_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`

		UserCount int `json:"user_count"`
	}

	OrgInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	IOrgInteractor interface {
		List(context.Context) ([]Org, error)
		Detail(context.Context, OrgID) (*Org, error)
		DetailBySlug(context.Context, string) (*Org, error)
		Create(context.Context, OrgInput) (*string, error)
		// Update
		// Delete
	}
)

func (o *OrgInput) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(o.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(o.Name, validation.Required, validation.Length(1, 64)),
	)
}

var (
// ValidationErrorOrgMemberInput = perr.New("")
)

// func (o *Org) IsMember(u *User) bool {
// 	if u == nil {
// 		return false
// 	}
// 	for _, user := range o.Users {
// 		if user.ID == u.ID {
// 			return true
// 		}
// 	}

// 	return o.IsAdmin(u)
// }

// func (o *Org) IsAdmin(u *User) bool {
// 	if u == nil {
// 		return false
// 	}
// 	for _, user := range o.Admins {
// 		if user.ID == u.ID {
// 			return true
// 		}
// 	}

// 	return false
// }
