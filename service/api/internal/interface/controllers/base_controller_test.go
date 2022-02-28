package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maru44/enva/service/api/internal/interface/controllers"
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type (
	testContextViewBody struct {
		Access domain.CtxAccess `json:"access"`
		User   *domain.User     `json:"user"`
	}
)

func newBaseControllerForTest(t *testing.T, cookieIdToken cookieIdToken, user *domain.User) *controllers.BaseController {
	return controllers.NewBaseControllerFromUsecase(
		&jwtInteractor{
			cookieIdToken: cookieIdToken,
			user:          user,
		},
	)
}

func testContextView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, _ := domain.UserFromCtx(ctx)

	access, _ := ctx.Value(domain.CtxAccessKey).(domain.CtxAccess)
	w.WriteHeader(200)
	j, _ := json.Marshal(map[string]interface{}{
		"user":   user,
		"access": access,
	})
	w.Write(j) //nolint:errcheck
}

/**************************
**************************/

func Test_BaseMiddlewareCors(t *testing.T) {
	con := newBaseControllerForTest(t, cookieIdTokenBlank, nil)

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
			path:   "/abc/efg",
			headers: map[string]string{
				"Origin": config.FRONT_URL,
			},
			wantStatus: http.StatusOK,
			wantAccess: domain.CtxAccess{
				Method: http.MethodGet,
				URL:    "/abc/efg",
			},
		},
		{
			name:   "fail for origin",
			method: http.MethodOptions,
			path:   "/xyz",
			headers: map[string]string{
				"Origin":     "https://front.example.com",
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
			},
			wantStatus: 419,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.Handler(con.BaseMiddleware(http.HandlerFunc(testContextView)))
			ts := httptest.NewServer(h)
			defer ts.Close()

			r := httptest.NewRequest(tt.method, ts.URL+tt.path, nil)
			defer r.Body.Close()
			for k, v := range tt.headers {
				r.Header.Add(k, v)
			}
			r.RequestURI = ""
			cli := &http.Client{}
			got, err := cli.Do(r)
			if err != nil {
				t.Fatal(err)
			}
			defer got.Body.Close()

			var access testContextViewBody
			if err := json.NewDecoder(got.Body).Decode(&access); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.wantStatus, got.StatusCode)
			assert.Equal(t, tt.wantAccess, access.Access)
			got.Body.Close()
		})
	}
}

func Test_GiveUserMiddleware(t *testing.T) {
	conAnonymous := newBaseControllerForTest(t, cookieIdTokenBlank, nil)
	conAuth := newBaseControllerForTest(t, cookieIdTokenValid, nil)
	conInvalid := newBaseControllerForTest(t, cookieIdTokenInvalid, nil)
	baseUrl := "http://example.com/user-test"

	tests := []struct {
		name       string
		method     string
		con        *controllers.BaseController
		wantStatus int
		wantUser   *domain.User
	}{
		{
			name:       "success anonymous",
			method:     http.MethodPost,
			con:        conAnonymous,
			wantStatus: http.StatusOK,
			wantUser:   nil,
		},
		{
			name:       "success annonymous invalid cookie",
			method:     http.MethodGet,
			con:        conInvalid,
			wantStatus: http.StatusOK,
			wantUser:   nil,
		},
		{
			name:       "success authenticated",
			method:     http.MethodDelete,
			con:        conAuth,
			wantStatus: http.StatusOK,
			wantUser:   &testUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, baseUrl, nil)
			defer r.Body.Close()
			con := tt.con
			if con == conAuth || con == conInvalid {
				r.Header.Add("Cookie", domain.JwtCookieKeyIdToken+"=a")
			}

			got := httptest.NewRecorder()
			mid := con.BaseMiddleware(con.GiveUserMiddleware(http.HandlerFunc(testContextView)))
			mid.ServeHTTP(got, r)

			var bod testContextViewBody
			if err := json.NewDecoder(got.Result().Body).Decode(&bod); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.wantStatus, got.Code)
			assert.Equal(t, tt.wantUser, bod.User)
			got.Result().Body.Close()
		})
	}
}

func Test_LoginRequiredMiddleware(t *testing.T) {
	conAnonymous := newBaseControllerForTest(t, cookieIdTokenBlank, nil)
	conAuth := newBaseControllerForTest(t, cookieIdTokenValid, nil)
	conInvalid := newBaseControllerForTest(t, cookieIdTokenInvalid, nil)
	baseUrl := "http://example.com/user-test"

	tests := []struct {
		name       string
		method     string
		con        *controllers.BaseController
		wantStatus int
		wantUser   *domain.User
	}{
		{
			name:       "failed anonymous",
			method:     http.MethodPost,
			con:        conAnonymous,
			wantStatus: http.StatusUnauthorized,
			wantUser:   nil,
		},
		{
			name:       "failed annonymous invalid cookie",
			method:     http.MethodPost,
			con:        conInvalid,
			wantStatus: http.StatusForbidden,
			wantUser:   nil,
		},
		{
			name:       "success authenticated",
			method:     http.MethodDelete,
			con:        conAuth,
			wantStatus: http.StatusOK,
			wantUser:   &testUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, baseUrl, nil)
			defer r.Body.Close()
			con := tt.con
			if con == conAuth || con == conInvalid {
				r.Header.Add("Cookie", domain.JwtCookieKeyIdToken+"=a")
			}

			got := httptest.NewRecorder()
			mid := con.BaseMiddleware(con.LoginRequiredMiddleware(http.HandlerFunc(testContextView)))
			mid.ServeHTTP(got, r)

			assert.Equal(t, tt.wantStatus, got.Code)

			if tt.wantStatus == 200 {
				var bod testContextViewBody
				if err := json.NewDecoder(got.Result().Body).Decode(&bod); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantUser, bod.User)
			}
			got.Result().Body.Close()
		})
	}
}
