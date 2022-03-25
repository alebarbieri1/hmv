package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEmergency(t *testing.T) {
	type args struct {
		e *entity.Emergency
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name string
		args args
		want *Emergency
	}{
		{
			name: "Given an entity.Emergency, a new Emergency should be created",
			args: args{
				e: &entity.Emergency{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
			want: &Emergency{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEmergency(tt.args.e))
		})
	}
}

func TestEmergency_toEntity(t *testing.T) {
	type fields struct {
		ID         string
		PacientID  string
		CreatedAt  time.Time
		UpdatedAt  time.Time
		Form       valueobject.EmergencyForm
		StatusFlow valueobject.EmergencyStatusFlow
		Status     valueobject.EmergencyStatus
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name   string
		fields fields
		want   *entity.Emergency
	}{
		{
			name: "Given an entity.Emergency, a new Emergency should be created",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
			want: &entity.Emergency{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emergency := &Emergency{
				ID:         tt.fields.ID,
				PacientID:  tt.fields.PacientID,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
				Form:       tt.fields.Form,
				StatusFlow: tt.fields.StatusFlow,
				Status:     tt.fields.Status,
			}

			assert.Equal(t, tt.want, emergency.toEntity())
		})
	}
}

func TestNewEmergenciesRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *EmergenciesRepository
		wantErr bool
	}{
		{
			name: "If I call NewEmergenciesRepository, a new EmergenciesRepository should be created",
			want: &EmergenciesRepository{
				emergencies: make(map[string]*Emergency),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmergenciesRepository()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestEmergenciesRepository_CreateEmergency(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	type args struct {
		emergency *entity.Emergency
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantEmergencys map[string]*Emergency
	}{
		{
			name: "Given an entity.Emergency that has not been added to the repository yet, a new Emergency should be added to the repository",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergency: &entity.Emergency{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					PacientID: "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
			wantErr: false,
			wantEmergencys: map[string]*Emergency{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
				"ee7f37e4-c165-4a35-9109-41ced42ee1fc": {
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					PacientID: "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
		{
			name: "Given an entity.Emergency that has already been added to the repository yet, an error should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergency: &entity.Emergency{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
			wantErr: true,
			wantEmergencys: map[string]*Emergency{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			err := r.CreateEmergency(tt.args.emergency)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantEmergencys, r.emergencies)
		})
	}
}

func TestEmergenciesRepository_FindEmergencyByID(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	type args struct {
		emergencyID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.Emergency
	}{
		{
			name: "Given an emergencyID that matches an Emergency that has been added to the repository, its entity.Emergency should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergencyID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want: &entity.Emergency{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
		},
		{
			name: "Given an emergencyID unrelated to any Emergencys that has been added to the repository, an error should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergencyID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			emergency, err := r.FindEmergencyByID(tt.args.emergencyID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, emergency)
		})
	}
}

func TestEmergenciesRepository_ListEmergencies(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    []*entity.Emergency
	}{
		{
			name: "If I call ListEmergencies, all Emergency entities in the repository should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			wantErr: false,
			want: []*entity.Emergency{
				{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			emergencies, err := r.ListEmergencies()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, emergencies)
		})
	}
}

func TestEmergenciesRepository_ListEmergenciesByStatus(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	type args struct {
		status valueobject.EmergencyStatus
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []*entity.Emergency
	}{
		{
			name: "Given a status related to some Emergency entities in the repository, all of them should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				status: valueobject.Undefined_EmergencyStatus,
			},
			wantErr: false,
			want: []*entity.Emergency{
				{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
		{
			name: "Given a status unrelated to any Emergency entities in the repository, no entities should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				status: valueobject.AmbulanceToHospital_EmergencyStatus,
			},
			wantErr: false,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			emergencies, err := r.ListEmergenciesByStatus(tt.args.status)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, emergencies)
		})
	}
}

func TestEmergenciesRepository_ListEmergenciesByPacientID(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	type args struct {
		pacientID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []*entity.Emergency
	}{
		{
			name: "Given a pacientID related to some Emergency entities in the repository, all of them should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				pacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: false,
			want: []*entity.Emergency{
				{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
		{
			name: "Given a pacientID unrelated to any Emergency entities in the repository, no entities should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				pacientID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			emergencies, err := r.ListEmergenciesByPacientID(tt.args.pacientID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.ElementsMatch(t, tt.want, emergencies)
		})
	}
}

func TestEmergenciesRepository_UpdateEmergency(t *testing.T) {
	type fields struct {
		emergencies map[string]*Emergency
	}

	type args struct {
		emergency *entity.Emergency
	}

	today := time.Now().Truncate(24 * time.Hour)

	_true := true

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         bool
		wantEmergencies map[string]*Emergency
	}{
		{
			name: "Given an entity.Emergency that has already been added to the repository, it should be updated with the input value",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergency: &entity.Emergency{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.Medium_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
			},
			wantErr: false,
			wantEmergencies: map[string]*Emergency{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.Medium_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
			},
		},
		{
			name: "Given an entity.Emergency that has not been added to the repository yet, an error should be returned",
			fields: fields{
				emergencies: map[string]*Emergency{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						Form: valueobject.EmergencyForm{
							Headache: valueobject.HeadacheEmergencyFormSession{
								Has:       &_true,
								Intensity: valueobject.VeryHigh_HeadacheIntensity,
							},
						},
						StatusFlow: valueobject.DefaultEmergencyStatusFlow,
						Status:     valueobject.Undefined_EmergencyStatus,
					},
				},
			},
			args: args{
				emergency: &entity.Emergency{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					PacientID: "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.Medium_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Triage_EmergencyStatus,
				},
			},
			wantErr: true,
			wantEmergencies: map[string]*Emergency{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					PacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					Form: valueobject.EmergencyForm{
						Headache: valueobject.HeadacheEmergencyFormSession{
							Has:       &_true,
							Intensity: valueobject.VeryHigh_HeadacheIntensity,
						},
					},
					StatusFlow: valueobject.DefaultEmergencyStatusFlow,
					Status:     valueobject.Undefined_EmergencyStatus,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EmergenciesRepository{
				emergencies: tt.fields.emergencies,
			}

			err := r.UpdateEmergency(tt.args.emergency)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantEmergencies, r.emergencies)
		})
	}
}
