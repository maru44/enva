package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type jwtInteractorForTest struct {
	usecase.JwtInteractor
}

func newBaseControllerForTest(t *testing.T) *BaseController {
	return &BaseController{
		ji: &jwtInteractorForTest{},
	}
}

func (in *jwtInteractorForTest) Evaluate(context.Context, string) (*jwt.Token, error) {
	return nil, nil
}

func (in *jwtInteractorForTest) GetUserByJwt(context.Context, string) (*domain.User, error) {
	return nil, nil
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

func Test_GetKeySet(t *testing.T) {
	_, err := jwk.Fetch(context.Background(), config.COGNITO_KEYS_URL)
	assert.NoError(t, err)
}

func Test_BaseMiddlewareCors(t *testing.T) {
	con := newBaseControllerForTest(t)
	baseUrl := "http://example.com/"

	// pseudo server
	keySet, err := con.GetKeySet()
	if err != nil {
		t.Fatal(err)
	}

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
			},
		},
		// {
		// 	name:   "fail for origin",
		// 	method: http.MethodGet,
		// 	path:   "xyz",
		// 	headers: map[string]string{
		// 		"Origin": "https://front.example.com",
		// 	},
		// 	wantStatus: 419,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, baseUrl+tt.path, nil)
			for k, v := range tt.headers {
				r.Header.Add(k, v)
			}
			defer r.Body.Close()

			got := httptest.NewRecorder()
			mid := con.BaseMiddleware(keySet, http.HandlerFunc(con.testContextView))
			mid.ServeHTTP(got, r)

			assert.Equal(t, tt.wantStatus, got.Result().StatusCode)
		})
	}
}
