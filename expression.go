package expressions

import (
	"math"
	"fmt"
	"errors"
)

type WrongTypeError struct {
	v interface{}
}

func NewWrongTypeError(v interface{}) *WrongTypeError {
	return &WrongTypeError{
		v: v,
	}
}

func (err *WrongTypeError) Error() string {
	return fmt.Sprintf("%s has not a valid type.", fmt.Sprint(err.v))
}

type Expression interface {
	Solve(ctx Context) (interface{}, error)
}

type ExpressionValue struct {
	value interface{}
}

func NewExpressionValue(value interface{}) *ExpressionValue {
	return &ExpressionValue{
		value: value,
	}
}

func (e *ExpressionValue) Solve(ctx Context) (interface{}, error) {
	return e.value, nil
}

type ExpressionField struct {
	field string
}

func NewExpressionField(field string) *ExpressionField {
	return &ExpressionField{
		field: field,
	}
}

func (e *ExpressionField) Solve(ctx Context) (interface{}, error) {
	return ctx.Resolver().Resolve(e.field)
}

type ExpressionMultiple struct {
	terms []*ExpressionMultiplePart
}

func NewExpressionMultiple() *ExpressionMultiple {
	return &ExpressionMultiple{}
}

func (e *ExpressionMultiple) Solve(ctx Context) (interface{}, error) {
	result := float64(0)
	var r float64
	for i, p := range e.terms {
		if i == 0 {
			rTemp, err := p.Solve(ctx)
			if err != nil {
				return nil, err
			}
			switch rr := rTemp.(type) {
			case int:
				r = float64(rr)
			case float64:
				r = rr
			default:
				return nil, NewWrongTypeError(rTemp)
			}

			result = r
		} else {
			ctx.SetAccumulated(result)
			rTemp, err := p.Solve(ctx)
			if err != nil {
				return nil, err
			}
			switch rr := rTemp.(type) {
			case int:
				r = float64(rr)
			case float64:
				r = rr
			}
			result = r
		}
	}
	return result, nil
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

func NewExpressionMultiplePart(operator string, expression Expression) *ExpressionMultiplePart {
	return &ExpressionMultiplePart{
		operator:   operator,
		expression: expression,
	}
}

func (e *ExpressionMultiplePart) Solve(ctx Context) (interface{}, error) {
	accumulated := ctx.Accumulated()
	v, err := e.expression.Solve(ctx)
	if err != nil {
		return v, err
	}
	switch e.operator {
	case "":
		return v, nil
	case "+":
		switch vv := v.(type) {
		case int:
			return accumulated + float64(vv), nil
		case float64:
			return accumulated + vv, nil
		default:
			return nil, NewWrongTypeError(v)
		}
	case "-":
		switch vv := v.(type) {
		case int:
			return accumulated - float64(vv), nil
		case float64:
			return accumulated - vv, nil
		default:
			return nil, NewWrongTypeError(v)
		}
	case "*":
		switch vv := v.(type) {
		case int:
			return accumulated * float64(vv), nil
		case float64:
			return accumulated * vv, nil
		default:
			return nil, NewWrongTypeError(v)
		}
	case "/":
		switch vv := v.(type) {
		case int:
			return accumulated / float64(vv), nil
		case float64:
			return accumulated / vv, nil
		default:
			return nil, NewWrongTypeError(v)
		}
	case "%":
		switch vv := v.(type) {
		case int:
			return int(accumulated+0.5) % vv, nil
		case float64:
			return int(accumulated+0.5) % int(vv+0.5), nil
		default:
			return nil, NewWrongTypeError(v)
		}
	case "^":
		switch vv := v.(type) {
		case int:
			return math.Pow(accumulated, float64(vv)), nil
		case float64:
			return math.Pow(accumulated, vv), nil
		default:
			return nil, NewWrongTypeError(v)
		}
	default:
		return nil, errors.New(fmt.Sprintf("The operator '%s' is not supported.", e.operator))
	}
}

type ExpressionBinary struct {
	left     Expression
	operator string
	right    Expression
}

func NewExpressionBinary(left Expression, operator string, right Expression) *ExpressionBinary {
	return &ExpressionBinary{
		left:     left,
		operator: operator,
		right:    right,
	}
}

func (e *ExpressionBinary) Solve(ctx Context) (interface{}, error) {
	rLeft, err := e.left.Solve(ctx)
	if err != nil {
		return nil, err
	}
	rRight, err := e.right.Solve(ctx)
	if err != nil {
		return nil, err
	}
	switch e.operator {
	case ">", "<", ">=", "<=":
		var (
			left  float64
			right float64
		)
		switch r := rLeft.(type) {
		case int:
			left = float64(r)
		case float64:
			left = r
		default:
			return nil, NewWrongTypeError(rLeft)
		}
		switch r := rRight.(type) {
		case int:
			right = float64(r)
		case float64:
			right = r
		default:
			return nil, NewWrongTypeError(rRight)
		}
		switch e.operator {
		case ">":
			return left > right, nil
		case "<":
			return left < right, nil
		case ">=":
			return left >= right, nil
		case "<=":
			return left <= right, nil
		}
	case "==", "!=":
		switch e.operator {
		case "==":
			return rLeft == rRight, nil
		case "!=":
			return rLeft != rRight, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("The operator '%s' is not supported", e.operator))
}

type ExpressionBrackets struct {
	inner Expression
}

func NewExpressionBrackets(expression Expression) *ExpressionBrackets {
	return &ExpressionBrackets{
		inner: expression,
	}
}

func (e *ExpressionBrackets) Solve(ctx Context) (interface{}, error) {
	return e.inner.Solve(ctx)
}

type ExpressionFunction struct {
	name   string
	params []Expression
}

func NewExpressionFunction(name string, params ... Expression) *ExpressionFunction {
	return &ExpressionFunction{
		name:   name,
		params: params,
	}
}

func (e *ExpressionFunction) Solve(ctx Context) (interface{}, error) {
	return ctx.Functions().Call(ctx, e.name, e.params...)
}
