package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackacheEmergencyFormSession_IsSet(t *testing.T) {
	type fields struct {
		Has *bool
	}

	_false := false

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "If Has is not set, IsSet should return false",
			fields: fields{
				Has: nil,
			},
			want: false,
		},
		{
			name: "If Has is set with any values, IsSet should return true",
			fields: fields{
				Has: &_false,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BackacheEmergencyFormSession{
				Has: tt.fields.Has,
			}

			assert.Equal(t, tt.want, f.IsSet())
		})
	}
}

func TestBackacheEmergencyFormSession_Score(t *testing.T) {
	type fields struct {
		Has *bool
	}

	var (
		_true  = true
		_false = false
	)

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "If the form session is not set, BackacheEmergencyFormSession.Score() should return 0",
			fields: fields{
				Has: nil,
			},
			want: 0,
		},
		{
			name: "If the form session is set as false, BackacheEmergencyFormSession.Score() should return 0",
			fields: fields{
				Has: &_false,
			},
			want: 0,
		},
		{
			name: "If the form session is set as true, BackacheEmergencyFormSession.Score() should return 1",
			fields: fields{
				Has: &_true,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := BackacheEmergencyFormSession{
				Has: tt.fields.Has,
			}

			assert.Equal(t, tt.want, f.Score())
		})
	}
}
