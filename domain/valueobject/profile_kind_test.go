package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileKind_String(t *testing.T) {
	tests := []struct {
		name string
		p    ProfileKind
		want string
	}{
		{
			name: "Given a ProfileKind, its string value should be returned",
			p:    Pacient_ProfileKind,
			want: "pacient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.p.String())
		})
	}
}
