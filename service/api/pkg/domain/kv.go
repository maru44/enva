package domain

import (
	"context"
	"regexp"
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
		ProjectID ProjectID `json:"project_id,omitempty"`
		Input     KvInput   `json:"input"`
	}

	IKvInteractor interface {
		ListValid(context.Context, ProjectID) ([]Kv, error)
		DetailValid(context.Context, KvKey, ProjectID) (*Kv, error)
		Create(context.Context, KvInputWithProjectID) (*KvID, error)
		Update(context.Context, KvInputWithProjectID) (*KvID, error)
		Delete(context.Context, KvID, ProjectID) (int, error)
		DeleteByKey(context.Context, KvKey, ProjectID) (int, error)
	}

	ICliKvInteractor interface {
		BulkInsert(context.Context, ProjectID, []KvInput) error
	}
)

func (k *KvKey) String() string {
	return string(*k)
}

func (v *KvValue) String() string {
	return string(*v)
}

func (kv *KvValid) ToInput() *KvInput {
	return &KvInput{
		Key:   kv.Key,
		Value: kv.Value,
	}
}

func (in *KvInput) Validate() error {
	if err := validation.Validate(in.Key,
		validation.Required,
		validation.Length(1, 1024),
		validation.Match(regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func MapFromKv(ks []KvValid) map[KvKey]KvValue {
	ret := make(map[KvKey]KvValue, len(ks))
	for _, k := range ks {
		ret[k.Key] = k.Value
	}
	return ret
}
