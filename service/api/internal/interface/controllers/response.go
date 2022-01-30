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
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/perr"
)

func response(w http.ResponseWriter, r *http.Request, err error, body map[string]interface{}) {
	status := getStatusCode(err, w)
	ctx := r.Context()

	if status == http.StatusOK {
		data, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			mess, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
			if _, err := w.Write(mess); err != nil {
				sendSentryErr(ctx, err)
			}
			return
		}
		w.WriteHeader(status)
		if _, err := w.Write(data); err != nil {
			sendSentryErr(ctx, err)
		}
		return
	}

	w.WriteHeader(status)
	mess := map[string]interface{}{
		"status": status,
	}
	if perror, ok := perr.IsPerror(err); ok {
		mess["error"] = perror.Output().Error()
		sendSentryPerror(ctx, perror)
	} else {
		mess["error"] = err.Error()
		sendSentryErr(ctx, err)
	}

	data, _ := json.Marshal(mess)
	if _, err := w.Write(data); err != nil {
		sendSentryErr(ctx, err)
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

func sendSentryPerror(ctx context.Context, err perr.Perror) {
	if config.IsEnvDevelopment {
		log.Println(err)
		fmt.Println("stack traces:\n", err.Traces())
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		panic(err)
	}

	var (
		sLevel = sentry.LevelWarning
		pm     = err.Map()
		cat    = "Error"
	)
	switch err.Level() {
	case perr.ErrLevelAlert:
		sLevel = sentry.LevelWarning
	case perr.ErrLevelInternal:
		sLevel = sentry.LevelError
	case perr.ErrLevelExternal:
		sLevel = sentry.LevelInfo
	default:
		panic("must not reach here")
	}

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(map[string]string{
			"category": cat,
		})
		scope.SetTag("level", string(err.Level()))
	})
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category:  cat,
		Level:     sLevel,
		Message:   err.Unwrap().Error(),
		Timestamp: pm.OccurredAt,
		Data: map[string]interface{}{
			"treated_as": err.Output().Error(),
		},
	})

	defer sentry.Flush(3 * time.Second)
	sentry.CaptureMessage(err.Unwrap().Error())
}

func sendSentryErr(ctx context.Context, err error) {
	if config.IsEnvDevelopment {
		log.Println(err)
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		panic(err)
	}

	var (
		sLevel = sentry.LevelWarning
		cat    = "Unexpected Error"
	)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(map[string]string{
			"category": cat,
		})
	})
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category:  cat,
		Level:     sLevel,
		Message:   err.Error(),
		Timestamp: time.Now(),
	})

	defer sentry.Flush(3 * time.Second)
	sentry.CaptureMessage(err.Error())
}
