package expressions

type Resolver interface {
	Resolve(name string) float64
}

type MapResolver struct {
	m map[string]float64
}

func NewMapResolver(m map[string]float64) *MapResolver {
	return &MapResolver{
		m: m,
	}
}

func (resolver *MapResolver) Resolve(name string) float64 {
	return resolver.m[name]
}
