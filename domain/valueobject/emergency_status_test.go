package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmergencyStatus_String(t *testing.T) {
	tests := []struct {
		name string
		i    EmergencyStatus
		want string
	}{
		{
			name: "If EmergencyStatus is Triage_EmergencyStatus, EmergencyStatus.String() should return 'triage''",
			i:    Triage_EmergencyStatus,
			want: "triage",
		},
		{
			name: "If EmergencyStatus is AmbulanceToPacient_EmergencyStatus, EmergencyStatus.String() should return 'ambulance-to-pacient''",
			i:    AmbulanceToPacient_EmergencyStatus,
			want: "ambulance-to-pacient",
		},
		{
			name: "If EmergencyStatus is AmbulanceToHospital_EmergencyStatus, EmergencyStatus.String() should return 'ambulance-to-hospital''",
			i:    AmbulanceToHospital_EmergencyStatus,
			want: "ambulance-to-hospital",
		},
		{
			name: "If EmergencyStatus is Finished_EmergencyStatus, EmergencyStatus.String() should return 'finished''",
			i:    Finished_EmergencyStatus,
			want: "finished",
		},
		{
			name: "If EmergencyStatus is Cancelled_EmergencyStatus, EmergencyStatus.String() should return 'cancelled''",
			i:    Cancelled_EmergencyStatus,
			want: "cancelled",
		},
		{
			name: "If EmergencyStatus is Undefined_EmergencyStatus, EmergencyStatus.String() should return 'undefined''",
			i:    Undefined_EmergencyStatus,
			want: "undefined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.i.String())
			})
		})
	}
}

func TestNewEmergencyStatusFromString(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want EmergencyStatus
	}{
		{
			name: "Given a string with value 'triage', Triage_EmergencyStatus should be returned",
			args: args{
				s: "triage",
			},
			want: Triage_EmergencyStatus,
		},
		{
			name: "Given a string with value 'ambulance-to-pacient', AmbulanceToPacient_EmergencyStatus should be returned",
			args: args{
				s: "ambulance-to-pacient",
			},
			want: AmbulanceToPacient_EmergencyStatus,
		},
		{
			name: "Given a string with value 'ambulance-to-hospital', AmbulanceToHospital_EmergencyStatus should be returned",
			args: args{
				s: "ambulance-to-hospital",
			},
			want: AmbulanceToHospital_EmergencyStatus,
		},
		{
			name: "Given a string with value 'finished', Finished_EmergencyStatus should be returned",
			args: args{
				s: "finished",
			},
			want: Finished_EmergencyStatus,
		},
		{
			name: "Given a string with value 'cancelled', Cancelled_EmergencyStatus should be returned",
			args: args{
				s: "cancelled",
			},
			want: Cancelled_EmergencyStatus,
		},
		{
			name: "Given a string unrelated to any EmergencyStatus levels, Undefined_EmergencyStatus should be returned",
			args: args{
				s: "banana",
			},
			want: Undefined_EmergencyStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, NewEmergencyStatusFromString(tt.args.s))
			})
		})
	}
}
