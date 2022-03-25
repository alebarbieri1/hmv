package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type args struct {
		e *entity.User
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "Given an entity.User, a new User should be created",
			args: args{
				e: &entity.User{
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Pacient_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
			want: &User{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Pacient_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUser(tt.args.e))
		})
	}
}

func TestUser_toEntity(t *testing.T) {
	type fields struct {
		ID          string
		Username    string
		Password    string
		ProfileKind valueobject.ProfileKind
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name   string
		fields fields
		want   *entity.User
	}{
		{
			name: "Given an entity.User, a new User should be created",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			want: &entity.User{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				ID:          tt.fields.ID,
				Username:    tt.fields.Username,
				Password:    tt.fields.Password,
				ProfileKind: tt.fields.ProfileKind,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}

			assert.Equal(t, tt.want, user.toEntity())
		})
	}
}

func TestNewUsersRepository(t *testing.T) {
	tests := []struct {
		name    string
		want    *UsersRepository
		wantErr bool
	}{
		{
			name: "If I call NewUsersRepository, a new UsersRepository should be created",
			want: &UsersRepository{
				users: make(map[string]*User),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUsersRepository()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUsersRepository_CreateUser(t *testing.T) {
	type fields struct {
		users map[string]*User
	}

	type args struct {
		user *entity.User
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantUsers map[string]*User
	}{
		{
			name: "Given an entity.User that has not been added to the repository yet, a new User should be added to the repository",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Rescuer_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				user: &entity.User{
					ID:          "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Rescuer_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
			wantErr: false,
			wantUsers: map[string]*User{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Rescuer_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
				"ee7f37e4-c165-4a35-9109-41ced42ee1fc": {
					ID:          "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Rescuer_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
		},
		{
			name: "Given an entity.User that has already been added to the repository yet, an error should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Rescuer_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				user: &entity.User{
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Rescuer_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
			wantErr: true,
			wantUsers: map[string]*User{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Rescuer_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepository{
				users: tt.fields.users,
			}

			err := r.CreateUser(tt.args.user)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantUsers, r.users)
		})
	}
}

func TestUsersRepository_FindUserByID(t *testing.T) {
	type fields struct {
		users map[string]*User
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
		want    *entity.User
	}{
		{
			name: "Given an userID that matches an User that has been added to the repository, its entity.User should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				userID: "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
			},
			wantErr: false,
			want: &entity.User{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Undefined_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
		},
		{
			name: "Given an userID unrelated to any Users that has been added to the repository, an error should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				userID: "6453415b-ea7f-4519-bb55-0f66bc50621b",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepository{
				users: tt.fields.users,
			}

			user, err := r.FindUserByID(tt.args.userID)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, user)
		})
	}
}

func TestUsersRepository_FindUserByUsername(t *testing.T) {
	type fields struct {
		users map[string]*User
	}

	type args struct {
		username string
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *entity.User
	}{
		{
			name: "Given an username that matches an User that has been added to the repository, its entity.User should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				username: "foo",
			},
			wantErr: false,
			want: &entity.User{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Undefined_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
		},
		{
			name: "Given an username unrelated to any Users that has been added to the repository, an error should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				username: "baz",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepository{
				users: tt.fields.users,
			}

			user, err := r.FindUserByUsername(tt.args.username)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, user)
		})
	}
}

func TestEmergenciesRepository_ListUsers(t *testing.T) {
	type fields struct {
		users map[string]*User
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    []*entity.User
	}{
		{
			name: "If I call ListUsers, all User entities in the repository should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			wantErr: false,
			want: []*entity.User{
				{
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Undefined_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepository{
				users: tt.fields.users,
			}

			users, err := r.ListUsers()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, users)
		})
	}
}

func TestUsersRepository_UpdateUser(t *testing.T) {
	type fields struct {
		users map[string]*User
	}

	type args struct {
		user *entity.User
	}

	today := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantUsers map[string]*User
	}{
		{
			name: "Given an entity.User that has already been added to the repository, it should be updated with the input value",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				user: &entity.User{
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Pacient_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
			wantErr: false,
			wantUsers: map[string]*User{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Pacient_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
		},
		{
			name: "Given an entity.User that has not been added to the repository yet, an error should be returned",
			fields: fields{
				users: map[string]*User{
					"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
						ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   today,
						UpdatedAt:   today,
					},
				},
			},
			args: args{
				user: &entity.User{
					ID:          "ee7f37e4-c165-4a35-9109-41ced42ee1fc",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Pacient_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
			wantErr: true,
			wantUsers: map[string]*User{
				"6fe98880-b181-4c1a-a17e-b6947af7f1c6": {
					ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Undefined_ProfileKind,
					CreatedAt:   today,
					UpdatedAt:   today,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepository{
				users: tt.fields.users,
			}

			err := r.UpdateUser(tt.args.user)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantUsers, r.users)
		})
	}
}
