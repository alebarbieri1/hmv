package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRescuer(t *testing.T) {
	type args struct {
		userID string
	}

	tests := []struct {
		name    string
		args    args
		want    *Rescuer
		wantErr bool
	}{
		{
			name: "Given a valid userID, a new Rescuer should be created",
			args: args{
				userID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			want: &Rescuer{
				UserID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
		},
		{
			name: "Given an invalid userID, an error should be returned",
			args: args{
				userID: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRescuer(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				assert.Equal(t, tt.want.UserID, got.UserID)
			}
		})
	}
}

func TestRescuer_Validate(t *testing.T) {
	type fields struct {
		ID        string
		UserID    string
		CreatedAt time.Time
		UpdatedAt time.Time
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
			name: "If all fields in Rescuer are valid, Rescuer.Validate() should return no errors",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: false,
		},
		{
			name: "If ID is empty, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If ID contains a non-UUIDv4 value, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "foo",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UserID is empty, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UserID contains a non-UUIDv4 value, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "foo",
				CreatedAt: today,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If CreatedAt contains data in the future, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: tomorrow,
				UpdatedAt: today,
			},
			wantErr: true,
		},
		{
			name: "If UpdatedAt contains data in the future, Rescuer.Validate() should return an error",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: tomorrow,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Rescuer{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			err := p.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
