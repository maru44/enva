package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/perr"
)

func response(w http.ResponseWriter, r *http.Request, err error, body map[string]interface{}) {
	status := getStatusCode(err, w)

	if status == http.StatusOK {
		data, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mess, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
			if _, err := w.Write(mess); err != nil {
				log.Fatal(err)
			}
			return
		}
		w.WriteHeader(status)
		if _, err := w.Write(data); err != nil {
			log.Fatal(err)
		}
		return
	}

	w.WriteHeader(status)
	var mess map[string]interface{}
	if perror, ok := perr.IsPerror(err); ok {
		mess = map[string]interface{}{
			"error":  perror.Output().Error(),
			"status": status,
		}
	} else {
		mess = map[string]interface{}{
			"error":  err.Error(),
			"status": status,
		}
	}

	if config.IsEnvDevelopment {
		log.Println(err)

		if perror, ok := perr.IsPerror(err); ok {
			fmt.Println("stack traces:\n", perror.Traces())
		}
	}

	data, _ := json.Marshal(mess)
	if _, err := w.Write(data); err != nil {
		log.Fatal(err)
	}
	return
}

func getStatusCode(err error, w http.ResponseWriter) int {
	if err == nil || reflect.ValueOf(err).IsNil() {
		return http.StatusOK
	}

	if perror, ok := perr.IsPerror(err); ok {
		switch {
		case perror.IsOutput(perr.InternalServerError), perror.IsOutput(perr.InternalServerErrorWithUrgency):
			return http.StatusInternalServerError
		case perror.IsOutput(perr.NotFound):
			return http.StatusNotFound
		case perror.IsOutput(perr.Forbidden):
			return http.StatusForbidden
		case perror.IsOutput(perr.Unauthorized), perror.IsOutput(perr.Expired), perror.IsOutput(perr.InvalidToken):
			return http.StatusUnauthorized
		case perror.IsOutput(perr.BadRequest):
			return http.StatusBadRequest
		case perror.IsOutput(perr.Created):
			return http.StatusCreated
		case perror.IsOutput(perr.UnsupportedMediaType):
			return http.StatusUnsupportedMediaType
		case perror.IsOutput(perr.MethodNotAllowed):
			return http.StatusMethodNotAllowed
		case perror.IsOutput(perr.UnsupportedMediaType):
			return http.StatusUnsupportedMediaType
		default:
			return http.StatusInternalServerError
		}
	}

	switch err {
	case perr.InternalServerError, perr.InternalServerErrorWithUrgency:
		return http.StatusInternalServerError
	case perr.NotFound:
		return http.StatusNotFound
	case perr.Forbidden:
		return http.StatusForbidden
	case perr.Unauthorized, perr.Expired, perr.InvalidToken:
		return http.StatusUnauthorized
	case perr.BadRequest:
		return http.StatusBadRequest
	case perr.Created:
		return http.StatusCreated
	case perr.UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	case perr.MethodNotAllowed:
		return http.StatusMethodNotAllowed
	case perr.UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}
