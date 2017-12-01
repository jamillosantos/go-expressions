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
		return &ExpressionValue{
			value: v,
		}
	case *parser.AtomContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else {
			return newExpression(e.GetChild(1))
		}
	case *parser.SignedAtomContext:
		if e.GetChildCount() == 1 {
			return newExpression(e.GetChild(0))
		} else {
			// TODO
			panic("TODO!")
		}
	case *parser.ExpressionContext:
		r := NewExpressionMultiple()
		for i, me := range e.AllMultiplyingExpression() {
			if i == 0 {
				r.Add("", newExpression(me))
			} else {
				r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), newExpression(me))
			}
		}
		return r
	case *parser.MultiplyingExpressionContext:
		powExpressions := e.AllPowExpression()
		if len(powExpressions) == 1 {
			return newExpression(powExpressions[0])
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
