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

// EmergencyServiceTestSuite defines a test suite for the package
type EmergencyServiceTestSuite struct {
	suite.Suite

	today time.Time

	emergencies repositories.EmergenciesRepository
	pacients    repositories.PacientsRepository
	users       repositories.UsersRepository
	logger      logging.Logger

	emergencyService usecases.EmergencyUsecase
}

// SetupTest sets the test suite up
func (suite *EmergencyServiceTestSuite) SetupTest() {
	suite.today = time.Date(2022, time.February, 22, 0, 0, 0, 0, time.UTC)

	_true := true

	suite.emergencies = memory.NewEmergenciesRepository()
	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:        "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
		PacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
		CreatedAt: suite.today,
		UpdatedAt: suite.today,
		Form: valueobject.EmergencyForm{
			Headache: valueobject.HeadacheEmergencyFormSession{
				Has:       &_true,
				Intensity: valueobject.High_HeadacheIntensity,
			},
		},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Triage_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "acffa8b8-5624-4160-8daf-77d53e89ddb2",
		PacientID:  "47322c6f-5883-4596-a305-29be7395ddd1",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Cancelled_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "0526c40e-045a-48f6-895f-4f5030ba797d",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.AmbulanceToPacient_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "ff23c286-5321-47a9-973d-b34c4b1ff511",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Undefined_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "70d3ca7c-9491-4aa2-8aaa-507517b782c2",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Triage_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "44639d78-4501-4d9b-aeac-59295e636b82",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.AmbulanceToPacient_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "8ffa7ebf-7dc7-4bff-8fba-bba0661c9a51",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.AmbulanceToHospital_EmergencyStatus,
	})

	suite.emergencies.CreateEmergency(&entity.Emergency{
		ID:         "4373dae2-cec1-42f5-9080-37528ef21245",
		PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		CreatedAt:  suite.today,
		UpdatedAt:  suite.today,
		Form:       valueobject.EmergencyForm{},
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Finished_EmergencyStatus,
	})

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

	suite.pacients.CreatePacient(&entity.Pacient{
		ID:     "28535166-c4a4-429a-8a8c-d26ea12ee132",
		UserID: "5f1f0411-4d68-4b9a-a393-3671fd655a19",
		EmergencyContact: valueobject.EmergencyContact{
			Name:         "foo",
			MobileNumber: "5511999999999",
		},
		CreatedAt: suite.today,
		UpdatedAt: suite.today,
	})

	suite.pacients.CreatePacient(&entity.Pacient{
		ID:     "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
		UserID: "9b55f726-e47c-4371-b7bd-3ce6e622b147",
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
		ID:          "fbff91e2-b65a-478d-bda6-a1699a4bcecf",
		Username:    "pacient",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "5f1f0411-4d68-4b9a-a393-3671fd655a19",
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
		ID:          "9b55f726-e47c-4371-b7bd-3ce6e622b147",
		Username:    "rescuer",
		Password:    "bar",
		ProfileKind: valueobject.Rescuer_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.users.CreateUser(&entity.User{
		ID:          "f24d25de-4531-40e5-aef8-eedc01d1f103",
		Username:    "pacient",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
		CreatedAt:   suite.today,
		UpdatedAt:   suite.today,
	})

	suite.logger, _ = logging.NewNopLogger()

	suite.emergencyService, _ = NewEmergencyService(suite.emergencies, suite.pacients, suite.users, suite.logger)
}

func (suite *EmergencyServiceTestSuite) TestNewEmergencyService() {
	suite.T().Run("Given a set of drivers, a new EmergencyService should be created", func(t *testing.T) {
		got, err := NewEmergencyService(suite.emergencies, suite.pacients, suite.users, suite.logger)
		assert.Equal(t,
			&EmergencyService{
				emergencies: suite.emergencies,
				pacients:    suite.pacients,
				users:       suite.users,
				logger:      suite.logger,
			},
			got,
		)
		assert.Equal(t, false, err != nil)
	})
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_FindEmergencyByID() {
	type args struct {
		emergencyID string
	}

	_true := true

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
		wantErr bool
	}{
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the emergencyID provided is related to an emergency in the emergencies repository, an entity.Emergency should be returned",
			args: args{
				emergencyID: "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
			},
			want: &entity.Emergency{
				ID:        "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
				PacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
				CreatedAt: suite.today,
				UpdatedAt: suite.today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.High_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.FindEmergencyByID(tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_ListEmergencies() {
	_true := true

	tests := []struct {
		name    string
		want    []*entity.Emergency
		wantErr bool
	}{
		{
			name: "If I call EmergencyService.ListEmergencies, all entity.User in the repository should be returned",
			want: []*entity.Emergency{
				{
					ID:        "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
					PacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
					CreatedAt: suite.today,
					UpdatedAt: suite.today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.High_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
				{
					ID:         "acffa8b8-5624-4160-8daf-77d53e89ddb2",
					PacientID:  "47322c6f-5883-4596-a305-29be7395ddd1",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Cancelled_EmergencyStatus,
				},
				{
					ID:         "0526c40e-045a-48f6-895f-4f5030ba797d",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.AmbulanceToPacient_EmergencyStatus,
				},
				{
					ID:         "ff23c286-5321-47a9-973d-b34c4b1ff511",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
				{
					ID:         "70d3ca7c-9491-4aa2-8aaa-507517b782c2",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
				{
					ID:         "44639d78-4501-4d9b-aeac-59295e636b82",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.AmbulanceToPacient_EmergencyStatus,
				},
				{
					ID:         "8ffa7ebf-7dc7-4bff-8fba-bba0661c9a51",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.AmbulanceToHospital_EmergencyStatus,
				},
				{
					ID:         "4373dae2-cec1-42f5-9080-37528ef21245",
					PacientID:  "50e8fca5-30d8-4ff8-9782-1083d81bf1ea",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Finished_EmergencyStatus,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.ListEmergencies()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_ListEmergenciesByStatus() {
	type args struct {
		status valueobject.EmergencyStatus
	}

	tests := []struct {
		name    string
		args    args
		want    []*entity.Emergency
		wantErr bool
	}{
		{
			name: "Given a EmergencyStatus related to some Emergency in the database, all matches should be returned",
			args: args{
				status: valueobject.Cancelled_EmergencyStatus,
			},
			want: []*entity.Emergency{
				{
					ID:         "acffa8b8-5624-4160-8daf-77d53e89ddb2",
					PacientID:  "47322c6f-5883-4596-a305-29be7395ddd1",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Cancelled_EmergencyStatus,
				},
			},
			wantErr: false,
		},
		{
			name: "Given a EmergencyStatus unrelated to any Emergency in the database, no matches should be returned",
			args: args{
				status: -1,
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.ListEmergenciesByStatus(tt.args.status)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_ListEmergenciesByUser() {
	type args struct {
		userID string
	}

	_true := true

	tests := []struct {
		name    string
		args    args
		want    []*entity.Emergency
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
				userID: "9b55f726-e47c-4371-b7bd-3ce6e622b147",
			},
			wantErr: true,
		},
		{
			name: "If the user provided is a Pacient, but is not related to any pacients in the pacients repository, an error should be returned",
			args: args{
				userID: "fbff91e2-b65a-478d-bda6-a1699a4bcecf",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the user provided is a Pacient but has no Emergencies related to him, no matches should be returned",
			args: args{
				userID: "5f1f0411-4d68-4b9a-a393-3671fd655a19",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "If the user provided is a Pacient and has any Emergencies related to him, all matches should be returned",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			want: []*entity.Emergency{
				{
					ID:        "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
					PacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
					CreatedAt: suite.today,
					UpdatedAt: suite.today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.High_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
				{
					ID:         "acffa8b8-5624-4160-8daf-77d53e89ddb2",
					PacientID:  "47322c6f-5883-4596-a305-29be7395ddd1",
					CreatedAt:  suite.today,
					UpdatedAt:  suite.today,
					Form:       valueobject.EmergencyForm{},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Cancelled_EmergencyStatus,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.ListEmergenciesByUser(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_UpdateEmergencyForm() {
	type args struct {
		userID        string
		emergencyID   string
		emergencyForm valueobject.EmergencyForm
	}

	_true := true

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
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
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				userID:      "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the user is a Pacient but the Emergency is not related to him, an error should be returned",
			args: args{
				userID:      "fbff91e2-b65a-478d-bda6-a1699a4bcecf",
				emergencyID: "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the userID provided is related to a user but it is not a Pacient or a Rescuer, an error should be returned",
			args: args{
				userID: "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the emergency contact data provided is not valid, an error should be returned",
			args: args{
				userID:      "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				emergencyID: "47322c6f-5883-4596-a305-29be7395ddd1",
				emergencyForm: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.Undefined_HeadacheIntensity,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If all input data are valid, the Emergency.Form should be updated",
			args: args{
				userID:      "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
				emergencyID: "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
				emergencyForm: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.Medium_HeadacheIntensity,
					},
				},
			},
			want: &entity.Emergency{
				ID:        "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
				PacientID: "47322c6f-5883-4596-a305-29be7395ddd1",
				CreatedAt: suite.today,
				UpdatedAt: suite.today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.Medium_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.UpdateEmergencyForm(tt.args.userID, tt.args.emergencyID, tt.args.emergencyForm)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_UpdateEmergencyStatus() {
	type args struct {
		emergencyID string
		status      valueobject.EmergencyStatus
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
		wantErr bool
	}{
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status cannot be achieved from the current EmergencyStatus, an error should be returned",
			args: args{
				emergencyID: "d08de3ca-e09f-4ba9-bee1-e0c12c1c560c",
				status:      valueobject.Finished_EmergencyStatus,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status can be achieved from the current EmergencyStatus, the Emergency.Status should be updated",
			args: args{
				emergencyID: "0526c40e-045a-48f6-895f-4f5030ba797d",
				status:      valueobject.AmbulanceToHospital_EmergencyStatus,
			},
			want: &entity.Emergency{
				Status: valueobject.AmbulanceToHospital_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.UpdateEmergencyStatus(tt.args.emergencyID, tt.args.status)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_SendAmbulance() {
	type args struct {
		userID      string
		emergencyID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
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
				userID: "9b55f726-e47c-4371-b7bd-3ce6e622b147",
			},
			wantErr: true,
		},
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status cannot be achieved from the current EmergencyStatus, an error should be returned",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "4373dae2-cec1-42f5-9080-37528ef21245",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status can be achieved from the current EmergencyStatus, the Emergency.Status should be updated",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "70d3ca7c-9491-4aa2-8aaa-507517b782c2",
			},
			want: &entity.Emergency{
				Status: valueobject.AmbulanceToPacient_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.SendAmbulance(tt.args.userID, tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_RemovePacient() {
	type args struct {
		userID      string
		emergencyID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
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
				userID: "f24d25de-4531-40e5-aef8-eedc01d1f103",
			},
			wantErr: true,
		},
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				userID:      "9b55f726-e47c-4371-b7bd-3ce6e622b147",
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status cannot be achieved from the current EmergencyStatus, an error should be returned",
			args: args{
				userID:      "9b55f726-e47c-4371-b7bd-3ce6e622b147",
				emergencyID: "4373dae2-cec1-42f5-9080-37528ef21245",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status can be achieved from the current EmergencyStatus, the Emergency.Status should be updated",
			args: args{
				userID:      "9b55f726-e47c-4371-b7bd-3ce6e622b147",
				emergencyID: "44639d78-4501-4d9b-aeac-59295e636b82",
			},
			want: &entity.Emergency{
				Status: valueobject.AmbulanceToHospital_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.RemovePacient(tt.args.userID, tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NoError(t, err)
			}
			if err == nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_FinishEmergencyCare() {
	type args struct {
		userID      string
		emergencyID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
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
				userID: "196f8233-4d02-44d3-a4a5-1c4a89b17b66",
			},
			wantErr: true,
		},
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status cannot be achieved from the current EmergencyStatus, an error should be returned",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "4373dae2-cec1-42f5-9080-37528ef21245",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status can be achieved from the current EmergencyStatus, the Emergency.Status should be updated",
			args: args{
				userID:      "be13d9a6-1e30-4fa4-9bd8-26c5c46279c2",
				emergencyID: "8ffa7ebf-7dc7-4bff-8fba-bba0661c9a51",
			},
			want: &entity.Emergency{
				Status: valueobject.Finished_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.FinishEmergencyCare(tt.args.userID, tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NoError(t, err)
			}
			if err == nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_CancelEmergency() {
	type args struct {
		emergencyID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
		wantErr bool
	}{
		{
			name: "If the emergencyID provided is not related to any emergencies in the emergencies repository, an error should be returned",
			args: args{
				emergencyID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status cannot be achieved from the current EmergencyStatus, an error should be returned",
			args: args{
				emergencyID: "4373dae2-cec1-42f5-9080-37528ef21245",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the new status can be achieved from the current EmergencyStatus, the Emergency.Status should be updated",
			args: args{
				emergencyID: "ff23c286-5321-47a9-973d-b34c4b1ff511",
			},
			want: &entity.Emergency{
				Status: valueobject.Cancelled_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.CancelEmergency(tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func (suite *EmergencyServiceTestSuite) TestEmergencyService_CreateEmergency() {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Emergency
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
			name: "If the user is a Pacient but still has no entity.Pacient data in the repository, an error should be returned",
			args: args{
				userID: "fbff91e2-b65a-478d-bda6-a1699a4bcecf",
			},
			wantErr: true,
		},
		{
			name: "If the user is a Pacient and also has entity.Pacient data in the repository, a new Emergency should be created",
			args: args{
				userID: "0ae23a9d-c9f0-4088-8e64-3ad341c07821",
			},
			want: &entity.Emergency{
				PacientID:  "47322c6f-5883-4596-a305-29be7395ddd1",
				Form:       valueobject.EmergencyForm{},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			got, err := suite.emergencyService.CreateEmergency(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, tt.want.PacientID, got.PacientID)
				assert.Equal(t, tt.want.Form, got.Form)
				assert.Equal(t, tt.want.StatusFlow, got.StatusFlow)
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}
