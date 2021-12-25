package domain

import (
	"context"
	"time"
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

		// fk

		Admins []User `json:"admin"`
		Users  []User `json:"users"`
		Guests []User `json:"guests"`
	}

	OrgInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	IOrgInteractor interface {
		// List(ctx context.Context)
		Detail(context.Context, string) (*Org, error)
		Create(context.Context)
	}
)

func (o *Org) IsMember(u *User) bool {
	if u == nil {
		return false
	}
	if o.Users != nil {
		for _, user := range o.Users {
			if user.ID == u.ID {
				return true
			}
		}
	}

	return o.IsAdmin(u)
}

func (o *Org) IsAdmin(u *User) bool {
	if u == nil {
		return false
	}
	if o.Admins != nil {
		for _, user := range o.Admins {
			if user.ID == u.ID {
				return true
			}
		}
	}

	return false
}
