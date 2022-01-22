package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/getsentry/sentry-go"
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

		if config.IsEnvDevelopment {
			log.Println(err)
			fmt.Println("stack traces:\n", perror.Traces())
		}
	} else {
		mess = map[string]interface{}{
			"error":  err.Error(),
			"status": status,
		}

		if config.IsEnvDevelopment {
			log.Println(err)
		}
	}

	// only production env
	if !config.IsEnvDevelopment {
		sendSentry(r.Context(), err)
	}

	data, _ := json.Marshal(mess)
	if _, err := w.Write(data); err != nil {
		// only production env
		if !config.IsEnvDevelopment {
			sendSentry(r.Context(), err)
		}
	}
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

func sendSentry(ctx context.Context, err error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		panic(err)
	}

	var (
		message   string
		data      = map[string]interface{}{}
		sLevel    = sentry.LevelWarning
		cat       = "Unexpected Error"
		timeStamp = time.Now()
	)
	if perror, ok := perr.IsPerror(err); ok {
		switch perror.Level() {
		case perr.ErrLevelAlert:
			sLevel = sentry.LevelWarning
		case perr.ErrLevelInternal:
			sLevel = sentry.LevelError
		case perr.ErrLevelExternal:
			sLevel = sentry.LevelInfo
		default:
			panic("must not reach here")
		}

		cat = "Error"
		pm := perror.Map()

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTags(map[string]string{
				"category": cat,
			})
			scope.SetTag("level", string(perror.Level()))
		})
		message = perror.Unwrap().Error()
		timeStamp = pm.OccurredAt
		data["treated_as"] = perror.Output().Error()
	} else {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTags(map[string]string{
				"category": cat,
			})
		})
	}

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category:  cat,
		Level:     sLevel,
		Message:   message,
		Timestamp: timeStamp,
		Data:      data,
	})

	defer sentry.Flush(3 * time.Second)
}
