package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmergencyStatusFlow_CanChange(t *testing.T) {
	type args struct {
		from EmergencyStatus
		to   EmergencyStatus
	}

	tests := []struct {
		name string
		f    EmergencyStatusFlow
		args args
		want bool
	}{
		{
			name: "If an EmergencyStatus change is set as allowed in the EmergencyStatusFlow, CanChange should return true",
			f:    DefaultEmergencyStatusFlow,
			args: args{
				from: Triage_EmergencyStatus,
				to:   AmbulanceToPacient_EmergencyStatus,
			},
			want: true,
		},
		{
			name: "If an EmergencyStatus change is not set as allowed in the EmergencyStatusFlow, CanChange should return false",
			f:    DefaultEmergencyStatusFlow,
			args: args{
				from: Triage_EmergencyStatus,
				to:   Finished_EmergencyStatus,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.f.CanChange(tt.args.from, tt.args.to))
		})
	}
}
