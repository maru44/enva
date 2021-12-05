package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/ichigo/service/api/internal/interface/database"
	"github.com/maru44/ichigo/service/api/internal/usecase"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type PostController struct {
	in usecase.PostInteractor
}

func NewPostController(sql database.SqlHandlerAbstract) *PostController {
	return &PostController{
		in: *usecase.NewPostInteractor(
			&database.PostRepository{
				SqlHandlerAbstract: sql,
			},
		),
	}
}

func (con *PostController) ListView(w http.ResponseWriter, r *http.Request) {
	posts, err := con.in.List(r.Context())
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": posts})
	return
}

func (con *PostController) CreateView(w http.ResponseWriter, r *http.Request) {
	var input domain.PostInput
	json.NewDecoder(r.Body).Decode(&input)
	id, err := con.in.Create(r.Context(), input)

	response(w, r, perr.Wrap(err, perr.BadRequest), map[string]interface{}{"data": id})
	return
}
