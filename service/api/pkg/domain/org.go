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

		Users []User `json:"users"`
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
