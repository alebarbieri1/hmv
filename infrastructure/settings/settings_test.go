package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFile(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *Settings
		wantErr bool
	}{
		{
			name: "If I call New, a new Settings with its default values should be created",
			args: args{
				path: "testdata/settings.yaml",
			},
			want: &Settings{
				Server: &ServerSettings{
					Address:                ":1234",
					DevelopmentEnvironment: false,
				},
				Logging: &LoggingSettings{
					DevelopmentEnvironment: true,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromFile(tt.args.path)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
