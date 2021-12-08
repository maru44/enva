package domain

import (
	"context"
	"time"
)

type (
	KvID    string
	KvKey   string
	KvValue string

	// @TODO add craeted_by, updated_by
	Kv struct {
		ID        KvID      `json:"id"`
		ProjectID ProjectID `json:"project_id"`
		Key       KvKey     `json:"kv_key"`
		Value     KvValue   `json:"kv_value"`
		IsValid   bool      `json:"is_valid"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	KvValid struct {
		Key   KvKey   `json:"key"`
		Value KvValue `json:"value"`
	}

	KvInput struct {
		Key   KvKey   `json:"kv_key"`
		Value KvValue `json:"kv_value"`
	}

	IKvInteractor interface {
		ListValid(context.Context, ProjectID) ([]Kv, error)
		DetailValid(context.Context, KvKey, ProjectID) (*Kv, error)
		Create(context.Context, KvInput, ProjectID) (*KvID, error)
		Update(context.Context, KvInput, ProjectID) (*KvID, error)
	}
)
