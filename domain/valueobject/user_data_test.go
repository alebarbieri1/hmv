package valueobject

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserData_Validate(t *testing.T) {
	type fields struct {
		Name string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "If Name is valid, no errors should be returned",
			fields: fields{
				Name: "foo",
			},
			wantErr: false,
		},
		{
			name: "If Name is empty, an error should be returned",
			fields: fields{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "If Name has more than 256 characters, an error should be returned",
			fields: fields{
				Name: strings.Repeat("a", 257),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UserData{
				Name: tt.fields.Name,
			}

			err := d.Validate()

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
