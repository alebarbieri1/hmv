package context

import (
	"context"
	"fmt"
)

// Key is a key that should be used to map values in a Context
type Key string

// UserKey is a default key reserved for storing user data in a Context
const UserKey Key = "user"

// Context extends context.Context
type Context struct {
	context.Context
}

// New creates a new Context from a context.Context
func New(source context.Context) *Context {
	return &Context{Context: source}
}

// Parse parses a context.Context into a Context. If the input cannot be cast to Context type, an error should be returned.
func Parse(input context.Context) (*Context, error) {
	if ctx, implements := input.(*Context); implements {
		return ctx, nil
	}

	return nil, fmt.Errorf("failed to parse context.Context: input does not implement type Context")
}

// Add adds a Key/value pair to the Context
func (c *Context) Add(key Key, value interface{}) {
	c.Context = context.WithValue(c.Context, key, value)
}

// Get returns a value related to a Key in the Context
func (c *Context) Get(key Key) interface{} {
	return c.Context.Value(key)
}
