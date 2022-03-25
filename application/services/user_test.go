package services

import (
	"testing"
	"time"

	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/repository/memory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// UserServiceTestSuite defines a test suite for the package
type UserServiceTestSuite struct {
	suite.Suite

	today time.Time

	users  repositories.UsersRepository
	logger logging.Logger

	userService usecases.UserUsecase
}

// SetupTest sets the test suite up
func (suite *UserServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

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

	suite.logger, _ = logging.NewMockLogger()

	suite.userService, _ = NewUserService(suite.users, suite.logger)
}

func (suite *UserServiceTestSuite) TestNewUserService() {
	suite.T().Run("Given a set of drivers, a new UserService should be created", func(t *testing.T) {
		got, err := NewUserService(suite.users, suite.logger)
		assert.Equal(t, &UserService{users: suite.users, logger: suite.logger}, got)
		assert.Equal(t, false, err != nil)
	})
}

func (suite *UserServiceTestSuite) TestUserService_FindUserByID() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "If the userID provided is not related to any users in the users repository, an error should be returned",
			args: args{
				userID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the userID provided is related to an user in the users repository, an entity.User should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			want: &entity.User{
				ID:          "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				Username:    "analyst",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   suite.today,
				UpdatedAt:   suite.today,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.userService.FindUserByID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *UserServiceTestSuite) TestUserService_ListUsers() {
	tests := []struct {
		name    string
		want    []*entity.User
		wantErr bool
	}{
		{
			name: "If I call UserService.ListUsers, all entity.User in the repository should be returned",
			want: []*entity.User{
				{
					ID:          "e01f33c3-074f-4f89-b4df-9708ba248599",
					Username:    "undefined",
					Password:    "bar",
					ProfileKind: valueobject.Undefined_ProfileKind,
					CreatedAt:   suite.today,
					UpdatedAt:   suite.today,
				},
				{
					ID:          "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
					Username:    "analyst",
					Password:    "bar",
					ProfileKind: valueobject.Analyst_ProfileKind,
					CreatedAt:   suite.today,
					UpdatedAt:   suite.today,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.userService.ListUsers()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func (suite *UserServiceTestSuite) TestUserService_CreateUser() {
	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "If the username provided is related to any users in the users repository, an error should be returned",
			args: args{
				username: "analyst",
				password: "bar",
			},
			wantErr: true,
		},
		{
			name: "If the username provided is not related to any users in the users repository but the input password is not valid, an error should be returned",
			args: args{
				username: "pacient",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "If the username provided is not related to any users in the users repository and the input password is valid, an entity.User should be created",
			args: args{
				username: "pacient",
				password: "foo",
			},
			wantErr: false,
			want: &entity.User{
				Username:    "pacient",
				Password:    "foo",
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.userService.CreateUser(tt.args.username, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.Username, got.Username)
				assert.Equal(t, tt.want.Password, got.Password)
				assert.Equal(t, tt.want.ProfileKind, got.ProfileKind)
			}
		})
	}
}
