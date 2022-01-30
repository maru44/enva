package controllers

import (
	"net/http"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/perr"
)

type HealthController struct {
	sql database.ISqlHandler
}

func NewHealthController(sql database.ISqlHandler) *HealthController {
	return &HealthController{
		sql: sql,
	}
}

func (con *HealthController) PostgresCheckView(w http.ResponseWriter, r *http.Request) {
	row := con.sql.QueryRowContext(r.Context(), "SELECT 1 FROM DUAL")
	if err := row.Err(); err != nil {
		response(w, r, perr.Wrap(err, perr.InternalServerErrorWithUrgency), nil)
		return
	}
	var one int
	if err := row.Scan(&one); err != nil {
		response(w, r, perr.Wrap(err, perr.InternalServerErrorWithUrgency), nil)
		return
	}
	response(w, r, nil, map[string]interface{}{"data": one})
}
