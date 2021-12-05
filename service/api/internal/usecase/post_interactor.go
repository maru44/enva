package usecase

import (
	"context"

	"github.com/maru44/ichigo/service/api/pkg/domain"
)

type PostInteractor struct {
	post PostRepositoryAbstract
}

// init post interactor
func NewPostInteractor(post PostRepositoryAbstract) *PostInteractor {
	return &PostInteractor{
		post: post,
	}
}

/************************
    repository
*************************/

type PostRepositoryAbstract interface {
	List(context.Context) ([]domain.Post, error)
	Create(context.Context, domain.PostInput) (string, error)
}

/************************
    post interactor methods
*************************/

// list posts
func (in *PostInteractor) List(ctx context.Context) ([]domain.Post, error) {
	return in.post.List(ctx)
}

// create post
func (in *PostInteractor) Create(ctx context.Context, input domain.PostInput) (string, error) {
	return in.post.Create(ctx, input)
}
