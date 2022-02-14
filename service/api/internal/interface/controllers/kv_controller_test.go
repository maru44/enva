package controllers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/maru44/enva/service/api/internal/interface/controllers"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/domain/mockdomain"
)

func newKvControllerForTest(t *testing.T) *controllers.KvController {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kI := mockdomain.NewMockIKvInteractor(ctrl)
	pI := mockdomain.NewMockIProjectInteractor(ctrl)
	oI := mockdomain.NewMockIOrgInteractor(ctrl)
	return controllers.NewKvControllerFromUsecase(kI, pI, oI)
}

func Test_KvController_ListView(t *testing.T) {
	cu := createUser(t)
	u2 := createUser(t)

	org1 := createOrg(t)

	projectUserOK := createProjectWithOwnerUser(t, cu)
	projectUserNotOK := createProjectWithOwnerUser(t, u2)

	projectOrg := createProjectWithOwnerOrg(t, org1)

	tests := []struct {
		name      string
		projectID domain.ProjectID

		mockKvs []domain.Kv

		wantStatus int
		wantKvs    []domain.Kv
		wantErr    string
	}{
		{
			name: "success",
		},
	}
}
