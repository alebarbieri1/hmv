package memory

import (
	"flavioltonon/hmv/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRescuer(t *testing.T) {
	type args struct {
		e *entity.Rescuer
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name string
		args args
		want *Rescuer
	}{
		{
			name: "Given an entity.Rescuer, a new Rescuer should be created",
			args: args{
				e: &entity.Rescuer{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			want: &Rescuer{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRescuer(tt.args.e))
		})
	}
}

func TestRescuer_toEntity(t *testing.T) {
	type fields struct {
		ID        string
		UserID    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name   string
		fields fields
		want   *entity.Rescuer
	}{
		{
			name: "Given an entity.Rescuer, a new Rescuer should be created",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			want: &entity.Rescuer{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescuer := &Rescuer{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			assert.Equal(t, tt.want, rescuer.toEntity())
		})
	}
}

func TestNewRescuersRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *RescuersRepository
		wantErr bool
	}{
		{
			name: "If I call NewRescuersRepository, a new RescuersRepository should be created",
			want: &RescuersRepository{
				rescuers: make(map[string]*Rescuer),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRescuersRepository())
		})
	}
}

func TestRescuersRepository_CreateRescuer(t *testing.T) {
	type fields struct {
		rescuers map[string]*Rescuer
	}

	type args struct {
		rescuer *entity.Rescuer
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantRescuers map[string]*Rescuer
	}{
		{
			name: "Given an entity.Rescuer that has not been added to the repository yet, a new Rescuer should be added to the repository",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				rescuer: &entity.Rescuer{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			wantErr: false,
			wantRescuers: map[string]*Rescuer{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
				},
				"ee7f37e4-c165-4a35-9109-41ced42ee1fc": {
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
		},
		{
			name: "Given an entity.Rescuer that has already been added to the repository yet, an error should be returned",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				rescuer: &entity.Rescuer{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			wantErr: true,
			wantRescuers: map[string]*Rescuer{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RescuersRepository{
				rescuers: tt.fields.rescuers,
			}

			err := r.CreateRescuer(tt.args.rescuer)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantRescuers, r.rescuers)
		})
	}
}

func TestRescuersRepository_FindRescuerByID(t *testing.T) {
	type fields struct {
		rescuers map[string]*Rescuer
	}

	type args struct {
		rescuerID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.Rescuer
	}{
		{
			name: "Given an rescuerID that matches an Rescuer that has been added to the repository, its entity.Rescuer should be returned",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				rescuerID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want: &entity.Rescuer{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
		{
			name: "Given an rescuerID unrelated to any Rescuers that has been added to the repository, an error should be returned",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				rescuerID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RescuersRepository{
				rescuers: tt.fields.rescuers,
			}

			rescuer, err := r.FindRescuerByID(tt.args.rescuerID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, rescuer)
		})
	}
}

func TestRescuersRepository_FindRescuerByUserID(t *testing.T) {
	type fields struct {
		rescuers map[string]*Rescuer
	}

	type args struct {
		userID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.Rescuer
	}{
		{
			name: "Given an userID that matches an Rescuer that has been added to the repository, its entity.Rescuer should be returned",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				userID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: false,
			want: &entity.Rescuer{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
		{
			name: "Given an userID unrelated to any Rescuers that has been added to the repository, an error should be returned",
			fields: fields{
				rescuers: map[string]*Rescuer{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				userID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RescuersRepository{
				rescuers: tt.fields.rescuers,
			}

			rescuer, err := r.FindRescuerByUserID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, rescuer)
		})
	}
}
