package presenter

import (
	"flavioltonon/hmv/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmergencyContact(t *testing.T) {
	type args struct {
		c valueobject.EmergencyContact
	}

	tests := []struct {
		name string
		args args
		want *EmergencyContact
	}{
		{
			name: "Given a EmergencyContact, a valid presentation should be returned",
			args: args{
				c: valueobject.EmergencyContact{
					Name:         "baz",
					MobileNumber: "qux",
				},
			},
			want: &EmergencyContact{
				Name:         "baz",
				MobileNumber: "qux",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEmergencyContact(tt.args.c))
		})
	}
}
