package expressions

import (
	"math"
)

type Expression interface {
	Solve(ctx Context) float64
}

type ExpressionValue struct {
	value float64
}

func (e *ExpressionValue) Solve(ctx Context) float64 {
	return e.value
}

type ExpressionField struct {
	field string
}

func (e *ExpressionField) Solve(ctx Context) float64 {
	return ctx.Resolver().Resolve(e.field)
}

type ExpressionMultiple struct {
	terms []*ExpressionMultiplePart
}

func NewExpressionMultiple() *ExpressionMultiple {
	return &ExpressionMultiple{}
}

func (e *ExpressionMultiple) Solve(ctx Context) float64 {
	r := float64(0)
	for i, p := range e.terms {
		if i == 0 {
			r = p.Solve(ctx)
		} else {
			ctx.SetAccumulated(r)
			r = p.Solve(ctx)
		}
	}
	return r
}

func (e *ExpressionMultiple) Add(operator string, exp Expression) {
	e.terms = append(e.terms, &ExpressionMultiplePart{
		operator:   operator,
		expression: exp,
	})
}

type ExpressionMultiplePart struct {
	operator   string
	expression Expression
}

func (e *ExpressionMultiplePart) Solve(ctx Context) float64 {
	switch e.operator {
	case "":
		return e.expression.Solve(ctx)
	case "+":
		return ctx.Accumulated() + e.expression.Solve(ctx)
	case "-":
		return ctx.Accumulated() - e.expression.Solve(ctx)
	case "*":
		return ctx.Accumulated() * e.expression.Solve(ctx)
	case "/":
		return ctx.Accumulated() / e.expression.Solve(ctx)
	case "^":
		return math.Pow(ctx.Accumulated(), e.expression.Solve(ctx))
	}
	return math.NaN()
}

type ExpressionBinary struct {
	left     Expression
	operator string
	right    Expression
}

func (e *ExpressionBinary) Solve(ctx Context) float64 {
	left := e.left.Solve(ctx)
	right := e.right.Solve(ctx)
	switch e.operator {
	case "^":
		return math.Pow(left, right)
	case "*":
		return left * right
	case "/":
		return left / right
	case "+":
		return left + right
	case "-":
		return left - right
	}
	return math.NaN()
}

type ExpressionBrackets struct {
	inner Expression
}

func (e *ExpressionBrackets) Solve(ctx Context) float64 {
	return e.inner.Solve(ctx)
}
