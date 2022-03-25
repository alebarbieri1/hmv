package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmergencyPriority_String(t *testing.T) {
	tests := []struct {
		name string
		i    EmergencyPriority
		want string
	}{
		{
			name: "If EmergencyPriority is Low_EmergencyPriority, EmergencyPriority.String() should return 'low''",
			i:    Low_EmergencyPriority,
			want: "low",
		},
		{
			name: "If EmergencyPriority is Medium_EmergencyPriority, EmergencyPriority.String() should return 'medium''",
			i:    Medium_EmergencyPriority,
			want: "medium",
		},
		{
			name: "If EmergencyPriority is High_EmergencyPriority, EmergencyPriority.String() should return 'high''",
			i:    High_EmergencyPriority,
			want: "high",
		},
		{
			name: "If EmergencyPriority is VeryHigh_EmergencyPriority, EmergencyPriority.String() should return 'very-high''",
			i:    VeryHigh_EmergencyPriority,
			want: "very-high",
		},
		{
			name: "If EmergencyPriority is Undefined_EmergencyPriority, EmergencyPriority.String() should return 'undefined''",
			i:    Undefined_EmergencyPriority,
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
