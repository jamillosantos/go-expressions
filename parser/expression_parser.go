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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 94, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9, 
	13, 3, 2, 3, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 3, 3, 
	3, 3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 4, 3, 4, 3, 4, 7, 
	4, 46, 10, 4, 12, 4, 14, 4, 49, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 
	5, 56, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 
	7, 3, 7, 3, 7, 5, 7, 70, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 85, 10, 12, 12, 12, 
	14, 12, 88, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 6, 3, 2, 5, 6, 3, 2, 7, 9, 3, 
	2, 20, 22, 3, 2, 10, 16, 2, 92, 2, 26, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 
	42, 3, 2, 2, 2, 8, 55, 3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 69, 3, 2, 2, 
	2, 14, 71, 3, 2, 2, 2, 16, 73, 3, 2, 2, 2, 18, 75, 3, 2, 2, 2, 20, 77, 
	3, 2, 2, 2, 22, 79, 3, 2, 2, 2, 24, 91, 3, 2, 2, 2, 26, 31, 5, 4, 3, 2, 
	27, 28, 9, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 27, 3, 2, 2, 2, 30, 33, 3, 
	2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 3, 3, 2, 2, 2, 33, 
	31, 3, 2, 2, 2, 34, 39, 5, 6, 4, 2, 35, 36, 9, 3, 2, 2, 36, 38, 5, 6, 4, 
	2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 
	3, 2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 47, 5, 8, 5, 2, 
	43, 44, 7, 19, 2, 2, 44, 46, 5, 8, 5, 2, 45, 43, 3, 2, 2, 2, 46, 49, 3, 
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 7, 3, 2, 2, 2, 49, 
	47, 3, 2, 2, 2, 50, 51, 9, 2, 2, 2, 51, 56, 5, 8, 5, 2, 52, 56, 5, 22, 
	12, 2, 53, 56, 5, 12, 7, 2, 54, 56, 5, 10, 6, 2, 55, 50, 3, 2, 2, 2, 55, 
	52, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 9, 3, 2, 2, 
	2, 57, 58, 5, 12, 7, 2, 58, 59, 5, 24, 13, 2, 59, 60, 5, 12, 7, 2, 60, 
	11, 3, 2, 2, 2, 61, 70, 5, 16, 9, 2, 62, 70, 5, 20, 11, 2, 63, 70, 5, 18, 
	10, 2, 64, 65, 7, 3, 2, 2, 65, 66, 5, 2, 2, 2, 66, 67, 7, 4, 2, 2, 67, 
	70, 3, 2, 2, 2, 68, 70, 5, 14, 8, 2, 69, 61, 3, 2, 2, 2, 69, 62, 3, 2, 
	2, 2, 69, 63, 3, 2, 2, 2, 69, 64, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 13, 
	3, 2, 2, 2, 71, 72, 7, 24, 2, 2, 72, 15, 3, 2, 2, 2, 73, 74, 7, 26, 2, 
	2, 74, 17, 3, 2, 2, 2, 75, 76, 9, 4, 2, 2, 76, 19, 3, 2, 2, 2, 77, 78, 
	7, 23, 2, 2, 78, 21, 3, 2, 2, 2, 79, 80, 7, 23, 2, 2, 80, 81, 7, 3, 2, 
	2, 81, 86, 5, 2, 2, 2, 82, 83, 7, 17, 2, 2, 83, 85, 5, 2, 2, 2, 84, 82, 
	3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 
	87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 7, 4, 2, 2, 90, 23, 3, 
	2, 2, 2, 91, 92, 9, 5, 2, 2, 92, 25, 3, 2, 2, 2, 8, 31, 39, 47, 55, 69, 
	86,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'=='", 
	"'!='", "'||'", "'&&'", "'xor'", "','", "'.'", "'^'", "'pi'", "", "'i'", 
	"", "", "'\"'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "NOT_EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", 
	"I", "VARIABLE", "QUOTED_STRING", "QUOTE", "SCIENTIFIC_NUMBER", "WS",
}

var ruleNames = []string{
	"expression", "multiplyingExpression", "powExpression", "signedAtom", "binaryOp", 
	"atom", "str", "scientific", "constant", "variable", "function", "relop",
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
	ExpressionParserLPAREN = 1
	ExpressionParserRPAREN = 2
	ExpressionParserPLUS = 3
	ExpressionParserMINUS = 4
	ExpressionParserTIMES = 5
	ExpressionParserDIV = 6
	ExpressionParserMOD = 7
	ExpressionParserGT = 8
	ExpressionParserLT = 9
	ExpressionParserEQ = 10
	ExpressionParserNOT_EQ = 11
	ExpressionParserOR = 12
	ExpressionParserAND = 13
	ExpressionParserXOR = 14
	ExpressionParserCOMMA = 15
	ExpressionParserPOINT = 16
	ExpressionParserPOW = 17
	ExpressionParserPI = 18
	ExpressionParserEULER = 19
	ExpressionParserI = 20
	ExpressionParserVARIABLE = 21
	ExpressionParserQUOTED_STRING = 22
	ExpressionParserQUOTE = 23
	ExpressionParserSCIENTIFIC_NUMBER = 24
	ExpressionParserWS = 25
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression = 0
	ExpressionParserRULE_multiplyingExpression = 1
	ExpressionParserRULE_powExpression = 2
	ExpressionParserRULE_signedAtom = 3
	ExpressionParserRULE_binaryOp = 4
	ExpressionParserRULE_atom = 5
	ExpressionParserRULE_str = 6
	ExpressionParserRULE_scientific = 7
	ExpressionParserRULE_constant = 8
	ExpressionParserRULE_variable = 9
	ExpressionParserRULE_function = 10
	ExpressionParserRULE_relop = 11
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
		p.SetState(24)
		p.MultiplyingExpression()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserPLUS || _la == ExpressionParserMINUS {
		p.SetState(25)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ExpressionParserPLUS || _la == ExpressionParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(26)
			p.MultiplyingExpression()
		}


		p.SetState(31)
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
		p.SetState(32)
		p.PowExpression()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserTIMES) | (1 << ExpressionParserDIV) | (1 << ExpressionParserMOD))) != 0) {
		p.SetState(33)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserTIMES) | (1 << ExpressionParserDIV) | (1 << ExpressionParserMOD))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(34)
			p.PowExpression()
		}


		p.SetState(39)
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
		p.SetState(40)
		p.SignedAtom()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserPOW {
		{
			p.SetState(41)
			p.Match(ExpressionParserPOW)
		}
		{
			p.SetState(42)
			p.SignedAtom()
		}


		p.SetState(47)
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

func (s *SignedAtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SignedAtomContext) BinaryOp() IBinaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(48)

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
			p.SetState(49)
			p.SignedAtom()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Function()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Atom()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)
			p.BinaryOp()
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
		p.SetState(55)
		p.Atom()
	}
	{
		p.SetState(56)
		p.Relop()
	}
	{
		p.SetState(57)
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

func (s *AtomContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserSCIENTIFIC_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Scientific()
		}


	case ExpressionParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Variable()
		}


	case ExpressionParserPI, ExpressionParserEULER, ExpressionParserI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Constant()
		}


	case ExpressionParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Match(ExpressionParserLPAREN)
		}
		{
			p.SetState(63)
			p.Expression()
		}
		{
			p.SetState(64)
			p.Match(ExpressionParserRPAREN)
		}


	case ExpressionParserQUOTED_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(66)
			p.Str()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(ExpressionParserQUOTED_STRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitStr(s)
	}
}




func (p *ExpressionParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExpressionParserRULE_str)

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
		p.SetState(69)
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
		p.SetState(71)
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
	p.SetState(73)
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
		p.SetState(75)
		p.Match(ExpressionParserVARIABLE)
	}



	return localctx
}


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFname returns the fname token.
	GetFname() antlr.Token 


	// SetFname sets the fname token.
	SetFname(antlr.Token) 


	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fname antlr.Token
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

func (s *FunctionContext) GetFname() antlr.Token { return s.fname }


func (s *FunctionContext) SetFname(v antlr.Token) { s.fname = v }


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

func (s *FunctionContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ExpressionParserVARIABLE, 0)
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
		p.SetState(77)

		var _m = p.Match(ExpressionParserVARIABLE)

		localctx.(*FunctionContext).fname = _m
	}
	{
		p.SetState(78)
		p.Match(ExpressionParserLPAREN)
	}
	{
		p.SetState(79)
		p.Expression()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserCOMMA {
		{
			p.SetState(80)
			p.Match(ExpressionParserCOMMA)
		}
		{
			p.SetState(81)
			p.Expression()
		}


		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(ExpressionParserRPAREN)
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

func (s *RelopContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNOT_EQ, 0)
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
	p.EnterRule(localctx, 22, ExpressionParserRULE_relop)
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
	p.SetState(89)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionParserGT) | (1 << ExpressionParserLT) | (1 << ExpressionParserEQ) | (1 << ExpressionParserNOT_EQ) | (1 << ExpressionParserOR) | (1 << ExpressionParserAND) | (1 << ExpressionParserXOR))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


