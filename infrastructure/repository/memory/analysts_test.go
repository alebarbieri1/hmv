package memory

import (
	"flavioltonon/hmv/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAnalyst(t *testing.T) {
	type args struct {
		e *entity.Analyst
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name string
		args args
		want *Analyst
	}{
		{
			name: "Given an entity.Analyst, a new Analyst should be created",
			args: args{
				e: &entity.Analyst{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			want: &Analyst{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewAnalyst(tt.args.e))
		})
	}
}

func TestAnalyst_toEntity(t *testing.T) {
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
		want   *entity.Analyst
	}{
		{
			name: "Given an entity.Analyst, a new Analyst should be created",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
			want: &entity.Analyst{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyst := &Analyst{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			assert.Equal(t, tt.want, analyst.toEntity())
		})
	}
}

func TestNewAnalystsRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *AnalystsRepository
		wantErr bool
	}{
		{
			name: "If I call NewAnalystsRepository, a new AnalystsRepository should be created",
			want: &AnalystsRepository{
				analysts: make(map[string]*Analyst),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAnalystsRepository()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAnalystsRepository_CreateAnalyst(t *testing.T) {
	type fields struct {
		analysts map[string]*Analyst
	}

	type args struct {
		analyst *entity.Analyst
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantAnalysts map[string]*Analyst
	}{
		{
			name: "Given an entity.Analyst that has not been added to the repository yet, a new Analyst should be added to the repository",
			fields: fields{
				analysts: map[string]*Analyst{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				analyst: &entity.Analyst{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			wantErr: false,
			wantAnalysts: map[string]*Analyst{
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
			name: "Given an entity.Analyst that has already been added to the repository yet, an error should be returned",
			fields: fields{
				analysts: map[string]*Analyst{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				analyst: &entity.Analyst{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
				},
			},
			wantErr: true,
			wantAnalysts: map[string]*Analyst{
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
			r := &AnalystsRepository{
				analysts: tt.fields.analysts,
			}

			err := r.CreateAnalyst(tt.args.analyst)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantAnalysts, r.analysts)
		})
	}
}

func TestAnalystsRepository_FindAnalystByID(t *testing.T) {
	type fields struct {
		analysts map[string]*Analyst
	}

	type args struct {
		analystID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.Analyst
	}{
		{
			name: "Given an analystID that matches an Analyst that has been added to the repository, its entity.Analyst should be returned",
			fields: fields{
				analysts: map[string]*Analyst{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				analystID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want: &entity.Analyst{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
		{
			name: "Given an analystID unrelated to any Analysts that has been added to the repository, an error should be returned",
			fields: fields{
				analysts: map[string]*Analyst{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
					},
				},
			},
			args: args{
				analystID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AnalystsRepository{
				analysts: tt.fields.analysts,
			}

			analyst, err := r.FindAnalystByID(tt.args.analystID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, analyst)
		})
	}
}

func TestAnalystsRepository_FindAnalystByUserID(t *testing.T) {
	type fields struct {
		analysts map[string]*Analyst
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
		want    *entity.Analyst
	}{
		{
			name: "Given an userID that matches an Analyst that has been added to the repository, its entity.Analyst should be returned",
			fields: fields{
				analysts: map[string]*Analyst{
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
			want: &entity.Analyst{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
			},
		},
		{
			name: "Given an userID unrelated to any Analysts that has been added to the repository, an error should be returned",
			fields: fields{
				analysts: map[string]*Analyst{
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
			r := &AnalystsRepository{
				analysts: tt.fields.analysts,
			}

			analyst, err := r.FindAnalystByUserID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, analyst)
		})
	}
}
