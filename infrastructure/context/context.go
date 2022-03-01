package context

import (
	"context"
	"fmt"
)

type Key string

const UserKey Key = "user"

type Context struct {
	context.Context
}

func New() *Context {
	return &Context{Context: context.Background()}
}

func Parse(input context.Context) (*Context, error) {
	if ctx, implements := input.(*Context); implements {
		return ctx, nil
	}

	return nil, fmt.Errorf("failed to parse context.Context: input does not implement type Context")
}

func (c *Context) Add(key Key, value interface{}) {
	c.Context = context.WithValue(c.Context, key, value)
}

func (c *Context) Get(key Key) interface{} {
	return c.Context.Value(key)
}
