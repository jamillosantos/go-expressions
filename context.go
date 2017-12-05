package expressions

type Context interface {
	SetAccumulated(e float64)
	Accumulated() float64

	Resolver() Resolver
	Functions() Functions
}

type BaseContext struct {
	accumulated float64
	resolver    Resolver
	functions   Functions
}

func NewContext(resolver Resolver, functions Functions) *BaseContext {
	return &BaseContext{
		resolver:  resolver,
		functions: functions,
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

func (ctx *BaseContext) Functions() Functions {
	return ctx.functions
}
