package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		source context.Context
	}

	tests := []struct {
		name string
		args args
		want *Context
	}{
		{
			name: "Given a context.Context, New should create a new Context",
			args: args{
				source: context.Background(),
			},
			want: &Context{
				Context: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.source))
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		input context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    *Context
		wantErr bool
	}{
		{
			name: "Given a context.Context that can be cast to Context type, a parsed Context should be returned",
			args: args{
				input: &Context{
					Context: context.Background(),
				},
			},
			want: &Context{
				Context: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Given a context.Context that cannot be cast to Context type, an error should be returned",
			args: args{
				input: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.input)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestContext_Add(t *testing.T) {
	type fields struct {
		Context context.Context
	}

	type args struct {
		key   Key
		value interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
	}{
		{
			name: "If I call Add passing a Key and a value, they should be stored in the Context",
			fields: fields{
				Context: &Context{
					Context: context.Background(),
				},
			},
			args: args{
				key:   UserKey,
				value: "foo",
			},
			want: &Context{
				Context: context.WithValue(context.Background(), UserKey, "foo"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Context: tt.fields.Context,
			}

			c.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestContext_Get(t *testing.T) {
	type fields struct {
		Context context.Context
	}

	type args struct {
		key Key
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "Given a Key, its associated value in the Context should be returned",
			fields: fields{
				Context: &Context{
					Context: context.WithValue(context.Background(), UserKey, "foo"),
				},
			},
			args: args{
				key: UserKey,
			},
			want: "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				Context: tt.fields.Context,
			}

			assert.Equal(t, tt.want, c.Get(tt.args.key))
		})
	}
}
