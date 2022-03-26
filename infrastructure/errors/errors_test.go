package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	type args struct {
		err    error
		target error
	}

	err1 := errors.New("err1")
	err2 := errors.New("err2")

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "If the input errors are equal, Is should return true",
			args: args{
				err:    err1,
				target: err1,
			},
			want: true,
		},
		{
			name: "If the input errors are not equal, Is should return false",
			args: args{
				err:    err1,
				target: err2,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Is(tt.args.err, tt.args.target))
		})
	}
}

func TestWithMessage(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Given an error and a message, WithMessage should return a formatted error composed by both",
			args: args{
				message: "error context",
				err:     errors.New("error cause"),
			},
			want: errors.New("error context: error cause"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, WithMessage(tt.args.message, tt.args.err))
		})
	}
}
