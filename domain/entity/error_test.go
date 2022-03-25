package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrProfileKindAlreadySet(t *testing.T) {
	type args struct {
		profileKind valueobject.ProfileKind
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Given a valid valueobject.ProfileKind, ErrProfileKindAlreadySet should return an error message accordingly",
			args: args{
				profileKind: valueobject.Analyst_ProfileKind,
			},
			want: "profile kind already set as analyst",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ErrProfileKindAlreadySet(tt.args.profileKind).Error())
		})
	}
}

func TestErrInvalidStatusChange(t *testing.T) {
	type args struct {
		from valueobject.EmergencyStatus
		to   valueobject.EmergencyStatus
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Given valid from and to valueobject.EmergencyStatus, ErrInvalidStatusChange should return an error message accordingly",
			args: args{
				from: valueobject.Triage_EmergencyStatus,
				to:   valueobject.Finished_EmergencyStatus,
			},
			want: "status cannot be changed from triage to finished",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ErrInvalidStatusChange(tt.args.from, tt.args.to).Error())
		})
	}
}
