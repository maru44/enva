package domain

import (
	"context"
	"time"
)

type (
	CommitID string

	Commit struct {
		ID           CommitID  `json:"id"`
		CommitNumber int       `json:"commit_number"`
		IsHead       bool      `json:"is_head"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`

		Project Project `json:"project"`
	}

	CommitInput struct {
		ProjectID ProjectID `json:"project_id"`
		KVs       []KV      `json:"kvs"`
	}

	ICommitInteractor interface {
		List(context.Context, ProjectID) ([]Commit, error)
		Commit(context.Context, CommitID) (*Commit, error)
		Create(context.Context)
	}
)
