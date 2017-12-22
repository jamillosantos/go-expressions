package expressions

import (
	"errors"
	"fmt"
)

type Resolver interface {
	Resolve(name string) (interface{}, error)
}

type MapResolver struct {
	m map[string]interface{}
}

func NewMapResolver(m map[string]interface{}) *MapResolver {
	return &MapResolver{
		m: m,
	}
}

func (resolver *MapResolver) Resolve(name string) (interface{}, error) {
	if v, ok := resolver.m[name]; ok {
		return v, nil
	}
	return nil, errors.New(fmt.Sprintf("Value of %s was not found.", name))
}
