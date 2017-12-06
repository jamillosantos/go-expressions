package expressions

import (
	"github.com/jamillosantos/go-expressions/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type ExpressionError error

func newExpression(expression antlr.Tree) Expression {
	switch e := expression.(type) {
	case *parser.VariableContext:
		return &ExpressionField{
			field: e.GetText(),
		}
	case *parser.ScientificContext:
		v, _ := strconv.ParseFloat(e.GetText(), 64)
		return NewExpressionValue(v)
	case *parser.AtomContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else {
			return newExpression(e.GetChild(1))
		}
	case *parser.SignedAtomContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else if e.GetChildCount() == 2 {
			expr := NewExpressionMultiple()
			expr.Add("", NewExpressionValue(0))
			expr.Add(e.GetOperator().GetText(), newExpression(e.GetChild(1)))
			return expr
		} else {
			// TODO
			panic("TODO! " + e.GetText())
		}
	case *parser.ExpressionContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllMultiplyingExpression() {
				if i == 0 {
					r.Add("", newExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), newExpression(me))
				}
			}
			return r
		}
	case *parser.MultiplyingExpressionContext:
		childCount := e.GetChildCount()
		if childCount == 1 {
			return newExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllPowExpression() {
				if i == 0 {
					r.Add("", newExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), newExpression(me))
				}
			}
			return r
		}
	case *parser.PowExpressionContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllSignedAtom() {
				if i == 0 {
					r.Add("", newExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), newExpression(me))
				}
			}
			return r
		}
	case *parser.FunctionContext:
		eParams := e.AllExpression()
		params := make([]Expression, 0, len(eParams))
		for _, p := range eParams {
			params = append(params, newExpression(p))
		}
		return NewExpressionFunction(e.GetFname().GetText(), params...)
	case *parser.BinaryOpContext:
		atoms := e.AllAtom()
		return NewExpressionBinary(newExpression(atoms[0]), e.Relop().GetText(), newExpression(atoms[1]))
	case *parser.StrContext:
		str := e.GetText()
		return NewExpressionValue(str[1:len(str)-1])
	}
	return nil
}

func Compile(expression string) (Expression, error) {
	input := antlr.NewInputStream(expression)
	lexer := parser.NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	return newExpression(p.Expression()), nil
}
