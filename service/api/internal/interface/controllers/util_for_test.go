package controllers_test

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
)

func uu() string {
	return uuid.New().String()
}

/**************************
	jwt
**************************/

type (
	jwtInteractor struct {
		usecase.JwtInteractor
		cookieIdToken cookieIdToken
	}

	cookieIdToken string
)

const (
	cookieIdTokenBlank   = cookieIdToken("blank")
	cookieIdTokenInvalid = cookieIdToken("invalid")
	cookieIdTokenValid   = cookieIdToken("valid")
)

var (
	testUser = domain.User{
		ID:              domain.UserID(uu()),
		Username:        "username",
		Email:           "aaa@example.com",
		IsValid:         true,
		IsEmailVerified: true,
	}

	testUser2 = domain.User{
		ID:              domain.UserID(uu()),
		Username:        "username2",
		Email:           "bbb@example.com",
		IsValid:         true,
		IsEmailVerified: true,
	}

	testOrg1 = domain.Org{
		ID:          domain.OrgID(uu()),
		Slug:        "org1",
		Name:        "org 1",
		Description: tools.StringPtr("test user belongs to"),
		IsValid:     true,
		CreatedBy:   testUser2,
		UserCount:   2,
	}

	testOrg2 = domain.Org{
		ID:          domain.OrgID(uu()),
		Slug:        "org2",
		Name:        "org 2",
		Description: tools.StringPtr("test user1 not belongs to"),
		IsValid:     true,
		CreatedBy:   testUser2,
		UserCount:   1,
	}

	testProject1 = domain.Project{
		ID:        domain.ProjectID(uu()),
		Slug:      "projectSlug1",
		Name:      "project 1",
		OwnerType: domain.OwnerTypeUser,
		IsValid:   true,
		OwnerUser: &testUser,
	}

	testProject2 = domain.Project{
		ID:        domain.ProjectID(uu()),
		Slug:      "projectSlug2",
		Name:      "project 2",
		OwnerType: domain.OwnerTypeUser,
		IsValid:   true,
		OwnerUser: &testUser,
	}

	testProject3 = domain.Project{
		ID:        domain.ProjectID(uu()),
		Slug:      "projectSlug3",
		Name:      "project 3",
		OwnerType: domain.OwnerTypeOrg,
		IsValid:   true,
		OwnerOrg:  &testOrg1,
	}
)

func (in *jwtInteractor) FetchJwk(context.Context, string) (jwk.Set, error) {
	return nil, nil
}

func (in *jwtInteractor) GetUserByJwt(context.Context, string) (*domain.User, error) {
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
	kv
**************************/

type (
	kvInteractor struct {
		usecase.KvInteractor
		hasError bool
	}
)

var (
	listKvValidSample = []domain.Kv{
		{
			ID:        domain.KvID(uu()),
			ProjectID: "projectID",
			Key:       "KEY1",
			Value:     "VALUE1",
			IsValid:   true,
		},
		{
			ID:        domain.KvID(uu()),
			ProjectID: "projectID",
			Key:       "KEY",
			Value:     "VALUE",
			IsValid:   true,
		},
		{
			ID:        domain.KvID(uu()),
			ProjectID: "projectID",
			Key:       "KEY3",
			Value:     "VALUE3",
			IsValid:   true,
		},
	}
)

func (in *kvInteractor) ListValid(ctx context.Context) ([]domain.Kv, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return listKvValidSample, nil
}

func (in *kvInteractor) DetailValid(ctx context.Context, key domain.KvKey) (*domain.Kv, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return &domain.Kv{
		ID:        "id",
		ProjectID: "projectID",
		Key:       key,
		Value:     "VALUE",
		IsValid:   true,
	}, nil
}

func (in *kvInteractor) Create(ctx context.Context, input domain.KvInput) (*domain.KvID, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return nil, nil
}

func (in *kvInteractor) Update(ctx context.Context, input domain.KvInputWithProjectID) (*domain.KvID, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return nil, nil
}

func (in *kvInteractor) Delete(ctx context.Context, kvID domain.KvID) (int, error) {
	if in.hasError {
		return 0, errors.New("some error")
	}
	return 1, nil
}

func (in *kvInteractor) DeleteByKey(ctx context.Context, key domain.KvKey, projectID domain.ProjectID) (int, error) {
	if in.hasError {
		return 0, errors.New("some error")
	}
	return 1, nil
}

/**************************
	project
**************************/

type (
	projectInteractor struct {
		usecase.ProjectInteractor
		hasError            bool
		createdProjectCount int
	}
)

func (in *projectInteractor) ListAll(context.Context) ([]domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return []domain.Project{
		testProject1, testProject2, testProject3,
	}, nil
}

func (in *projectInteractor) ListByUser(ctx context.Context) ([]domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return []domain.Project{
		testProject1, testProject2,
	}, nil
}

func (in *projectInteractor) ListByOrg(ctx context.Context, orgID domain.OrgID) ([]domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return []domain.Project{
		testProject3,
	}, nil
}

func (in *projectInteractor) SlugListByUser(ctx context.Context) ([]string, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return []string{"projectSlug1", "projectSlug2"}, nil
}

func (in *projectInteractor) GetBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return &testProject1, nil
}

func (in *projectInteractor) GetBySlugAndOrgID(ctx context.Context, slug string, orgID domain.OrgID) (*domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return &testProject3, nil
}

func (in *projectInteractor) GetBySlugAndOrgSlug(ctx context.Context, slug, orgSlug string) (*domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return &testProject3, nil
}

func (in *projectInteractor) GetByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return &testProject2, nil
}

func (in *projectInteractor) Create(ctx context.Context, input domain.ProjectInput) (*string, error) {
	if in.hasError {
		return nil, errors.New("some error")
	}
	return nil, nil
}

func (in *projectInteractor) Delete(ctx context.Context, projectID domain.ProjectID) (int, error) {
	if in.hasError {
		return 0, errors.New("some error")
	}
	return 1, nil
}

func (in *projectInteractor) CountValidByOrgID(ctx context.Context, orgID domain.OrgID) (*int, *domain.Subscription, error) {
	return tools.IntPtrAbleZero(in.createdProjectCount), nil, nil
}

func (in *projectInteractor) CountValidByOrgSlug(ctx context.Context, orgSlug string) (*int, *domain.Subscription, error) {
	return tools.IntPtrAbleZero(in.createdProjectCount), nil, nil
}

func (in *projectInteractor) CountValidByUser(ctx context.Context, userID domain.UserID) (*int, *domain.Subscription, error) {
	return tools.IntPtrAbleZero(in.createdProjectCount), nil, nil
}

/**************************
	org
**************************/

type (
	orgInteractor struct {
		usecase.OrgInteractor
		hasError            bool
		createdProjectCount int
		cretaedMemberCount  int
	}
)

func (in *orgInteractor) List(ctx context.Context) ([]domain.Org, error) {
	return nil, nil
}

func (in *orgInteractor) ListOwnerAdmin(ctx context.Context) ([]domain.Org, error) {
	return nil, nil
}

func (in *orgInteractor) DetailBySlug(ctx context.Context, slug string) (*domain.Org, *domain.UserType, error) {
	return nil, nil, nil
}

func (in *orgInteractor) Create(ctx context.Context, input domain.OrgInput) (*string, error) {
	return nil, nil
}

func (in *orgInteractor) OrgValidCount(ctx context.Context, userID domain.UserID) (*int, *domain.Subscription, error) {
	return nil, nil, nil
}

func (in *orgInteractor) Invite(ctx context.Context) error {
	return nil
}

func (in *orgInteractor) InvitationListFromOrg(context.Context, domain.OrgID) ([]domain.OrgInvitation, error) {
	return nil, nil
}

func (in *orgInteractor) InvitationList(context.Context) ([]domain.OrgInvitation, error) {
	return nil, nil
}

func (in *orgInteractor) InvitationDetail(context.Context, domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	return nil, nil
}

func (in *orgInteractor) InvitationPastList(context.Context, domain.OrgID) ([]domain.OrgInvitationID, error) {
	return nil, nil
}

func (in *orgInteractor) InvitationUpdateStatus(context.Context, domain.OrgInvitationID, domain.OrgInvitationStatus) error {
	return nil
}

func (in *orgInteractor) InvitationDeny(context.Context, domain.OrgInvitationID) error {
	return nil
}

func (in *orgInteractor) MemberCreate(context.Context, domain.OrgMemberInput) error {
	return nil
}

func (in *orgInteractor) MemberList(context.Context, domain.OrgID) (map[domain.UserType][]domain.User, error) {
	return nil, nil
}

func (in *orgInteractor) MemberGetCurrentUserType(context.Context, domain.OrgID) (*domain.UserType, error) {
	return nil, nil
}

func (in *orgInteractor) MemberGetUserType(context.Context, domain.UserID, domain.OrgID) (*domain.UserType, error) {
	return nil, nil
}

func (in *orgInteractor) MemberUpdateUserType(context.Context, domain.OrgMemberUpdateInput) error {
	return nil
}

func (in *orgInteractor) MemberDelete(context.Context, domain.UserID, domain.OrgID) error {
	return nil
}

func (in *orgInteractor) MemberValidCount(context.Context, domain.OrgID) (*int, *domain.Subscription, error) {
	return nil, nil, nil
}
