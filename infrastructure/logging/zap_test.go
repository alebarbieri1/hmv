package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_toZapFields(t *testing.T) {
	type args struct {
		fields []field
	}

	tests := []struct {
		name string
		args args
		want []zap.Field
	}{
		{
			name: "Given a set of fields, toZapFields should return a set of zap.Fields",
			args: args{
				fields: []field{
					{
						Name:  "foo",
						Value: "bar",
					},
					{
						Name:  "baz",
						Value: "qux",
					},
				},
			},
			want: []zap.Field{
				zap.String("foo", "bar"),
				zap.String("baz", "qux"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, toZapFields(tt.args.fields))
		})
	}
}
