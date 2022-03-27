package services

import (
	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/repository/memory"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// AuthenticationServiceTestSuite defines a test suite for the package
type AuthenticationServiceTestSuite struct {
	suite.Suite

	today time.Time

	users  repositories.UsersRepository
	logger logging.Logger

	authenticationService usecases.AuthenticationUsecase
}

// SetupTest sets the test suite up
func (suite *AuthenticationServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

	suite.users = memory.NewUsersRepository()
	suite.users.CreateUser(&entity.User{
		ID:          "e01f33c3-074f-4f89-b4df-9708ba248599",
		Username:    "undefined",
		Password:    "bar",
		ProfileKind: valueobject.Undefined_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.logger, _ = logging.NewNopLogger()

	suite.authenticationService, _ = NewAuthenticationService(suite.users, suite.logger)
}

func (suite *AuthenticationServiceTestSuite) TestNewAuthenticationService() {
	suite.T().Run("Given a set of drivers, a new AuthenticationService should be created", func(t *testing.T) {
		got, err := NewAuthenticationService(suite.users, suite.logger)
		assert.Equal(t, &AuthenticationService{users: suite.users, logger: suite.logger}, got)
		assert.Equal(t, false, err != nil)
	})
}

func (suite *AuthenticationServiceTestSuite) TestAuthenticationService_AuthenticateUser() {
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
			name: "If the username provided is not related to any users in the users repository, an error should be returned",
			args: args{
				username: "foo",
			},
			wantErr: true,
		},
		{
			name: "If the username provided exists in the repository but its password doesn't match, an error should be returned",
			args: args{
				username: "undefined",
				password: "foo",
			},
			wantErr: true,
		},
		{
			name: "If the username provided exists in the repository and its password matches, an entity.User should be returned",
			args: args{
				username: "undefined",
				password: "bar",
			},
			wantErr: false,
			want: &entity.User{
				ID:          "e01f33c3-074f-4f89-b4df-9708ba248599",
				Username:    "undefined",
				Password:    "bar",
				ProfileKind: valueobject.Undefined_ProfileKind,
				CreatedAt:   suite.today,
				UpdatedAt:   suite.today,
			},
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.authenticationService.AuthenticateUser(tt.args.username, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *AuthenticationServiceTestSuite) TestAuthenticationService_AuthenticateUserFromRequest() {
	type args struct {
		r *http.Request
	}

	r0, _ := http.NewRequest(http.MethodGet, "", nil)

	r1, _ := http.NewRequest(http.MethodGet, "", nil)
	r1.SetBasicAuth("foo", "bar")

	r2, _ := http.NewRequest(http.MethodGet, "", nil)
	r2.SetBasicAuth("undefined", "foo")

	r3, _ := http.NewRequest(http.MethodGet, "", nil)
	r3.SetBasicAuth("undefined", "bar")

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "If the request provided does not have a Basic Auth header set, an error should be returned",
			args: args{
				r: r0,
			},
			wantErr: true,
		},
		{
			name: "If the username provided is not related to any users in the users repository, an error should be returned",
			args: args{
				r: r1,
			},
			wantErr: true,
		},
		{
			name: "If the username provided exists in the repository but its password doesn't match, an error should be returned",
			args: args{
				r: r2,
			},
			wantErr: true,
		},
		{
			name: "If the username provided exists in the repository and its password matches, an entity.User should be returned",
			args: args{
				r: r3,
			},
			wantErr: false,
			want: &entity.User{
				ID:          "e01f33c3-074f-4f89-b4df-9708ba248599",
				Username:    "undefined",
				Password:    "bar",
				ProfileKind: valueobject.Undefined_ProfileKind,
				CreatedAt:   suite.today,
				UpdatedAt:   suite.today,
			},
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.authenticationService.AuthenticateUserFromRequest(tt.args.r)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
