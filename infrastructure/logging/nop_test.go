package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNopLogger(t *testing.T) {
	tests := []struct {
		name    string
		want    Logger
		wantErr bool
	}{
		{
			name:    "If I call NewNopLogger, a new NopLogger should be returned",
			want:    new(NopLogger),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNopLogger()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
