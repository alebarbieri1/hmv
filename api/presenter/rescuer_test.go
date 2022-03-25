package presenter

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

	tests := []struct {
		name string
		args args
		want *Rescuer
	}{
		{
			name: "Given a Rescuer, a valid presentation should be returned",
			args: args{
				e: &entity.Rescuer{
					ID:        "foo",
					UserID:    "bar",
					CreatedAt: time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, time.January, 25, 0, 0, 0, 0, time.UTC),
				},
			},
			want: &Rescuer{
				ID:        "foo",
				UserID:    "bar",
				CreatedAt: "25/01/2022 - 00:00:00h",
				UpdatedAt: "25/01/2022 - 00:00:00h",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRescuer(tt.args.e))
		})
	}
}
