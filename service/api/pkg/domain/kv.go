package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/maru44/perr"
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
		Key   KvKey   `json:"kv_key"`
		Value KvValue `json:"kv_value"`
	}

	KvInput struct {
		Key   KvKey   `json:"kv_key"`
		Value KvValue `json:"kv_value"`
	}

	KvInputWithProjectID struct {
		ProjectID ProjectID `json:"project_id"`
		Input     KvInput   `json:"input"`
	}

	IKvInteractor interface {
		ListValid(context.Context, ProjectID) ([]Kv, error)
		DetailValid(context.Context, KvKey, ProjectID) (*Kv, error)
		Create(context.Context, KvInputWithProjectID) (*KvID, error)
		Update(context.Context, KvInputWithProjectID) (*KvID, error)
		Delete(context.Context, KvID, ProjectID) (int, error) // @TODO fix arg type
	}
)

func (k *KvKey) String() string {
	return string(*k)
}

func (v *KvValue) String() string {
	return string(*v)
}

func (in *KvInput) Validate() error {
	if err := validation.Validate(in.Key, validation.Required, validation.RuneLength(1, 1024)); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}
