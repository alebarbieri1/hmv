package presenter

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

	tests := []struct {
		name string
		args args
		want *Pacient
	}{
		{
			name: "Given a Pacient, a valid presentation should be returned",
			args: args{
				e: &entity.Pacient{
					ID:     "foo",
					UserID: "bar",
					EmergencyContact: valueobject.EmergencyContact{
						Name:         "baz",
						MobileNumber: "qux",
					},
					CreatedAt: time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &Pacient{
				ID:     "foo",
				UserID: "bar",
				EmergencyContact: &EmergencyContact{
					Name:         "baz",
					MobileNumber: "qux",
				},
				CreatedAt: "25/01/2022 - 00:00:00h",
				UpdatedAt: "25/01/2022 - 00:00:00h",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPacient(tt.args.e))
		})
	}
}
