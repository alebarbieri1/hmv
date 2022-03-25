package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEmergency(t *testing.T) {
	type args struct {
		pacientID string
	}

	tests := []struct {
		name    string
		args    args
		want    *Emergency
		wantErr bool
	}{
		{
			name: "Given a valid pacientID, a new Analyst should be created",
			args: args{
				pacientID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			want: &Emergency{
				PacientID:  "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: false,
		},
		{
			name: "Given an invalid pacientID, an error should be returned",
			args: args{
				pacientID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmergency(tt.args.pacientID)
			assert.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				assert.Equal(t, tt.want.PacientID, got.PacientID)
				assert.Equal(t, tt.want.StatusFlow, got.StatusFlow)
				assert.Equal(t, tt.want.Status, got.Status)
			}
		})
	}
}

func TestEmergency_Validate(t *testing.T) {
	type fields struct {
		ID         string
		PacientID  string
		CreatedAt  time.Time
		UpdatedAt  time.Time
		StatusFlow valueobject.EmergencyStatusFlow
		Status     valueobject.EmergencyStatus
	}

	var (
		today    = time.Now().Truncate(24 * time.Hour)
		tomorrow = today.Add(24 * time.Hour)
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all fields in Emergency are valid, Emergency.Validate() should return no errors",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: false,
		},
		{
			name: "If ID is empty, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If ID contains a non-UUIDv4 value, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "foo",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If PacientID is empty, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If PacientID contains a non-UUIDv4 value, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "foo",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If CreatedAt contains data in the future, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  tomorrow,
				UpdatedAt:  today,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If UpdatedAt contains data in the future, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  tomorrow,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If StatusFlow is empty, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  today,
				StatusFlow: nil,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			wantErr: true,
		},
		{
			name: "If Status is empty, Emergency.Validate() should return an error",
			fields: fields{
				ID:         "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				PacientID:  "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt:  today,
				UpdatedAt:  tomorrow,
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emergency{
				ID:         tt.fields.ID,
				PacientID:  tt.fields.PacientID,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
				StatusFlow: tt.fields.StatusFlow,
				Status:     tt.fields.Status,
			}

			err := e.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestEmergency_UpdateForm(t *testing.T) {
	type fields struct {
		Form valueobject.EmergencyForm
	}

	type args struct {
		form valueobject.EmergencyForm
	}

	var (
		_true  = true
		_false = false
	)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Emergency
	}{
		{
			name: "Given a valid EmergencyForm, Emergency.Form should be updated",
			fields: fields{
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_false,
						Intensity: valueobject.Undefined_HeadacheIntensity,
					},
				},
			},
			args: args{
				form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
			},
			wantErr: false,
			want: &Emergency{
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
			},
		},
		{
			name: "Given an invalid EmergencyForm, Emergency.Form an error should be returned",
			fields: fields{
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
			},
			args: args{
				form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_false,
						Intensity: valueobject.Medium_HeadacheIntensity,
					},
				},
			},
			wantErr: true,
			want: &Emergency{
				Form: valueobject.EmergencyForm{
					Headache: valueobject.HeadacheEmergencyFormSession{
						Has:       &_true,
						Intensity: valueobject.VeryHigh_HeadacheIntensity,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emergency{
				Form: tt.fields.Form,
			}

			err := e.UpdateForm(tt.args.form)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, e)
		})
	}
}

func TestEmergency_UpdateStatus(t *testing.T) {
	type fields struct {
		StatusFlow valueobject.EmergencyStatusFlow
		Status     valueobject.EmergencyStatus
	}

	type args struct {
		status valueobject.EmergencyStatus
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantStatus valueobject.EmergencyStatus
	}{
		{
			name: "Given a status change that is mapped in the Emergency.StatusFlow, Emergency.Status should be updated",
			fields: fields{
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Undefined_EmergencyStatus,
			},
			args: args{
				status: valueobject.Triage_EmergencyStatus,
			},
			wantErr:    false,
			wantStatus: valueobject.Triage_EmergencyStatus,
		},
		{
			name: "Given a status change that is not mapped in the Emergency.StatusFlow, an error should be returned",
			fields: fields{
				StatusFlow: valueobject.DefaultEmergencyStatusFlow,
				Status:     valueobject.Triage_EmergencyStatus,
			},
			args: args{
				status: valueobject.Undefined_EmergencyStatus,
			},
			wantErr:    true,
			wantStatus: valueobject.Triage_EmergencyStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emergency{
				StatusFlow: tt.fields.StatusFlow,
				Status:     tt.fields.Status,
			}

			err := e.UpdateStatus(tt.args.status)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantStatus, e.Status)
		})
	}
}

func TestEmergency_Priority(t *testing.T) {
	type fields struct {
		Form valueobject.EmergencyForm
	}

	_true := true

	tests := []struct {
		name   string
		fields fields
		want   valueobject.EmergencyPriority
	}{
		{
			name: "Emergency.Priority() should should return the priority of the Emergency based on its Emergency.Form",
			fields: fields{
				Form: valueobject.EmergencyForm{
					ChestPain: valueobject.ChestPainEmergencyFormSession{
						Has:             &_true,
						Characteristics: valueobject.RadiatingToTheLeftArm_ChestPainCharacteristics,
					},
				},
			},
			want: valueobject.VeryHigh_EmergencyPriority,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emergency{
				Form: tt.fields.Form,
			}

			assert.Equal(t, tt.want, e.Priority())
		})
	}
}
