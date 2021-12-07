package domain

import "context"

type (
	KVKey   string
	KVValue string

	KV struct {
		CommitID CommitID `json:"commit_id"`
		Key      KVKey    `json:"kv_key"`
		Value    KVValue  `json:"kv_value"`
	}

	KVInput struct {
		Key   KVKey   `json:"kv_key"`
		Value KVValue `json:"kv_value"`
	}

	IKVInteractor interface {
		List(context.Context, CommitID) ([]KV, error)
		Create(context.Context, KVInput, CommitID) (int, error)
	}
)
