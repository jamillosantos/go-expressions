// Generated from /home/jsantos/projects/aure-api/src/github.com/jamillosantos/go-expressions/antlr/Expression.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 181, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 117, 10, 22, 12, 22, 14, 22, 
	120, 11, 22, 3, 23, 3, 23, 3, 23, 7, 23, 125, 10, 23, 12, 23, 14, 23, 128, 
	11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 5, 26, 
	138, 10, 26, 3, 27, 3, 27, 5, 27, 142, 10, 27, 3, 28, 3, 28, 3, 28, 5, 
	28, 147, 10, 28, 3, 28, 5, 28, 150, 10, 28, 3, 28, 3, 28, 5, 28, 154, 10, 
	28, 3, 29, 6, 29, 157, 10, 29, 13, 29, 14, 29, 158, 3, 29, 3, 29, 6, 29, 
	163, 10, 29, 13, 29, 14, 29, 164, 5, 29, 167, 10, 29, 3, 30, 3, 30, 3, 
	31, 3, 31, 3, 32, 3, 32, 3, 33, 6, 33, 176, 10, 33, 13, 33, 14, 33, 177, 
	3, 33, 3, 33, 3, 126, 2, 34, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 2, 51, 
	2, 53, 2, 55, 26, 57, 2, 59, 2, 61, 2, 63, 2, 65, 27, 3, 2, 6, 4, 2, 12, 
	12, 15, 15, 5, 2, 67, 92, 97, 97, 99, 124, 4, 2, 45, 45, 47, 47, 5, 2, 
	11, 12, 15, 15, 34, 34, 2, 184, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 
	67, 3, 2, 2, 2, 5, 69, 3, 2, 2, 2, 7, 71, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 
	11, 75, 3, 2, 2, 2, 13, 77, 3, 2, 2, 2, 15, 79, 3, 2, 2, 2, 17, 81, 3, 
	2, 2, 2, 19, 83, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2, 23, 88, 3, 2, 2, 2, 25, 
	91, 3, 2, 2, 2, 27, 94, 3, 2, 2, 2, 29, 97, 3, 2, 2, 2, 31, 101, 3, 2, 
	2, 2, 33, 103, 3, 2, 2, 2, 35, 105, 3, 2, 2, 2, 37, 107, 3, 2, 2, 2, 39, 
	110, 3, 2, 2, 2, 41, 112, 3, 2, 2, 2, 43, 114, 3, 2, 2, 2, 45, 121, 3, 
	2, 2, 2, 47, 131, 3, 2, 2, 2, 49, 133, 3, 2, 2, 2, 51, 137, 3, 2, 2, 2, 
	53, 141, 3, 2, 2, 2, 55, 143, 3, 2, 2, 2, 57, 156, 3, 2, 2, 2, 59, 168, 
	3, 2, 2, 2, 61, 170, 3, 2, 2, 2, 63, 172, 3, 2, 2, 2, 65, 175, 3, 2, 2, 
	2, 67, 68, 7, 42, 2, 2, 68, 4, 3, 2, 2, 2, 69, 70, 7, 43, 2, 2, 70, 6, 
	3, 2, 2, 2, 71, 72, 7, 45, 2, 2, 72, 8, 3, 2, 2, 2, 73, 74, 7, 47, 2, 2, 
	74, 10, 3, 2, 2, 2, 75, 76, 7, 44, 2, 2, 76, 12, 3, 2, 2, 2, 77, 78, 7, 
	49, 2, 2, 78, 14, 3, 2, 2, 2, 79, 80, 7, 39, 2, 2, 80, 16, 3, 2, 2, 2, 
	81, 82, 7, 64, 2, 2, 82, 18, 3, 2, 2, 2, 83, 84, 7, 62, 2, 2, 84, 20, 3, 
	2, 2, 2, 85, 86, 7, 63, 2, 2, 86, 87, 7, 63, 2, 2, 87, 22, 3, 2, 2, 2, 
	88, 89, 7, 35, 2, 2, 89, 90, 7, 63, 2, 2, 90, 24, 3, 2, 2, 2, 91, 92, 7, 
	126, 2, 2, 92, 93, 7, 126, 2, 2, 93, 26, 3, 2, 2, 2, 94, 95, 7, 40, 2, 
	2, 95, 96, 7, 40, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98, 7, 122, 2, 2, 98, 99, 
	7, 113, 2, 2, 99, 100, 7, 116, 2, 2, 100, 30, 3, 2, 2, 2, 101, 102, 7, 
	46, 2, 2, 102, 32, 3, 2, 2, 2, 103, 104, 7, 48, 2, 2, 104, 34, 3, 2, 2, 
	2, 105, 106, 7, 96, 2, 2, 106, 36, 3, 2, 2, 2, 107, 108, 7, 114, 2, 2, 
	108, 109, 7, 107, 2, 2, 109, 38, 3, 2, 2, 2, 110, 111, 5, 61, 31, 2, 111, 
	40, 3, 2, 2, 2, 112, 113, 7, 107, 2, 2, 113, 42, 3, 2, 2, 2, 114, 118, 
	5, 51, 26, 2, 115, 117, 5, 53, 27, 2, 116, 115, 3, 2, 2, 2, 117, 120, 3, 
	2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 44, 3, 2, 2, 
	2, 120, 118, 3, 2, 2, 2, 121, 126, 5, 47, 24, 2, 122, 125, 5, 49, 25, 2, 
	123, 125, 10, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 
	128, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 129, 
	3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 5, 47, 24, 2, 130, 46, 3, 2, 
	2, 2, 131, 132, 7, 36, 2, 2, 132, 48, 3, 2, 2, 2, 133, 134, 7, 94, 2, 2, 
	134, 135, 7, 36, 2, 2, 135, 50, 3, 2, 2, 2, 136, 138, 9, 3, 2, 2, 137, 
	136, 3, 2, 2, 2, 138, 52, 3, 2, 2, 2, 139, 142, 5, 51, 26, 2, 140, 142, 
	4, 50, 59, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2, 2, 2, 142, 54, 3, 2, 
	2, 2, 143, 153, 5, 57, 29, 2, 144, 147, 5, 59, 30, 2, 145, 147, 5, 61, 
	31, 2, 146, 144, 3, 2, 2, 2, 146, 145, 3, 2, 2, 2, 147, 149, 3, 2, 2, 2, 
	148, 150, 5, 63, 32, 2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 
	151, 3, 2, 2, 2, 151, 152, 5, 57, 29, 2, 152, 154, 3, 2, 2, 2, 153, 146, 
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 56, 3, 2, 2, 2, 155, 157, 4, 50, 
	59, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 
	158, 159, 3, 2, 2, 2, 159, 166, 3, 2, 2, 2, 160, 162, 7, 48, 2, 2, 161, 
	163, 4, 50, 59, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162, 
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 167, 3, 2, 2, 2, 166, 160, 3, 2, 
	2, 2, 166, 167, 3, 2, 2, 2, 167, 58, 3, 2, 2, 2, 168, 169, 7, 71, 2, 2, 
	169, 60, 3, 2, 2, 2, 170, 171, 7, 103, 2, 2, 171, 62, 3, 2, 2, 2, 172, 
	173, 9, 4, 2, 2, 173, 64, 3, 2, 2, 2, 174, 176, 9, 5, 2, 2, 175, 174, 3, 
	2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 
	2, 178, 179, 3, 2, 2, 2, 179, 180, 8, 33, 2, 2, 180, 66, 3, 2, 2, 2, 15, 
	2, 118, 124, 126, 137, 141, 146, 149, 153, 158, 164, 166, 177, 3, 8, 2, 
	2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'=='", 
	"'!='", "'||'", "'&&'", "'xor'", "','", "'.'", "'^'", "'pi'", "", "'i'", 
	"", "", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "NOT_EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", 
	"I", "VARIABLE", "QUOTED_STRING", "QUOTE", "SCIENTIFIC_NUMBER", "WS",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "NOT_EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", 
	"I", "VARIABLE", "QUOTED_STRING", "QUOTE", "ESCAPED_QUOTE", "VALID_ID_START", 
	"VALID_ID_CHAR", "SCIENTIFIC_NUMBER", "NUMBER", "E1", "E2", "SIGN", "WS",
}

type ExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExpressionLexer(input antlr.CharStream) *ExpressionLexer {

	l := new(ExpressionLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionLexer tokens.
const (
	ExpressionLexerLPAREN = 1
	ExpressionLexerRPAREN = 2
	ExpressionLexerPLUS = 3
	ExpressionLexerMINUS = 4
	ExpressionLexerTIMES = 5
	ExpressionLexerDIV = 6
	ExpressionLexerMOD = 7
	ExpressionLexerGT = 8
	ExpressionLexerLT = 9
	ExpressionLexerEQ = 10
	ExpressionLexerNOT_EQ = 11
	ExpressionLexerOR = 12
	ExpressionLexerAND = 13
	ExpressionLexerXOR = 14
	ExpressionLexerCOMMA = 15
	ExpressionLexerPOINT = 16
	ExpressionLexerPOW = 17
	ExpressionLexerPI = 18
	ExpressionLexerEULER = 19
	ExpressionLexerI = 20
	ExpressionLexerVARIABLE = 21
	ExpressionLexerQUOTED_STRING = 22
	ExpressionLexerQUOTE = 23
	ExpressionLexerSCIENTIFIC_NUMBER = 24
	ExpressionLexerWS = 25
)

