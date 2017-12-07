package expressions

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
	return resolver.m[name], nil
}
