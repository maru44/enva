package database

import (
	"context"

	"github.com/maru44/ichigo/service/api/internal/interface/database/queryset"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type PostRepository struct {
	SqlHandlerAbstract
}

func (repo *PostRepository) List(ctx context.Context) (posts []domain.Post, err error) {
	rows, err := repo.QueryContext(ctx, queryset.PostListQuery)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if rows.Err() != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	for rows.Next() {
		p := domain.Post{
			User: &domain.User{},
		}
		err = rows.Scan(
			&p.ID, &p.Title, &p.Abstract,
			&p.Content, &p.IsValid, &p.IsPublic,
			&p.CreatedAt, &p.UpdatedAt,
			&p.User.ID,
		)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (repo *PostRepository) Create(ctx context.Context, input domain.PostInput) (id string, err error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)
	if err := repo.QueryRowContext(
		ctx,
		queryset.PostInsertQuery,
		input.Title, input.Abstract, input.Content,
		true, true, user.ID, // @TODO fix true
	).Scan(&id); err != nil {
		return id, perr.Wrap(err, perr.BadRequest)
	}

	return id, nil
}
