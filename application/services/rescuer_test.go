package services

import (
	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/repository/memory"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// RescuerServiceTestSuite defines a test suite for the package
type RescuerServiceTestSuite struct {
	suite.Suite

	today time.Time

	rescuers repositories.RescuersRepository
	users    repositories.UsersRepository
	logger   logging.Logger

	rescuerService usecases.RescuerUsecase
}

// SetupTest sets the test suite up
func (suite *RescuerServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

	suite.rescuers, _ = memory.NewRescuersRepository()
	suite.rescuers.CreateRescuer(&entity.Rescuer{
		ID:        "47322c6f-5883-4596-a305-29be7395ddd1",
		UserID:    "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
		CreatedAt: suite.today,
		UpdatedAt: suite.today,
	})

	suite.users, _ = memory.NewUsersRepository()
	suite.users.CreateUser(&entity.User{
		ID:          "e01f33c3-074f-4f89-b4df-9708ba248599",
		Username:    "undefined",
		Password:    "bar",
		ProfileKind: valueobject.Undefined_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
		Username:    "rescuer",
		Password:    "bar",
		ProfileKind: valueobject.Rescuer_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
		Username:    "analyst",
		Password:    "bar",
		ProfileKind: valueobject.Analyst_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "b6d08f7a-e202-45a8-bbee-a827bc6ab8b5",
		Username:    "pacient",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "196f8233-4d02-44d3-a4a5-1c4a89b17b66",
		Username:    "rescuer",
		Password:    "bar",
		ProfileKind: valueobject.Rescuer_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.logger, _ = logging.NewMockLogger()

	suite.rescuerService, _ = NewRescuerService(suite.rescuers, suite.users, suite.logger)
}

func (suite *RescuerServiceTestSuite) TestNewRescuerService() {
	suite.T().Run("Given a set of drivers, a new RescuerService should be created", func(t *testing.T) {
		got, err := NewRescuerService(suite.rescuers, suite.users, suite.logger)
		assert.Equal(t, &RescuerService{rescuers: suite.rescuers, users: suite.users, logger: suite.logger}, got)
		assert.Equal(t, false, err != nil)
	})
}

func (suite *RescuerServiceTestSuite) TestRescuerService_CreateRescuer() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Rescuer
		wantErr bool
	}{
		{
			name: "If the userID provided is not related to any users in the users repository, an error should be returned",
			args: args{
				userID: "foo",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided is related to a user that already has a defined ProfileKind but is not an Rescuer, an error should be returned",
			args: args{
				userID: "b6d08f7a-e202-45a8-bbee-a827bc6ab8b5",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided is related to an Rescuer user that already has an entity.Rescuer, an error should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided has an undefined ProfileKind, a new Rescuer should be created",
			args: args{
				userID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			want: &entity.Rescuer{
				UserID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			wantErr: false,
		},
		{
			name: "If the userID provided is related to an Rescuer user that failed to create an entity.Rescuer previously, a new Rescuer should be created",
			args: args{
				userID: "196f8233-4d02-44d3-a4a5-1c4a89b17b66",
			},
			want: &entity.Rescuer{
				UserID: "196f8233-4d02-44d3-a4a5-1c4a89b17b66",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.rescuerService.CreateRescuer(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.UserID, got.UserID)
			}
		})
	}
}

func (suite *RescuerServiceTestSuite) TestRescuerService_FindRescuerByID() {
	type args struct {
		rescuerID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Rescuer
		wantErr bool
	}{
		{
			name: "If the rescuerID provided is not related to any rescuers in the rescuers repository, an error should be returned",
			args: args{
				rescuerID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the rescuerID provided is related to an rescuer in the rescuers repository, an entity.Rescuer should be returned",
			args: args{
				rescuerID: "47322c6f-5883-4596-a305-29be7395ddd1",
			},
			want: &entity.Rescuer{
				ID:        "47322c6f-5883-4596-a305-29be7395ddd1",
				UserID:    "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				CreatedAt: suite.today,
				UpdatedAt: suite.today,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.rescuerService.FindRescuerByID(tt.args.rescuerID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
