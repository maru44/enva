package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/maru44/enva/service/api/internal/interface/controllers"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/domain/mockdomain"
	"github.com/maru44/enva/service/api/pkg/tools"
	"github.com/maru44/perr"
	"github.com/stretchr/testify/assert"
)

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

			// dummy con and request
			con := newKvController(kI, pI, oI)
			r := httptest.NewRequest(http.MethodGet, url+"?projectId="+string(tt.project.ID), nil)
			defer r.Body.Close()
			r.Header.Add("Cookie", domain.JwtCookieKeyIdToken+"=a")

			got := httptest.NewRecorder()
			end := baseCon.BaseMiddleware(baseCon.GetOnlyMiddleware(baseCon.LoginRequiredMiddleware(http.HandlerFunc(con.ListView))))
			end.ServeHTTP(got, r)

			// evaluate
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

func Test_KvController_CreateView(t *testing.T) {
	type body struct {
		Data *domain.KvID `json:"data"`
	}

	url := "https://example.com/kv/create"

	cu := createUser(t)
	u2 := createUser(t)

	projectUserOK := createProjectWithOwnerUser(t, cu)
	projectUserNotOK := createProjectWithOwnerUser(t, u2)

	org1 := createOrg(t)
	projectOrg := createProjectWithOwnerOrg(t, org1)
	uTypeGuest := domain.UserTypeGuest
	uTypeUser := domain.UserTypeUser

	uuidOK := domain.KvID(uu())

	baseCon := newBaseControllerForTest(t, cookieIdTokenValid, cu)

	tests := []struct {
		name    string
		project *domain.Project
		input   domain.KvInputWithProjectID

		mockInserted *domain.KvID
		mockErr      error

		mockMemberCurrentUserType    *domain.UserType
		mockMemberCurrentUserTypeErr error

		wantStatus int
	}{
		{
			name:         "success",
			project:      projectUserOK,
			input:        kvInputWithProjectID(projectUserOK.ID, true),
			mockInserted: &uuidOK,
			mockErr:      nil,
			wantStatus:   200,
		},
		{
			name:         "failed not valid input",
			project:      projectUserOK,
			input:        kvInputWithProjectID(projectUserOK.ID, false),
			mockInserted: nil,
			mockErr:      perr.New("a", perr.ErrBadRequest),
			wantStatus:   400,
		},
		{
			name:       "failed cannot access to project",
			project:    projectUserNotOK,
			input:      kvInputWithProjectID(projectUserNotOK.ID, true),
			wantStatus: 403,
		},
		{
			name:                         "success ownerType org",
			project:                      projectOrg,
			input:                        kvInputWithProjectID(projectOrg.ID, true),
			mockInserted:                 &uuidOK,
			mockMemberCurrentUserType:    &uTypeUser,
			mockMemberCurrentUserTypeErr: nil,
			wantStatus:                   200,
		},
		{
			name:                         "failed cannot access to project (guest user)",
			project:                      projectOrg,
			input:                        kvInputWithProjectID(projectOrg.ID, true),
			mockInserted:                 nil,
			mockMemberCurrentUserType:    &uTypeGuest,
			mockMemberCurrentUserTypeErr: nil,
			wantStatus:                   403,
		},
		{
			name:                         "failed cannot access to project (not member)",
			project:                      projectOrg,
			input:                        kvInputWithProjectID(projectOrg.ID, true),
			mockInserted:                 nil,
			mockMemberCurrentUserType:    nil,
			mockMemberCurrentUserTypeErr: errors.New("some"),
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
			if tt.project.OwnerUser == cu || (tt.mockMemberCurrentUserType != nil && tt.mockMemberCurrentUserTypeErr == nil && *tt.mockMemberCurrentUserType != domain.UserTypeGuest) {
				kI.EXPECT().Create(gomock.Any(), tt.input).Return(tt.mockInserted, tt.mockErr).Times(1)
			}

			// dummy con and request
			con := newKvController(kI, pI, oI)
			j, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			r := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(j))
			defer r.Body.Close()
			r.Header.Add("Cookie", domain.JwtCookieKeyIdToken+"=a")

			got := httptest.NewRecorder()
			end := baseCon.BaseMiddleware(baseCon.PostOnlyMiddleware(baseCon.LoginRequiredMiddleware(http.HandlerFunc(con.CreateView))))
			end.ServeHTTP(got, r)

			assert.Equal(t, tt.wantStatus, got.Code)
			if got.Code == 200 {
				var bod body
				if err := json.NewDecoder(got.Result().Body).Decode(&bod); err != nil {
					t.Fatal(err)
				}
				assert.NotNil(t, bod.Data)
			}
			got.Result().Body.Close()
		})
	}
}

func kvInputWithProjectID(projectID domain.ProjectID, isValid bool) domain.KvInputWithProjectID {
	if !isValid {
		return domain.KvInputWithProjectID{
			ProjectID: projectID,
			Input: domain.KvInput{
				Value: domain.KvValue(tools.GenRandSlug(8)),
			},
		}
	}

	return domain.KvInputWithProjectID{
		ProjectID: projectID,
		Input: domain.KvInput{
			Key:   domain.KvKey(tools.GenRandSlug(8)),
			Value: domain.KvValue(tools.GenRandSlug(16)),
		},
	}
}
