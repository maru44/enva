package domain

import (
	"context"
	"time"
)

type (
	/***********************
	   struct
	***********************/

	Post struct {
		ID        string     `json:"id"`
		Title     string     `json:"title"`
		Abstract  string     `json:"abstract"`
		Content   string     `json:"content"`
		IsExpired bool       `json:"is_expired"`
		IsPublic  bool       `json:"is_public"`
		IsValid   bool       `json:"is_valid"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		ExpiredAt *time.Time `json:"expired_at"`

		// foreign key

		User     *User     `json:"user"`
		Category *Category `json:"category"`
		Tags     []*Tag    `json:"tags"`
	}

	PostInput struct {
		Title    string `json:"title"`
		Abstract string `json:"abstract"`
		Content  string `json:"content"`
	}

	Tag struct {
		ID        string     `json:"id"`
		Slug      string     `json:"slug"`
		Name      string     `json:"name"`
		IsActive  bool       `json:"is_active"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	Category struct {
		ID        string     `json:"id"`
		Slug      string     `json:"slug"`
		Name      string     `json:"name"`
		IsActive  bool       `json:"is_active"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	/***********************
	   interface
	***********************/

	PostInteractorAbstract interface {
		List(context.Context) ([]Post, error)
		Create(context.Context, PostInput) (string, error)
	}
)
