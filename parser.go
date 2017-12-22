package expressions

import (
	"github.com/jamillosantos/go-expressions/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"errors"
	"fmt"
)

type ExpressionError error

func NewExpression(expression antlr.Tree) Expression {
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
			return NewExpression(e.GetChild(0))
		} else {
			return NewExpression(e.GetChild(1))
		}
	case *parser.SignedAtomContext:
		if e.GetChildCount() == 1 {
			return NewExpression(e.GetChild(0))
		}
		expr := NewExpressionMultiple()
		expr.Add("", NewExpressionValue(0))
		expr.Add(e.GetOperator().GetText(), NewExpression(e.GetChild(1)))
		return expr
	case *parser.ExpressionContext:
		if e.GetChildCount() == 1 {
			return NewExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllMultiplyingExpression() {
				if i == 0 {
					r.Add("", NewExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), NewExpression(me))
				}
			}
			return r
		}
	case *parser.MultiplyingExpressionContext:
		childCount := e.GetChildCount()
		if childCount == 1 {
			return NewExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllPowExpression() {
				if i == 0 {
					r.Add("", NewExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), NewExpression(me))
				}
			}
			return r
		}
	case *parser.PowExpressionContext:
		if e.GetChildCount() == 1 {
			return NewExpression(e.GetChild(0))
		} else {
			r := NewExpressionMultiple()
			for i, me := range e.AllSignedAtom() {
				if i == 0 {
					r.Add("", NewExpression(me))
				} else {
					r.Add(e.GetChild(i*2 - 1).(*antlr.TerminalNodeImpl).GetText(), NewExpression(me))
				}
			}
			return r
		}
	case *parser.FunctionContext:
		eParams := e.AllExpression()
		params := make([]Expression, 0, len(eParams))
		for _, p := range eParams {
			params = append(params, NewExpression(p))
		}
		return NewExpressionFunction(e.GetFname().GetText(), params...)
	case *parser.BinaryOpContext:
		atoms := e.AllAtom()
		return NewExpressionBinary(NewExpression(atoms[0]), e.Relop().GetText(), NewExpression(atoms[1]))
	case *parser.StrContext:
		str := e.GetText()
		return NewExpressionValue(str[1:len(str)-1])
	}
	return nil
}

type CaptureErrorListener struct {
	errors []error
}

func NewCaptureErrorListener() *CaptureErrorListener {
	return &CaptureErrorListener{
		errors: make([]error, 0),
	}
}

func (errorListener *CaptureErrorListener) HasErrors() bool {
	return len(errorListener.errors) > 0
}

func (errorListener *CaptureErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorListener.errors = append(errorListener.errors, errors.New(fmt.Sprintf("%s at %d:%d", msg, line, column)))
}

func (errorListener *CaptureErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (errorListener *CaptureErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (errorListener *CaptureErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func Compile(expression string) (Expression, error) {
	input := antlr.NewInputStream(expression)
	lexer := parser.NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionParser(stream)

	errorListener := NewCaptureErrorListener()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true
	expr := p.Expression()
	if errorListener.HasErrors() {
		return nil, errorListener.errors[0]
	}
	return NewExpression(expr), nil
}
