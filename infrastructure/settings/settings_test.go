package settings

import (
	"flavioltonon/hmv/infrastructure/logging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *Settings
		wantErr bool
	}{
		{
			name: "If I call New, a new Settings with its default values should be created",
			want: &Settings{
				Server: &ServerSettings{
					Address:         ":8080",
					DevelopmentMode: true,
				},
				Logging: &logging.Settings{
					DevelopmentMode: true,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
