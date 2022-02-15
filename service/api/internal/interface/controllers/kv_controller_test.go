package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/maru44/enva/service/api/internal/interface/controllers"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/domain/mockdomain"
	"github.com/stretchr/testify/assert"
)

// func newKvController(t *testing.T) *controllers.KvController {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	kI := mockdomain.NewMockIKvInteractor(ctrl)
// 	pI := mockdomain.NewMockIProjectInteractor(ctrl)
// 	oI := mockdomain.NewMockIOrgInteractor(ctrl)
// 	return controllers.NewKvControllerFromUsecase(kI, pI, oI)
// }

func newKvController(kv domain.IKvInteractor, p domain.IProjectInteractor, o domain.IOrgInteractor) *controllers.KvController {
	return controllers.NewKvControllerFromUsecase(kv, p, o)
}

func Test_KvController_ListView(t *testing.T) {
	type body struct {
		Data []domain.Kv `json:"data"`
	}

	url := "https://example.com/kv"

	cu := createUser(t)
	u2 := createUser(t)

	projectUserOK := createProjectWithOwnerUser(t, cu)
	projectUserNotOK := createProjectWithOwnerUser(t, u2)

	org1 := createOrg(t)
	projectOrg := createProjectWithOwnerOrg(t, org1)
	uTypeGuest := domain.UserTypeGuest

	baseCon := newBaseControllerForTest(t, cookieIdTokenValid, cu)

	kvsUserOk := createKvs(t, projectUserOK.ID)
	kvsOrg := createKvs(t, projectOrg.ID)

	tests := []struct {
		name    string
		project *domain.Project

		mockKvs []domain.Kv
		mockErr error

		mockMemberCurrentUserType    *domain.UserType
		mockMemberCurrentUserTypeErr error

		wantStatus int
		wantKvs    []domain.Kv
	}{
		{
			name:       "success",
			project:    projectUserOK,
			mockKvs:    kvsUserOk,
			mockErr:    nil,
			wantKvs:    kvsUserOk,
			wantStatus: 200,
		},
		{
			name:       "failed cannot access to project",
			project:    projectUserNotOK,
			mockKvs:    nil,
			mockErr:    nil,
			wantKvs:    nil,
			wantStatus: 403,
		},
		{
			name:                         "success ownerType org",
			project:                      projectOrg,
			mockKvs:                      kvsOrg,
			mockErr:                      nil,
			mockMemberCurrentUserType:    &uTypeGuest,
			mockMemberCurrentUserTypeErr: nil,
			wantKvs:                      kvsOrg,
			wantStatus:                   200,
		},
		{
			name:                         "failed cannot access to project (not org member)",
			project:                      projectOrg,
			mockKvs:                      kvsOrg,
			mockErr:                      nil,
			mockMemberCurrentUserType:    nil,
			mockMemberCurrentUserTypeErr: errors.New("some"),
			wantKvs:                      kvsOrg,
			wantStatus:                   403,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mock usecase
			pI := mockdomain.NewMockIProjectInteractor(ctrl)
			kI := mockdomain.NewMockIKvInteractor(ctrl)
			oI := mockdomain.NewMockIOrgInteractor(ctrl)
			pI.EXPECT().GetByID(gomock.Any(), tt.project.ID).Return(tt.project, nil)
			if tt.project.OwnerOrg != nil {
				oI.EXPECT().MemberGetCurrentUserType(gomock.Any(), tt.project.OwnerOrg.ID).Return(tt.mockMemberCurrentUserType, tt.mockMemberCurrentUserTypeErr).Times(1)
			}
			if tt.project.OwnerUser == cu || (tt.mockMemberCurrentUserType != nil && tt.mockMemberCurrentUserTypeErr == nil) {
				kI.EXPECT().ListValid(gomock.Any(), tt.project.ID).Return(tt.mockKvs, tt.mockErr).Times(1)
			}
			con := newKvController(kI, pI, oI)

			r := httptest.NewRequest(http.MethodGet, url+"?projectId="+string(tt.project.ID), nil)
			defer r.Body.Close()
			r.Header.Add("Cookie", domain.JwtCookieKeyIdToken+"=a")

			got := httptest.NewRecorder()
			end := baseCon.BaseMiddleware(baseCon.GetOnlyMiddleware(baseCon.LoginRequiredMiddleware(http.HandlerFunc(con.ListView))))
			end.ServeHTTP(got, r)

			assert.Equal(t, tt.wantStatus, got.Code)
			if got.Code == 200 {
				var bod body
				if err := json.NewDecoder(got.Result().Body).Decode(&bod); err != nil {
					t.Fatal(err)
				}
				assert.ElementsMatch(t, tt.wantKvs, bod.Data)
			}
			got.Result().Body.Close()
		})
	}
}
