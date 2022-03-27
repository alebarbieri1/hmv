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

// PacientServiceTestSuite defines a test suite for the package
type PacientServiceTestSuite struct {
	suite.Suite

	today time.Time

	pacients repositories.PacientsRepository
	users    repositories.UsersRepository
	logger   logging.Logger

	pacientService usecases.PacientUsecase
}

// SetupTest sets the test suite up
func (suite *PacientServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

	suite.pacients = memory.NewPacientsRepository()
	suite.pacients.CreatePacient(&entity.Pacient{
		ID:     "47322c6f-5883-4596-a305-29be7395ddd1",
		UserID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
		EmergencyContact: valueobject.EmergencyContact{
			Name:         "foo",
			MobileNumber: "5511999999999",
		},
		CreatedAt: suite.today,
		UpdatedAt: suite.today,
	})

	suite.users = memory.NewUsersRepository()
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
		Username:    "pacient",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
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

	suite.logger = logging.NewNopLogger()

	suite.pacientService = NewPacientService(suite.pacients, suite.users, suite.logger)
}

func (suite *PacientServiceTestSuite) TestNewPacientService() {
	suite.T().Run("Given a set of drivers, a new PacientService should be created", func(t *testing.T) {
		assert.Equal(t, &PacientService{pacients: suite.pacients, users: suite.users, logger: suite.logger}, NewPacientService(suite.pacients, suite.users, suite.logger))
	})
}

func (suite *PacientServiceTestSuite) TestPacientService_FindPacientByID() {
	type args struct {
		pacientID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Pacient
		wantErr bool
	}{
		{
			name: "If the pacientID provided is not related to any pacients in the pacients repository, an error should be returned",
			args: args{
				pacientID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the pacientID provided is related to an pacient in the pacients repository, an entity.Pacient should be returned",
			args: args{
				pacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
			},
			want: &entity.Pacient{
				ID:     "47322c6f-5883-4596-a305-29be7395ddd1",
				UserID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
				CreatedAt: suite.today,
				UpdatedAt: suite.today,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.pacientService.FindPacientByID(tt.args.pacientID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *PacientServiceTestSuite) TestPacientService_FindPacientByUserID() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Pacient
		wantErr bool
	}{
		{
			name: "If the userID provided is not related to any pacients in the pacients repository, an error should be returned",
			args: args{
				userID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the userID provided is related to an pacient in the pacients repository, an entity.Pacient should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			want: &entity.Pacient{
				ID:     "47322c6f-5883-4596-a305-29be7395ddd1",
				UserID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
				CreatedAt: suite.today,
				UpdatedAt: suite.today,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.pacientService.FindPacientByUserID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *PacientServiceTestSuite) TestPacientService_UpdateEmergencyContact() {
	type args struct {
		userID           string
		pacientID        string
		emergencyContact valueobject.EmergencyContact
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Pacient
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
			name: "If the userID provided is related to a user but it is not a Pacient, an error should be returned",
			args: args{
				userID: "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the pacientID provided is not related to any pacients in the pacients repository, an error should be returned",
			args: args{
				userID:    "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				pacientID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the emergency contact data provided is not valid, an error should be returned",
			args: args{
				userID:    "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				pacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
				emergencyContact: valueobject.EmergencyContact{
					Name:         "",
					MobileNumber: "99999999999999999",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If all input data are valid, the Pacient.EmergencyContact should be updated",
			args: args{
				userID:    "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				pacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
				emergencyContact: valueobject.EmergencyContact{
					Name:         "bar",
					MobileNumber: "5519999999999",
				},
			},
			want: &entity.Pacient{
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "bar",
					MobileNumber: "5519999999999",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.pacientService.UpdateEmergencyContact(tt.args.userID, tt.args.pacientID, tt.args.emergencyContact)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.EmergencyContact, got.EmergencyContact)
			}
		})
	}
}

func (suite *PacientServiceTestSuite) TestPacientService_CreatePacient() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Pacient
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
			name: "If the userID provided is related to a user that already has a defined ProfileKind but is not an Pacient, an error should be returned",
			args: args{
				userID: "196f8233-4d02-44d3-a4a5-1c4a89b17b66",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided is related to an Pacient user that already has an entity.Pacient, an error should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			wantErr: true,
		},
		{
			name: "If the userID provided has an undefined ProfileKind, a new Pacient should be created",
			args: args{
				userID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			want: &entity.Pacient{
				UserID: "e01f33c3-074f-4f89-b4df-9708ba248599",
			},
			wantErr: false,
		},
		{
			name: "If the userID provided is related to an Pacient user that failed to create an entity.Pacient previously, a new Pacient should be created",
			args: args{
				userID: "b6d08f7a-e202-45a8-bbee-a827bc6ab8b5",
			},
			want: &entity.Pacient{
				UserID: "b6d08f7a-e202-45a8-bbee-a827bc6ab8b5",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.pacientService.CreatePacient(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.UserID, got.UserID)
			}
		})
	}
}
