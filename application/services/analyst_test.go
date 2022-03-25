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

// AnalystServiceTestSuite defines a test suite for the package
type AnalystServiceTestSuite struct {
	suite.Suite

	today time.Time

	analysts repositories.AnalystsRepository
	users    repositories.UsersRepository
	logger   logging.Logger

	analystService usecases.AnalystUsecase
}

// SetupTest sets the test suite up
func (suite *AnalystServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

	suite.analysts, _ = memory.NewAnalystsRepository()
	suite.analysts.CreateAnalyst(&entity.Analyst{
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
		Username:    "analyst",
		Password:    "bar",
		ProfileKind: valueobject.Analyst_ProfileKind,
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

	suite.analystService, _ = NewAnalystService(suite.analysts, suite.users, suite.logger)
}

func (suite *AnalystServiceTestSuite) TestNewAnalystService() {
	suite.T().Run("Given a set of drivers, a new AnalystService should be created", func(t *testing.T) {
		got, err := NewAnalystService(suite.analysts, suite.users, suite.logger)
		assert.Equal(t, &AnalystService{analysts: suite.analysts, users: suite.users, logger: suite.logger}, got)
		assert.Equal(t, false, err != nil)
	})
}

func (suite *AnalystServiceTestSuite) TestAnalystService_CreateAnalyst() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Analyst
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
			name: "If the userID provided is related to a user that already has a defined ProfileKind but is not an Analyst, an error should be returned",
			args: args{
				userID: "b6d08f7a-e202-45a8-bbee-a827bc6ab8b5",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided is related to an Analyst user that already has an entity.Analyst, an error should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided has an undefined ProfileKind, a new Analyst should be created",
			args: args{
				userID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			want: &entity.Analyst{
				UserID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			wantErr: false,
		},
		{
			name: "If the userID provided is related to an Analyst user that failed to create an entity.Analyst previously, a new Analyst should be created",
			args: args{
				userID: "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
			},
			want: &entity.Analyst{
				UserID: "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.analystService.CreateAnalyst(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.UserID, got.UserID)
			}
		})
	}
}

func (suite *AnalystServiceTestSuite) TestAnalystService_FindAnalystByID() {
	type args struct {
		analystID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Analyst
		wantErr bool
	}{
		{
			name: "If the analystID provided is not related to any analysts in the analysts repository, an error should be returned",
			args: args{
				analystID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the analystID provided is related to an analyst in the analysts repository, an entity.Analyst should be returned",
			args: args{
				analystID: "47322c6f-5883-4596-a305-29be7395ddd1",
			},
			want: &entity.Analyst{
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
			got, err := suite.analystService.FindAnalystByID(tt.args.analystID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
