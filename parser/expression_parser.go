// Generated from /home/jsantos/projects/aure-api/src/github.com/jamillosantos/go-expressions/antlr/Expression.g4 by ANTLR 4.7.

package parser // Expression

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 98, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 
	13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14, 2, 35, 
	11, 2, 3, 3, 3, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3, 43, 11, 3, 3, 4, 
	3, 4, 3, 4, 7, 4, 48, 10, 4, 12, 4, 14, 4, 51, 11, 4, 3, 5, 3, 5, 3, 5, 
	3, 5, 3, 5, 5, 5, 58, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 72, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 87, 
	10, 12, 12, 12, 14, 12, 90, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 
	3, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 
	2, 7, 3, 2, 15, 16, 3, 2, 17, 19, 3, 2, 29, 31, 3, 2, 3, 12, 3, 2, 20, 
	25, 2, 95, 2, 28, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 
	57, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2, 14, 73, 3, 2, 2, 
	2, 16, 75, 3, 2, 2, 2, 18, 77, 3, 2, 2, 2, 20, 79, 3, 2, 2, 2, 22, 81, 
	3, 2, 2, 2, 24, 93, 3, 2, 2, 2, 26, 95, 3, 2, 2, 2, 28, 33, 5, 4, 3, 2, 
	29, 30, 9, 2, 2, 2, 30, 32, 5, 4, 3, 2, 31, 29, 3, 2, 2, 2, 32, 35, 3, 
	2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 
	33, 3, 2, 2, 2, 36, 41, 5, 6, 4, 2, 37, 38, 9, 3, 2, 2, 38, 40, 5, 6, 4, 
	2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 
	3, 2, 2, 2, 42, 5, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 49, 5, 8, 5, 2, 
	45, 46, 7, 28, 2, 2, 46, 48, 5, 8, 5, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 
	2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 7, 3, 2, 2, 2, 51, 
	49, 3, 2, 2, 2, 52, 53, 9, 2, 2, 2, 53, 58, 5, 8, 5, 2, 54, 58, 5, 22, 
	12, 2, 55, 58, 5, 10, 6, 2, 56, 58, 5, 12, 7, 2, 57, 52, 3, 2, 2, 2, 57, 
	54, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 9, 3, 2, 2, 
	2, 59, 60, 5, 12, 7, 2, 60, 61, 5, 26, 14, 2, 61, 62, 5, 12, 7, 2, 62, 
	11, 3, 2, 2, 2, 63, 72, 5, 16, 9, 2, 64, 72, 5, 20, 11, 2, 65, 72, 5, 18, 
	10, 2, 66, 67, 7, 13, 2, 2, 67, 68, 5, 2, 2, 2, 68, 69, 7, 14, 2, 2, 69, 
	72, 3, 2, 2, 2, 70, 72, 5, 14, 8, 2, 71, 63, 3, 2, 2, 2, 71, 64, 3, 2, 
	2, 2, 71, 65, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 13, 
	3, 2, 2, 2, 73, 74, 7, 33, 2, 2, 74, 15, 3, 2, 2, 2, 75, 76, 7, 35, 2, 
	2, 76, 17, 3, 2, 2, 2, 77, 78, 9, 4, 2, 2, 78, 19, 3, 2, 2, 2, 79, 80, 
	7, 32, 2, 2, 80, 21, 3, 2, 2, 2, 81, 82, 5, 24, 13, 2, 82, 83, 7, 13, 2, 
	2, 83, 88, 5, 2, 2, 2, 84, 85, 7, 26, 2, 2, 85, 87, 5, 2, 2, 2, 86, 84, 
	3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 
	89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 92, 7, 14, 2, 2, 92, 23, 3, 
	2, 2, 2, 93, 94, 9, 5, 2, 2, 94, 25, 3, 2, 2, 2, 95, 96, 9, 6, 2, 2, 96, 
	27, 3, 2, 2, 2, 8, 33, 41, 49, 57, 71, 88,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'cos'", "'sin'", "'tan'", "'acos'", "'asin'", "'atan'", "'ln'", "'log'", 
	"'sqrt'", "'if'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", 
	"'<'", "'=='", "'||'", "'&&'", "'xor'", "','", "'.'", "'^'", "'pi'", "", 
	"'i'", "", "", "'\"'",
}
var symbolicNames = []string{
	"", "COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "IF", 
	"LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", "I", 
	"VARIABLE", "QUOTED_STRING", "QUOTE", "SCIENTIFIC_NUMBER", "WS",
}

var ruleNames = []string{
	"expression", "multiplyingExpression", "powExpression", "signedAtom", "binaryOp", 
	"atom", "string", "scientific", "constant", "variable", "function", "funcname", 
	"relop",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExpressionParser struct {
	*antlr.BaseParser
}

func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	this := new(ExpressionParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}

// ExpressionParser tokens.
const (
	ExpressionParserEOF = antlr.TokenEOF
	ExpressionParserCOS = 1
	ExpressionParserSIN = 2
	ExpressionParserTAN = 3
	ExpressionParserACOS = 4
	ExpressionParserASIN = 5
	ExpressionParserATAN = 6
	ExpressionParserLN = 7
	ExpressionParserLOG = 8
	ExpressionParserSQRT = 9
	ExpressionParserIF = 10
	ExpressionParserLPAREN = 11
	ExpressionParserRPAREN = 12
	ExpressionParserPLUS = 13
	ExpressionParserMINUS = 14
	ExpressionParserTIMES = 15
	ExpressionParserDIV = 16
	ExpressionParserMOD = 17
	ExpressionParserGT = 18
	ExpressionParserLT = 19
	ExpressionParserEQ = 20
	ExpressionParserOR = 21
	ExpressionParserAND = 22
	ExpressionParserXOR = 23
	ExpressionParserCOMMA = 24
	ExpressionParserPOINT = 25
	ExpressionParserPOW = 26
	ExpressionParserPI = 27
	ExpressionParserEULER = 28
	ExpressionParserI = 29
	ExpressionParserVARIABLE = 30
	ExpressionParserQUOTED_STRING = 31
	ExpressionParserQUOTE = 32
	ExpressionParserSCIENTIFIC_NUMBER = 33
	ExpressionParserWS = 34
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression = 0
	ExpressionParserRULE_multiplyingExpression = 1
	ExpressionParserRULE_powExpression = 2
	ExpressionParserRULE_signedAtom = 3
	ExpressionParserRULE_binaryOp = 4
	ExpressionParserRULE_atom = 5
	ExpressionParserRULE_string = 6
	ExpressionParserRULE_scientific = 7
	ExpressionParserRULE_constant = 8
	ExpressionParserRULE_variable = 9
	ExpressionParserRULE_function = 10
	ExpressionParserRULE_funcname = 11
	ExpressionParserRULE_relop = 12
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserPLUS)
}

func (s *ExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserPLUS, i)
}

func (s *ExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserMINUS)
}

func (s *ExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserMINUS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}




func (p *ExpressionParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionParserRULE_expression)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.MultiplyingExpression()
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserPLUS || _la == ExpressionParserMINUS {
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ExpressionParserPLUS || _la == ExpressionParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(28)
			p.MultiplyingExpression()
		}


		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllPowExpression() []IPowExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem())
	var tst = make([]IPowExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPowExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) PowExpression(i int) IPowExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPowExpressionContext)
}

func (s *MultiplyingExpressionContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserTIMES)
}

func (s *MultiplyingExpressionContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserTIMES, i)
}

func (s *MultiplyingExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserDIV)
}

func (s *MultiplyingExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserDIV, i)
}

func (s *MultiplyingExpressionContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserMOD)
}

func (s *MultiplyingExpressionContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserMOD, i)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}




func (p *ExpressionParser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpressionParserRULE_multiplyingExpression)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.PowExpression()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserTIMES) | (1 << ExpressionParserDIV) | (1 << ExpressionParserMOD))) != 0) {
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserTIMES) | (1 << ExpressionParserDIV) | (1 << ExpressionParserMOD))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(36)
			p.PowExpression()
		}


		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IPowExpressionContext is an interface to support dynamic dispatch.
type IPowExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowExpressionContext differentiates from other interfaces.
	IsPowExpressionContext()
}

type PowExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowExpressionContext() *PowExpressionContext {
	var p = new(PowExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_powExpression
	return p
}

func (*PowExpressionContext) IsPowExpressionContext() {}

func NewPowExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowExpressionContext {
	var p = new(PowExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_powExpression

	return p
}

func (s *PowExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PowExpressionContext) AllSignedAtom() []ISignedAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem())
	var tst = make([]ISignedAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISignedAtomContext)
		}
	}

	return tst
}

func (s *PowExpressionContext) SignedAtom(i int) ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *PowExpressionContext) AllPOW() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserPOW)
}

func (s *PowExpressionContext) POW(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserPOW, i)
}

func (s *PowExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PowExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterPowExpression(s)
	}
}

func (s *PowExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitPowExpression(s)
	}
}




func (p *ExpressionParser) PowExpression() (localctx IPowExpressionContext) {
	localctx = NewPowExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_powExpression)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.SignedAtom()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserPOW {
		{
			p.SetState(43)
			p.Match(ExpressionParserPOW)
		}
		{
			p.SetState(44)
			p.SignedAtom()
		}


		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISignedAtomContext is an interface to support dynamic dispatch.
type ISignedAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token 


	// SetOperator sets the operator token.
	SetOperator(antlr.Token) 


	// IsSignedAtomContext differentiates from other interfaces.
	IsSignedAtomContext()
}

type SignedAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	operator antlr.Token
}

func NewEmptySignedAtomContext() *SignedAtomContext {
	var p = new(SignedAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_signedAtom
	return p
}

func (*SignedAtomContext) IsSignedAtomContext() {}

func NewSignedAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedAtomContext {
	var p = new(SignedAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_signedAtom

	return p
}

func (s *SignedAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedAtomContext) GetOperator() antlr.Token { return s.operator }


func (s *SignedAtomContext) SetOperator(v antlr.Token) { s.operator = v }


func (s *SignedAtomContext) SignedAtom() ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *SignedAtomContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserPLUS, 0)
}

func (s *SignedAtomContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserMINUS, 0)
}

func (s *SignedAtomContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *SignedAtomContext) BinaryOp() IBinaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *SignedAtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SignedAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SignedAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterSignedAtom(s)
	}
}

func (s *SignedAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitSignedAtom(s)
	}
}




func (p *ExpressionParser) SignedAtom() (localctx ISignedAtomContext) {
	localctx = NewSignedAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionParserRULE_signedAtom)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(50)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*SignedAtomContext).operator = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ExpressionParserPLUS || _la == ExpressionParserMINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*SignedAtomContext).operator = _ri
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(51)
			p.SignedAtom()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Function()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.BinaryOp()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Atom()
		}

	}


	return localctx
}


// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOpContext differentiates from other interfaces.
	IsBinaryOpContext()
}

type BinaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpContext() *BinaryOpContext {
	var p = new(BinaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_binaryOp
	return p
}

func (*BinaryOpContext) IsBinaryOpContext() {}

func NewBinaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_binaryOp

	return p
}

func (s *BinaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *BinaryOpContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *BinaryOpContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BinaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBinaryOp(s)
	}
}

func (s *BinaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBinaryOp(s)
	}
}




func (p *ExpressionParser) BinaryOp() (localctx IBinaryOpContext) {
	localctx = NewBinaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExpressionParserRULE_binaryOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Atom()
	}
	{
		p.SetState(58)
		p.Relop()
	}
	{
		p.SetState(59)
		p.Atom()
	}



	return localctx
}


// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Scientific() IScientificContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScientificContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScientificContext)
}

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *AtomContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *AtomContext) String() IStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitAtom(s)
	}
}




func (p *ExpressionParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExpressionParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserSCIENTIFIC_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Scientific()
		}


	case ExpressionParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Variable()
		}


	case ExpressionParserPI, ExpressionParserEULER, ExpressionParserI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Constant()
		}


	case ExpressionParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Match(ExpressionParserLPAREN)
		}
		{
			p.SetState(65)
			p.Expression()
		}
		{
			p.SetState(66)
			p.Match(ExpressionParserRPAREN)
		}


	case ExpressionParserQUOTED_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.String()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_string
	return p
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(ExpressionParserQUOTED_STRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitString(s)
	}
}




func (p *ExpressionParser) String() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExpressionParserRULE_string)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(ExpressionParserQUOTED_STRING)
	}



	return localctx
}


// IScientificContext is an interface to support dynamic dispatch.
type IScientificContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScientificContext differentiates from other interfaces.
	IsScientificContext()
}

type ScientificContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScientificContext() *ScientificContext {
	var p = new(ScientificContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_scientific
	return p
}

func (*ScientificContext) IsScientificContext() {}

func NewScientificContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScientificContext {
	var p = new(ScientificContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_scientific

	return p
}

func (s *ScientificContext) GetParser() antlr.Parser { return s.parser }

func (s *ScientificContext) SCIENTIFIC_NUMBER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSCIENTIFIC_NUMBER, 0)
}

func (s *ScientificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ScientificContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterScientific(s)
	}
}

func (s *ScientificContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitScientific(s)
	}
}




func (p *ExpressionParser) Scientific() (localctx IScientificContext) {
	localctx = NewScientificContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExpressionParserRULE_scientific)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(ExpressionParserSCIENTIFIC_NUMBER)
	}



	return localctx
}


// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) PI() antlr.TerminalNode {
	return s.GetToken(ExpressionParserPI, 0)
}

func (s *ConstantContext) EULER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEULER, 0)
}

func (s *ConstantContext) I() antlr.TerminalNode {
	return s.GetToken(ExpressionParserI, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitConstant(s)
	}
}




func (p *ExpressionParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExpressionParserRULE_constant)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserPI) | (1 << ExpressionParserEULER) | (1 << ExpressionParserI))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ExpressionParserVARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitVariable(s)
	}
}




func (p *ExpressionParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExpressionParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(ExpressionParserVARIABLE)
	}



	return localctx
}


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *FunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *FunctionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserCOMMA)
}

func (s *FunctionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserCOMMA, i)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFunction(s)
	}
}




func (p *ExpressionParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ExpressionParserRULE_function)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Funcname()
	}
	{
		p.SetState(80)
		p.Match(ExpressionParserLPAREN)
	}
	{
		p.SetState(81)
		p.Expression()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserCOMMA {
		{
			p.SetState(82)
			p.Match(ExpressionParserCOMMA)
		}
		{
			p.SetState(83)
			p.Expression()
		}


		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(ExpressionParserRPAREN)
	}



	return localctx
}


// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) COS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserCOS, 0)
}

func (s *FuncnameContext) TAN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserTAN, 0)
}

func (s *FuncnameContext) SIN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSIN, 0)
}

func (s *FuncnameContext) ACOS() antlr.TerminalNode {
	return s.GetToken(ExpressionParserACOS, 0)
}

func (s *FuncnameContext) ATAN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserATAN, 0)
}

func (s *FuncnameContext) ASIN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserASIN, 0)
}

func (s *FuncnameContext) LOG() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLOG, 0)
}

func (s *FuncnameContext) LN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLN, 0)
}

func (s *FuncnameContext) SQRT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSQRT, 0)
}

func (s *FuncnameContext) IF() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIF, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFuncname(s)
	}
}




func (p *ExpressionParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ExpressionParserRULE_funcname)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserCOS) | (1 << ExpressionParserSIN) | (1 << ExpressionParserTAN) | (1 << ExpressionParserACOS) | (1 << ExpressionParserASIN) | (1 << ExpressionParserATAN) | (1 << ExpressionParserLN) | (1 << ExpressionParserLOG) | (1 << ExpressionParserSQRT) | (1 << ExpressionParserIF))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEQ, 0)
}

func (s *RelopContext) GT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserGT, 0)
}

func (s *RelopContext) LT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLT, 0)
}

func (s *RelopContext) OR() antlr.TerminalNode {
	return s.GetToken(ExpressionParserOR, 0)
}

func (s *RelopContext) AND() antlr.TerminalNode {
	return s.GetToken(ExpressionParserAND, 0)
}

func (s *RelopContext) XOR() antlr.TerminalNode {
	return s.GetToken(ExpressionParserXOR, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitRelop(s)
	}
}




func (p *ExpressionParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ExpressionParserRULE_relop)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserGT) | (1 << ExpressionParserLT) | (1 << ExpressionParserEQ) | (1 << ExpressionParserOR) | (1 << ExpressionParserAND) | (1 << ExpressionParserXOR))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


