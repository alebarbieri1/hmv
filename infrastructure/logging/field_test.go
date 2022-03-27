package logging

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	type args struct {
		name  string
		value string
	}

	tests := []struct {
		name string
		args args
		want field
	}{
		{
			name: "Given a key and a string value, String() should create a new field",
			args: args{
				name:  "foo",
				value: "bar",
			},
			want: field{
				Name:  "foo",
				Value: "bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, String(tt.args.name, tt.args.value))
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		name  string
		value int
	}

	tests := []struct {
		name string
		args args
		want field
	}{
		{
			name: "Given a key and a int value, Int() should create a new field",
			args: args{
				name:  "foo",
				value: 123,
			},
			want: field{
				Name:  "foo",
				Value: "123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Int(tt.args.name, tt.args.value))
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		value error
	}

	err := errors.New("foo")

	tests := []struct {
		name string
		args args
		want field
	}{
		{
			name: "Given an error value, Error() should create a new field with an 'error' Name",
			args: args{
				value: err,
			},
			want: field{
				Name:  "error",
				Value: err.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Error(tt.args.value))
		})
	}
}
