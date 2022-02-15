package controllers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
)

type (
	errBody struct {
		Err    string `json:"error"`
		Status int    `json:"status"`
	}
)

func uu() string {
	return uuid.New().String()
}

var (
	testUserID = uu()

	testUser = domain.User{
		ID:              domain.UserID(testUserID),
		Username:        "username",
		Email:           "aaa@example.com",
		IsValid:         true,
		IsEmailVerified: true,
	}
)

/**************************
	jwt
**************************/

type (
	jwtInteractor struct {
		usecase.JwtInteractor
		cookieIdToken cookieIdToken
		user          *domain.User
	}

	cookieIdToken string
)

const (
	cookieIdTokenBlank   = cookieIdToken("blank")
	cookieIdTokenInvalid = cookieIdToken("invalid")
	cookieIdTokenValid   = cookieIdToken("valid")
)

func (in *jwtInteractor) FetchJwk(context.Context, string) (jwk.Set, error) {
	return nil, nil
}

func (in *jwtInteractor) GetUserByJwt(context.Context, string) (*domain.User, error) {
	if in.user != nil {
		return in.user, nil
	}
	switch in.cookieIdToken {
	case cookieIdTokenBlank:
		return nil, nil
	case cookieIdTokenInvalid:
		return nil, errors.New("invalid cookie")
	case cookieIdTokenValid:
		return &testUser, nil
	default:
		panic("must not reach here")
	}
}

/**************************
	cu
**************************/

func createUser(t *testing.T) *domain.User {
	emailPre := tools.GenRandSlug(6)
	username := tools.GenRandSlug(8)
	return &domain.User{
		ID:              domain.UserID(uu()),
		Username:        username,
		Email:           emailPre + "@example.com",
		IsValid:         true,
		IsEmailVerified: true,
	}
}

func createKv(t *testing.T, projectID domain.ProjectID) *domain.Kv {
	return &domain.Kv{
		ID:        domain.KvID(uu()),
		ProjectID: projectID,
		Key:       domain.KvKey(tools.GenRandSlug(6)),
		Value:     domain.KvValue(tools.GenRandSlug(8)),
		IsValid:   true,
	}
}

func createKvs(t *testing.T, projectID domain.ProjectID) []domain.Kv {
	return []domain.Kv{
		{
			ID:        domain.KvID(uu()),
			ProjectID: projectID,
			Key:       domain.KvKey(tools.GenRandSlug(6)),
			Value:     domain.KvValue(tools.GenRandSlug(8)),
			IsValid:   true,
		},
		{
			ID:        domain.KvID(uu()),
			ProjectID: projectID,
			Key:       domain.KvKey(tools.GenRandSlug(6)),
			Value:     domain.KvValue(tools.GenRandSlug(8)),
			IsValid:   true,
		},
		{
			ID:        domain.KvID(uu()),
			ProjectID: projectID,
			Key:       domain.KvKey(tools.GenRandSlug(6)),
			Value:     domain.KvValue(tools.GenRandSlug(8)),
			IsValid:   true,
		},
	}
}

func createProjectWithOwnerUser(t *testing.T, user *domain.User) *domain.Project {
	slug := tools.GenRandSlug(6)
	return &domain.Project{
		ID:        domain.ProjectID(uu()),
		Slug:      slug,
		Name:      slug + " name",
		OwnerType: domain.OwnerTypeUser,
		IsValid:   true,
		OwnerUser: user,
	}
}

func createProjectWithOwnerOrg(t *testing.T, org *domain.Org) *domain.Project {
	slug := tools.GenRandSlug(6)
	return &domain.Project{
		ID:        domain.ProjectID(uu()),
		Slug:      slug,
		Name:      slug + " name",
		OwnerType: domain.OwnerTypeOrg,
		IsValid:   true,
		OwnerOrg:  org,
	}
}

func createOrg(t *testing.T) *domain.Org {
	slug := tools.GenRandSlug(7)
	u := createUser(t)
	return &domain.Org{
		ID:        domain.OrgID(uu()),
		Slug:      slug,
		Name:      slug + " name",
		CreatedBy: *u,
	}
}
