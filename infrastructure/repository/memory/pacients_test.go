package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPacient(t *testing.T) {
	type args struct {
		e *entity.Pacient
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name string
		args args
		want *Pacient
	}{
		{
			name: "Given an entity.Pacient, a new Pacient should be created",
			args: args{
				e: &entity.Pacient{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
			want: &Pacient{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPacient(tt.args.e))
		})
	}
}

func TestPacient_toEntity(t *testing.T) {
	type fields struct {
		ID               string
		UserID           string
		EmergencyContact valueobject.EmergencyContact
		CreatedAt        time.Time
		UpdatedAt        time.Time
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name   string
		fields fields
		want   *entity.Pacient
	}{
		{
			name: "Given an entity.Pacient, a new Pacient should be created",
			fields: fields{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
			want: &entity.Pacient{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pacient := &Pacient{
				ID:               tt.fields.ID,
				UserID:           tt.fields.UserID,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
				EmergencyContact: tt.fields.EmergencyContact,
			}

			assert.Equal(t, tt.want, pacient.toEntity())
		})
	}
}

func TestNewPacientsRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *PacientsRepository
		wantErr bool
	}{
		{
			name: "If I call NewPacientsRepository, a new PacientsRepository should be created",
			want: &PacientsRepository{
				pacients: make(map[string]*Pacient),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPacientsRepository())
		})
	}
}

func TestPacientsRepository_CreatePacient(t *testing.T) {
	type fields struct {
		pacients map[string]*Pacient
	}

	type args struct {
		pacient *entity.Pacient
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantPacients map[string]*Pacient
	}{
		{
			name: "Given an entity.Pacient that has not been added to the repository yet, a new Pacient should be added to the repository",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacient: &entity.Pacient{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
			wantErr: false,
			wantPacients: map[string]*Pacient{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
				"ee7f37e4-c165-4a35-9109-41ced42ee1fc": {
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
		},
		{
			name: "Given an entity.Pacient that has already been added to the repository yet, an error should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacient: &entity.Pacient{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
			wantErr: true,
			wantPacients: map[string]*Pacient{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PacientsRepository{
				pacients: tt.fields.pacients,
			}

			err := r.CreatePacient(tt.args.pacient)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPacients, r.pacients)
		})
	}
}

func TestPacientsRepository_FindPacientByID(t *testing.T) {
	type fields struct {
		pacients map[string]*Pacient
	}

	type args struct {
		pacientID string
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.Pacient
	}{
		{
			name: "Given an pacientID that matches an Pacient that has been added to the repository, its entity.Pacient should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacientID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want: &entity.Pacient{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
		},
		{
			name: "Given an pacientID unrelated to any Pacients that has been added to the repository, an error should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacientID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PacientsRepository{
				pacients: tt.fields.pacients,
			}

			pacient, err := r.FindPacientByID(tt.args.pacientID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, pacient)
		})
	}
}

func TestPacientsRepository_FindPacientByUserID(t *testing.T) {
	type fields struct {
		pacients map[string]*Pacient
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
		want    *entity.Pacient
	}{
		{
			name: "Given an userID that matches an Pacient that has been added to the repository, its entity.Pacient should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				userID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: false,
			want: &entity.Pacient{
				ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
				CreatedAt: today,
				UpdatedAt: today,
				EmergencyContact: valueobject.EmergencyContact{
					Name:         "foo",
					MobileNumber: "5511999999999",
				},
			},
		},
		{
			name: "Given an userID unrelated to any Pacients that has been added to the repository, an error should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
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
			r := &PacientsRepository{
				pacients: tt.fields.pacients,
			}

			pacient, err := r.FindPacientByUserID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, pacient)
		})
	}
}

func TestPacientsRepository_UpdatePacient(t *testing.T) {
	type fields struct {
		pacients map[string]*Pacient
	}

	type args struct {
		pacient *entity.Pacient
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		wantPacients map[string]*Pacient
	}{
		{
			name: "Given an entity.Pacient that has already been added to the repository, it should be updated with the input value",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacient: &entity.Pacient{
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "bar",
						MobileNumber: "5519999999999",
					},
				},
			},
			wantErr: false,
			wantPacients: map[string]*Pacient{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "bar",
						MobileNumber: "5519999999999",
					},
				},
			},
		},
		{
			name: "Given an entity.Pacient that has not been added to the repository yet, an error should be returned",
			fields: fields{
				pacients: map[string]*Pacient{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
						CreatedAt: today,
						UpdatedAt: today,
						EmergencyContact: valueobject.EmergencyContact{
							Name:         "foo",
							MobileNumber: "5511999999999",
						},
					},
				},
			},
			args: args{
				pacient: &entity.Pacient{
					ID:        "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					UserID:    "21a48c8f-1008-4e11-a05f-82fabab09b92",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "",
						MobileNumber: "5519999999999999999999999",
					},
				},
			},
			wantErr: true,
			wantPacients: map[string]*Pacient{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:        "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					UserID:    "6453415b-ea7f-4519-bb55-0f66bc50621b",
					CreatedAt: today,
					UpdatedAt: today,
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "foo",
						MobileNumber: "5511999999999",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PacientsRepository{
				pacients: tt.fields.pacients,
			}

			err := r.UpdatePacient(tt.args.pacient)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPacients, r.pacients)
		})
	}
}
