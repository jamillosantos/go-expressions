package expressions

type Context interface {
	SetAccumulated(e float64)
	Accumulated() float64
	Resolver() Resolver
}

type BaseContext struct {
	accumulated float64
	resolver    Resolver
}

func NewContext(resolver Resolver) *BaseContext {
	return &BaseContext{
		resolver: resolver,
	}
}

func (ctx *BaseContext) SetAccumulated(e float64) {
	ctx.accumulated = e
}

func (ctx *BaseContext) Accumulated() float64 {
	return ctx.accumulated
}

func (ctx *BaseContext) Resolver() Resolver {
	return ctx.resolver
}
