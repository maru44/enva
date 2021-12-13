package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/maru44/ichigo/service/api/internal/interface/database"
	"github.com/maru44/ichigo/service/api/internal/usecase"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type KvController struct {
	in domain.IKvInteractor
}

func NewKvController(sql database.ISqlHandler) *KvController {
	return &KvController{
		in: usecase.NewKvInteractor(
			&database.KvRepository{
				ISqlHandler: sql,
			},
		),
	}
}

func (con *KvController) ListView(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get(QueryParamsProjectID)
	kvs, err := con.in.ListValid(r.Context(), domain.ProjectID(projectID))
	if err != nil {
		response(w, r, perr.Wrap(err, perr.NotFound), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": kvs})
	return
}

func (con *KvController) CreateView(w http.ResponseWriter, r *http.Request) {
	var input domain.KvInputWithProjectID
	json.NewDecoder(r.Body).Decode(&input)

	key, value, err := con.in.Create(r.Context(), input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
		return
	}

	response(w, r, nil,
		map[string]interface{}{
			"data": map[string]string{
				"env_key":   *key,
				"env_value": *value,
			},
		})
	return
}

func (con *KvController) UpdateView(w http.ResponseWriter, r *http.Request) {
	var input domain.KvInputWithProjectID
	json.NewDecoder(r.Body).Decode(&input)

	id, err := con.in.Update(r.Context(), input)
	if err != nil {
		response(w, r, perr.Wrap(err, perr.BadRequest), nil)
	}

	response(w, r, nil, map[string]interface{}{"data": id})
	return
}
