package entity

import (
	"context"
	"flavioltonon/hmv/domain/valueobject"
	internalContext "flavioltonon/hmv/infrastructure/context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Given a valid username and password, a new User should be created",
			args: args{
				username: "foo",
				password: "bar",
			},
			want: &User{
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			wantErr: false,
		},
		{
			name: "Given an invalid username, an error should be returned",
			args: args{
				password: "bar",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Given an invalid password, an error should be returned",
			args: args{
				username: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username, tt.args.password)
			assert.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				assert.Equal(t, tt.want.Username, got.Username)
				assert.Equal(t, tt.want.Password, got.Password)
				assert.Equal(t, tt.want.ProfileKind, got.ProfileKind)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	type fields struct {
		ID          string
		Username    string
		Password    string
		ProfileKind valueobject.ProfileKind
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	var (
		today    = time.Now().Truncate(24 * time.Hour)
		tomorrow = today.Add(24 * time.Hour)

		a65 = strings.Repeat("a", 65)
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If all fields in User are valid, User.Validate() should return no errors",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: false,
		},
		{
			name: "If ID is empty, User.Validate() should return an error",
			fields: fields{
				ID:          "",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If ID contains a non-UUIDv4 value, User.Validate() should return an error",
			fields: fields{
				ID:          "baz",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If Username is empty, User.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If Username has more than 64 characters, User.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    a65,
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If Password is empty, User.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If Password has more than 64 characters, User.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    a65,
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If ProfileKind is empty, User.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: "",
				CreatedAt:   today,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If CreatedAt contains data in the future, Rescuer.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   tomorrow,
				UpdatedAt:   today,
			},
			wantErr: true,
		},
		{
			name: "If UpdatedAt contains data in the future, Rescuer.Validate() should return an error",
			fields: fields{
				ID:          "6fe98880-b181-4c1a-a17e-b6947af7f1c6",
				Username:    "foo",
				Password:    "bar",
				ProfileKind: valueobject.Analyst_ProfileKind,
				CreatedAt:   today,
				UpdatedAt:   tomorrow,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &User{
				ID:          tt.fields.ID,
				Username:    tt.fields.Username,
				Password:    tt.fields.Password,
				ProfileKind: tt.fields.ProfileKind,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}

			err := p.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestNewUserFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	today := time.Now().Truncate(24 * time.Hour)
	user := &User{
		ID:          uuid.NewString(),
		Username:    "foo",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
		CreatedAt:   today,
		UpdatedAt:   today,
	}
	baseContext := context.Background()
	validContext := internalContext.New(baseContext)
	validContext.Add(internalContext.UserKey, user)
	validContextWithNoUserData := internalContext.New(baseContext)

	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Given a valid context with user data, a new User should be returned",
			args: args{
				ctx: validContext,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Given a valid context with no user data, an error should be returned",
			args: args{
				ctx: validContextWithNoUserData,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Given an invalid context, an error should be returned",
			args: args{
				ctx: baseContext,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserFromContext(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestNewUserFromRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}

	today := time.Now().Truncate(24 * time.Hour)
	user := &User{
		ID:          uuid.NewString(),
		Username:    "foo",
		Password:    "bar",
		ProfileKind: valueobject.Pacient_ProfileKind,
		CreatedAt:   today,
		UpdatedAt:   today,
	}
	baseContext := context.Background()
	validContext := internalContext.New(baseContext)
	validContext.Add(internalContext.UserKey, user)
	validContextWithNoUserData := internalContext.New(baseContext)

	baseRequest, _ := http.NewRequest(http.MethodGet, "", nil)

	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Given a request with a valid context, a new User should be returned",
			args: args{
				r: baseRequest.WithContext(validContext),
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Given a valid context with no user data, an error should be returned",
			args: args{
				r: baseRequest.WithContext(validContextWithNoUserData),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Given an invalid context, an error should be returned",
			args: args{
				r: baseRequest.WithContext(baseContext),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserFromRequest(tt.args.r)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestUser_SetProfileKind(t *testing.T) {
	type fields struct {
		ProfileKind valueobject.ProfileKind
	}

	type args struct {
		profileKind valueobject.ProfileKind
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    valueobject.ProfileKind
	}{
		{
			name: "If User.ProfileKind is valueobject.Undefined_ProfileKind, User.SetProfileKind should return no errors",
			fields: fields{
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			args: args{
				profileKind: valueobject.Pacient_ProfileKind,
			},
			wantErr: false,
			want:    valueobject.Pacient_ProfileKind,
		},
		{
			name: "If User.ProfileKind is not valueobject.Undefined_ProfileKind, User.SetProfileKind should return an error",
			fields: fields{
				ProfileKind: valueobject.Pacient_ProfileKind,
			},
			args: args{
				profileKind: valueobject.Analyst_ProfileKind,
			},
			wantErr: true,
			want:    valueobject.Pacient_ProfileKind,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ProfileKind: tt.fields.ProfileKind,
			}

			err := u.SetProfileKind(tt.args.profileKind)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, u.ProfileKind)
		})
	}
}

func TestUser_HasProfileKind(t *testing.T) {
	type fields struct {
		ProfileKind valueobject.ProfileKind
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If User.ProfileKind is valueobject.Undefined_ProfileKind, User.HasProfileKind should return false",
			fields: fields{
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			want: false,
		},
		{
			name: "If User.ProfileKind is not valueobject.Undefined_ProfileKind, User.HasProfileKind should return true",
			fields: fields{
				ProfileKind: valueobject.Pacient_ProfileKind,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ProfileKind: tt.fields.ProfileKind,
			}

			assert.Equal(t, tt.want, u.HasProfileKind())
		})
	}
}

func TestUser_IsAnalyst(t *testing.T) {
	type fields struct {
		ProfileKind valueobject.ProfileKind
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If User.ProfileKind is valueobject.Analyst_ProfileKind, User.IsAnalyst should return true",
			fields: fields{
				ProfileKind: valueobject.Analyst_ProfileKind,
			},
			want: true,
		},
		{
			name: "If User.ProfileKind is not valueobject.Analyst_ProfileKind, User.IsAnalyst should return false",
			fields: fields{
				ProfileKind: valueobject.Pacient_ProfileKind,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ProfileKind: tt.fields.ProfileKind,
			}

			assert.Equal(t, tt.want, u.IsAnalyst())
		})
	}
}

func TestNewPacient(t *testing.T) {
	type fields struct {
		ID          string
		Data        valueobject.UserData
		ProfileKind valueobject.ProfileKind
	}

	userID := uuid.NewString()

	tests := []struct {
		name    string
		fields  fields
		want    *Pacient
		wantErr bool
	}{
		{
			name: "If the User.ID is not set, an error should be returned",
			fields: fields{
				ID:          "",
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "If the User has an ID, a new Pacient should be created",
			fields: fields{
				ID: userID,
				Data: valueobject.UserData{
					Name: "foo",
				},
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			want: &Pacient{
				UserID: userID,
				Data: valueobject.PacientData{
					Name: "foo",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				ID:          tt.fields.ID,
				Data:        tt.fields.Data,
				ProfileKind: tt.fields.ProfileKind,
			}

			got, err := user.NewPacient()
			assert.Equal(t, tt.wantErr, err != nil)

			if err == nil {
				assert.Equal(t, tt.want.UserID, got.UserID)
				assert.Equal(t, tt.want.Data, got.Data)
			}
		})
	}
}

func TestUser_IsPacient(t *testing.T) {
	type fields struct {
		ProfileKind valueobject.ProfileKind
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If User.ProfileKind is valueobject.Pacient_ProfileKind, User.IsPacient should return true",
			fields: fields{
				ProfileKind: valueobject.Pacient_ProfileKind,
			},
			want: true,
		},
		{
			name: "If User.ProfileKind is not valueobject.Pacient_ProfileKind, User.IsPacient should return false",
			fields: fields{
				ProfileKind: valueobject.Rescuer_ProfileKind,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ProfileKind: tt.fields.ProfileKind,
			}

			assert.Equal(t, tt.want, u.IsPacient())
		})
	}
}

func TestUser_IsRescuer(t *testing.T) {
	type fields struct {
		ProfileKind valueobject.ProfileKind
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If User.ProfileKind is valueobject.Rescuer_ProfileKind, User.IsRescuer should return true",
			fields: fields{
				ProfileKind: valueobject.Rescuer_ProfileKind,
			},
			want: true,
		},
		{
			name: "If User.ProfileKind is not valueobject.Rescuer_ProfileKind, User.IsRescuer should return false",
			fields: fields{
				ProfileKind: valueobject.Undefined_ProfileKind,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ProfileKind: tt.fields.ProfileKind,
			}

			assert.Equal(t, tt.want, u.IsRescuer())
		})
	}
}

func TestUser_SetUserData(t *testing.T) {
	type fields struct {
		Data valueobject.UserData
	}

	type args struct {
		data valueobject.UserData
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Given a valid valueobject.UserData, User.SetUserData should return no errors",
			fields: fields{
				Data: valueobject.UserData{
					Name: "baz",
				},
			},
			args: args{
				data: valueobject.UserData{
					Name: "qux",
				},
			},
			wantErr: false,
		},
		{
			name: "Given an invalid valueobject.UserData, User.SetUserData should return an error",
			fields: fields{
				Data: valueobject.UserData{
					Name: "baz",
				},
			},
			args: args{
				data: valueobject.UserData{
					Name: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Data: tt.fields.Data,
			}

			err := u.SetUserData(tt.args.data)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
