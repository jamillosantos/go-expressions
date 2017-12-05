package expressions

import (
	"math"
	"errors"
	"fmt"
)

type Functions interface {
	Call(ctx Context, name string, params ... Expression) (interface{}, error)
}

type ParameterError struct {
	name       string
	paramCount uint
}

func (err *ParameterError) Error() string {
	return fmt.Sprintf("'%s' expects %d parameters.", err.name, err.paramCount)
}

type DefaultFunctions struct {
}

func (*DefaultFunctions) Call(ctx Context, name string, params ... Expression) (interface{}, error) {
	switch name {
	case "cos":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Cos(float64(rr)), nil
		case float64:
			return math.Cos(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "cosh":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Cosh(float64(rr)), nil
		case float64:
			return math.Cosh(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "acos":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Acos(float64(rr)), nil
		case float64:
			return math.Acos(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "acosh":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Acosh(float64(rr)), nil
		case float64:
			return math.Acosh(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "sin":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Sin(float64(rr)), nil
		case float64:
			return math.Sin(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "sinh":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Sinh(float64(rr)), nil
		case float64:
			return math.Sinh(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "asin":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Asin(float64(rr)), nil
		case float64:
			return math.Asin(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "asinh":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Asinh(float64(rr)), nil
		case float64:
			return math.Asinh(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "sqrt":
		if len(params) != 1 && len(params) != 2 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		var (
			param1 float64
			param2 float64
		)
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			param1 = float64(rr)
		case float64:
			param1 = rr
		default:
			return nil, NewWrongTypeError(r)
		}
		if len(params) > 1 {
			r, err = params[1].Solve(ctx)
			if err != nil {
				return nil, err
			}
			switch rr := r.(type) {
			case int:
				param2 = float64(rr)
			case float64:
				param2 = rr
			default:
				return nil, NewWrongTypeError(r)
			}
			return math.Pow(param1, float64(1)/param2), nil
		}
		return math.Sqrt(param1), nil
	case "tan":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Tan(float64(rr)), nil
		case float64:
			return math.Tan(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "atan":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Atan(float64(rr)), nil
		case float64:
			return math.Atan(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "atan2":
		if len(params) != 2 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 2,
			}
		}
		var (
			param1 float64
			param2 float64
		)
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			param1 = float64(rr)
		case float64:
			param1 = rr
		default:
			return nil, NewWrongTypeError(r)
		}
		r, err = params[1].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			param2 = float64(rr)
		case float64:
			param2 = rr
		default:
			return nil, NewWrongTypeError(r)
		}
		return math.Atan2(param1, param2), nil
	case "atanh":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Atanh(float64(rr)), nil
		case float64:
			return math.Atanh(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "log":
		if len(params) != 1 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 1,
			}
		}
		r, err := params[0].Solve(ctx)
		if err != nil {
			return nil, err
		}
		switch rr := r.(type) {
		case int:
			return math.Log(float64(rr)), nil
		case float64:
			return math.Log(rr), nil
		default:
			return nil, NewWrongTypeError(r)
		}
	case "if":
		if len(params) != 3 {
			return nil, &ParameterError{
				name:       name,
				paramCount: 3,
			}
		}
		// TODO
		return nil, errors.New("Not implemented")
	default:
		return nil, errors.New(fmt.Sprintf("The function '%s' is not defined.", name))
	}
}
