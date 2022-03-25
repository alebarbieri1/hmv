package presenter

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

	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "Given a User, a valid presentation should be returned",
			args: args{
				e: &entity.User{
					ID:          "1",
					Username:    "foo",
					Password:    "bar",
					ProfileKind: valueobject.Undefined_ProfileKind,
					CreatedAt:   time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &User{
				ID:          "1",
				ProfileKind: "undefined",
				CreatedAt:   "25/01/2022 - 00:00:00h",
				UpdatedAt:   "25/01/2022 - 00:00:00h",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUser(tt.args.e))
		})
	}
}

func TestNewUsers(t *testing.T) {
	type args struct {
		es []*entity.User
	}

	tests := []struct {
		name string
		args args
		want []*User
	}{
		{
			name: "Given a set of Users, a valid presentation should be returned",
			args: args{
				es: []*entity.User{
					{
						ID:          "1",
						Username:    "foo",
						Password:    "bar",
						ProfileKind: valueobject.Undefined_ProfileKind,
						CreatedAt:   time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:          "2",
						Username:    "foo2",
						Password:    "bar2",
						ProfileKind: valueobject.Pacient_ProfileKind,
						CreatedAt:   time.Date(2022, time.February, 22, 1, 2, 3, 4, time.UTC),
						UpdatedAt:   time.Date(2022, time.March, 23, 4, 3, 2, 1, time.UTC),
					},
				},
			},
			want: []*User{
				{
					ID:          "1",
					ProfileKind: "undefined",
					CreatedAt:   "25/01/2022 - 00:00:00h",
					UpdatedAt:   "25/01/2022 - 00:00:00h",
				},
				{
					ID:          "2",
					ProfileKind: "pacient",
					CreatedAt:   "22/02/2022 - 01:02:03h",
					UpdatedAt:   "23/03/2022 - 04:03:02h",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUsers(tt.args.es))
		})
	}
}
