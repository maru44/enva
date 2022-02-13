package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type (
	jwtInteractorForTest struct {
		usecase.JwtInteractor
	}

	testContextViewBody struct {
		Access domain.CtxAccess `json:"access"`
	}
)

func newBaseControllerForTest(t *testing.T) *BaseController {
	return &BaseController{
		ji: &jwtInteractorForTest{},
	}
}

func (in *jwtInteractorForTest) GetUserByJwt(context.Context, string) (*domain.User, error) {
	return &domain.User{
		ID:              "id",
		Username:        "username",
		Email:           "aaa@example.com",
		IsValid:         true,
		IsEmailVerified: true,
	}, nil
}

func (con *BaseController) testContextView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		fmt.Println(err)
	}
	keySet, _ := ctx.Value(domain.CtxCognitoKeySetKey).(jwk.Set)
	access, _ := ctx.Value(domain.CtxAccessKey).(domain.CtxAccess)
	response(w, r, nil, map[string]interface{}{
		"user":   user,
		"key":    keySet,
		"access": access,
	})
}

/**************************
**************************/

func Test_BaseMiddlewareCors(t *testing.T) {
	con := newBaseControllerForTest(t)
	baseUrl := "http://example.com/"

	tests := []struct {
		name       string
		method     string
		path       string
		headers    map[string]string
		wantStatus int
		wantAccess domain.CtxAccess
	}{
		{
			name:   "success",
			method: http.MethodGet,
			path:   "abc/efg",
			headers: map[string]string{
				"Origin": config.FRONT_URL,
			},
			wantStatus: http.StatusOK,
			wantAccess: domain.CtxAccess{
				Method: http.MethodGet,
				URL:    "/abc/efg",
			},
		},
		// {
		// 	name:   "fail for origin",
		// 	method: http.MethodPut,
		// 	path:   "xyz",
		// 	headers: map[string]string{
		// 		"Origin":     "https://front.example.com",
		// 		"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
		// 	},
		// 	wantStatus: 419,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, baseUrl+tt.path, nil)
			for k, v := range tt.headers {
				r.Header.Add(k, v)
			}
			defer r.Body.Close()

			got := httptest.NewRecorder()
			mid := con.BaseMiddleware(http.HandlerFunc(con.testContextView))
			mid.ServeHTTP(got, r)

			var access testContextViewBody
			if err := json.NewDecoder(got.Result().Body).Decode(&access); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.wantStatus, got.Result().StatusCode)
			assert.Equal(t, tt.wantAccess, access.Access)
		})
	}
}
