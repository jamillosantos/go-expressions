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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 36, 238, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 2, 
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 
	20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 
	28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 7, 31, 174, 10, 31, 
	12, 31, 14, 31, 177, 11, 31, 3, 32, 3, 32, 3, 32, 7, 32, 182, 10, 32, 12, 
	32, 14, 32, 185, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 
	3, 35, 5, 35, 195, 10, 35, 3, 36, 3, 36, 5, 36, 199, 10, 36, 3, 37, 3, 
	37, 3, 37, 5, 37, 204, 10, 37, 3, 37, 5, 37, 207, 10, 37, 3, 37, 3, 37, 
	5, 37, 211, 10, 37, 3, 38, 6, 38, 214, 10, 38, 13, 38, 14, 38, 215, 3, 
	38, 3, 38, 6, 38, 220, 10, 38, 13, 38, 14, 38, 221, 5, 38, 224, 10, 38, 
	3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 6, 42, 233, 10, 42, 13, 
	42, 14, 42, 234, 3, 42, 3, 42, 3, 183, 2, 43, 3, 3, 5, 4, 7, 5, 9, 6, 11, 
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 
	67, 2, 69, 2, 71, 2, 73, 35, 75, 2, 77, 2, 79, 2, 81, 2, 83, 36, 3, 2, 
	6, 4, 2, 12, 12, 15, 15, 5, 2, 67, 92, 97, 97, 99, 124, 4, 2, 45, 45, 47, 
	47, 5, 2, 11, 12, 15, 15, 34, 34, 2, 241, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 73, 
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 3, 85, 3, 2, 2, 2, 5, 89, 3, 2, 2, 2, 7, 
	93, 3, 2, 2, 2, 9, 97, 3, 2, 2, 2, 11, 102, 3, 2, 2, 2, 13, 107, 3, 2, 
	2, 2, 15, 112, 3, 2, 2, 2, 17, 115, 3, 2, 2, 2, 19, 119, 3, 2, 2, 2, 21, 
	124, 3, 2, 2, 2, 23, 127, 3, 2, 2, 2, 25, 129, 3, 2, 2, 2, 27, 131, 3, 
	2, 2, 2, 29, 133, 3, 2, 2, 2, 31, 135, 3, 2, 2, 2, 33, 137, 3, 2, 2, 2, 
	35, 139, 3, 2, 2, 2, 37, 141, 3, 2, 2, 2, 39, 143, 3, 2, 2, 2, 41, 145, 
	3, 2, 2, 2, 43, 148, 3, 2, 2, 2, 45, 151, 3, 2, 2, 2, 47, 154, 3, 2, 2, 
	2, 49, 158, 3, 2, 2, 2, 51, 160, 3, 2, 2, 2, 53, 162, 3, 2, 2, 2, 55, 164, 
	3, 2, 2, 2, 57, 167, 3, 2, 2, 2, 59, 169, 3, 2, 2, 2, 61, 171, 3, 2, 2, 
	2, 63, 178, 3, 2, 2, 2, 65, 188, 3, 2, 2, 2, 67, 190, 3, 2, 2, 2, 69, 194, 
	3, 2, 2, 2, 71, 198, 3, 2, 2, 2, 73, 200, 3, 2, 2, 2, 75, 213, 3, 2, 2, 
	2, 77, 225, 3, 2, 2, 2, 79, 227, 3, 2, 2, 2, 81, 229, 3, 2, 2, 2, 83, 232, 
	3, 2, 2, 2, 85, 86, 7, 101, 2, 2, 86, 87, 7, 113, 2, 2, 87, 88, 7, 117, 
	2, 2, 88, 4, 3, 2, 2, 2, 89, 90, 7, 117, 2, 2, 90, 91, 7, 107, 2, 2, 91, 
	92, 7, 112, 2, 2, 92, 6, 3, 2, 2, 2, 93, 94, 7, 118, 2, 2, 94, 95, 7, 99, 
	2, 2, 95, 96, 7, 112, 2, 2, 96, 8, 3, 2, 2, 2, 97, 98, 7, 99, 2, 2, 98, 
	99, 7, 101, 2, 2, 99, 100, 7, 113, 2, 2, 100, 101, 7, 117, 2, 2, 101, 10, 
	3, 2, 2, 2, 102, 103, 7, 99, 2, 2, 103, 104, 7, 117, 2, 2, 104, 105, 7, 
	107, 2, 2, 105, 106, 7, 112, 2, 2, 106, 12, 3, 2, 2, 2, 107, 108, 7, 99, 
	2, 2, 108, 109, 7, 118, 2, 2, 109, 110, 7, 99, 2, 2, 110, 111, 7, 112, 
	2, 2, 111, 14, 3, 2, 2, 2, 112, 113, 7, 110, 2, 2, 113, 114, 7, 112, 2, 
	2, 114, 16, 3, 2, 2, 2, 115, 116, 7, 110, 2, 2, 116, 117, 7, 113, 2, 2, 
	117, 118, 7, 105, 2, 2, 118, 18, 3, 2, 2, 2, 119, 120, 7, 117, 2, 2, 120, 
	121, 7, 115, 2, 2, 121, 122, 7, 116, 2, 2, 122, 123, 7, 118, 2, 2, 123, 
	20, 3, 2, 2, 2, 124, 125, 7, 107, 2, 2, 125, 126, 7, 104, 2, 2, 126, 22, 
	3, 2, 2, 2, 127, 128, 7, 42, 2, 2, 128, 24, 3, 2, 2, 2, 129, 130, 7, 43, 
	2, 2, 130, 26, 3, 2, 2, 2, 131, 132, 7, 45, 2, 2, 132, 28, 3, 2, 2, 2, 
	133, 134, 7, 47, 2, 2, 134, 30, 3, 2, 2, 2, 135, 136, 7, 44, 2, 2, 136, 
	32, 3, 2, 2, 2, 137, 138, 7, 49, 2, 2, 138, 34, 3, 2, 2, 2, 139, 140, 7, 
	39, 2, 2, 140, 36, 3, 2, 2, 2, 141, 142, 7, 64, 2, 2, 142, 38, 3, 2, 2, 
	2, 143, 144, 7, 62, 2, 2, 144, 40, 3, 2, 2, 2, 145, 146, 7, 63, 2, 2, 146, 
	147, 7, 63, 2, 2, 147, 42, 3, 2, 2, 2, 148, 149, 7, 126, 2, 2, 149, 150, 
	7, 126, 2, 2, 150, 44, 3, 2, 2, 2, 151, 152, 7, 40, 2, 2, 152, 153, 7, 
	40, 2, 2, 153, 46, 3, 2, 2, 2, 154, 155, 7, 122, 2, 2, 155, 156, 7, 113, 
	2, 2, 156, 157, 7, 116, 2, 2, 157, 48, 3, 2, 2, 2, 158, 159, 7, 46, 2, 
	2, 159, 50, 3, 2, 2, 2, 160, 161, 7, 48, 2, 2, 161, 52, 3, 2, 2, 2, 162, 
	163, 7, 96, 2, 2, 163, 54, 3, 2, 2, 2, 164, 165, 7, 114, 2, 2, 165, 166, 
	7, 107, 2, 2, 166, 56, 3, 2, 2, 2, 167, 168, 5, 79, 40, 2, 168, 58, 3, 
	2, 2, 2, 169, 170, 7, 107, 2, 2, 170, 60, 3, 2, 2, 2, 171, 175, 5, 69, 
	35, 2, 172, 174, 5, 71, 36, 2, 173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 
	2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 62, 3, 2, 2, 2, 177, 
	175, 3, 2, 2, 2, 178, 183, 5, 65, 33, 2, 179, 182, 5, 67, 34, 2, 180, 182, 
	10, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 185, 3, 2, 
	2, 2, 183, 184, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 
	185, 183, 3, 2, 2, 2, 186, 187, 5, 65, 33, 2, 187, 64, 3, 2, 2, 2, 188, 
	189, 7, 36, 2, 2, 189, 66, 3, 2, 2, 2, 190, 191, 7, 94, 2, 2, 191, 192, 
	7, 36, 2, 2, 192, 68, 3, 2, 2, 2, 193, 195, 9, 3, 2, 2, 194, 193, 3, 2, 
	2, 2, 195, 70, 3, 2, 2, 2, 196, 199, 5, 69, 35, 2, 197, 199, 4, 50, 59, 
	2, 198, 196, 3, 2, 2, 2, 198, 197, 3, 2, 2, 2, 199, 72, 3, 2, 2, 2, 200, 
	210, 5, 75, 38, 2, 201, 204, 5, 77, 39, 2, 202, 204, 5, 79, 40, 2, 203, 
	201, 3, 2, 2, 2, 203, 202, 3, 2, 2, 2, 204, 206, 3, 2, 2, 2, 205, 207, 
	5, 81, 41, 2, 206, 205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 
	2, 2, 2, 208, 209, 5, 75, 38, 2, 209, 211, 3, 2, 2, 2, 210, 203, 3, 2, 
	2, 2, 210, 211, 3, 2, 2, 2, 211, 74, 3, 2, 2, 2, 212, 214, 4, 50, 59, 2, 
	213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 
	216, 3, 2, 2, 2, 216, 223, 3, 2, 2, 2, 217, 219, 7, 48, 2, 2, 218, 220, 
	4, 50, 59, 2, 219, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 219, 3, 
	2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 3, 2, 2, 2, 223, 217, 3, 2, 2, 
	2, 223, 224, 3, 2, 2, 2, 224, 76, 3, 2, 2, 2, 225, 226, 7, 71, 2, 2, 226, 
	78, 3, 2, 2, 2, 227, 228, 7, 103, 2, 2, 228, 80, 3, 2, 2, 2, 229, 230, 
	9, 4, 2, 2, 230, 82, 3, 2, 2, 2, 231, 233, 9, 5, 2, 2, 232, 231, 3, 2, 
	2, 2, 233, 234, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 
	235, 236, 3, 2, 2, 2, 236, 237, 8, 42, 2, 2, 237, 84, 3, 2, 2, 2, 15, 2, 
	175, 181, 183, 194, 198, 203, 206, 210, 215, 221, 223, 234, 3, 8, 2, 2,
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
	"", "'cos'", "'sin'", "'tan'", "'acos'", "'asin'", "'atan'", "'ln'", "'log'", 
	"'sqrt'", "'if'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", 
	"'<'", "'=='", "'||'", "'&&'", "'xor'", "','", "'.'", "'^'", "'pi'", "", 
	"'i'", "", "", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "IF", 
	"LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", "I", 
	"VARIABLE", "QUOTED_STRING", "QUOTE", "SCIENTIFIC_NUMBER", "WS",
}

var lexerRuleNames = []string{
	"COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "IF", 
	"LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "MOD", "GT", "LT", 
	"EQ", "OR", "AND", "XOR", "COMMA", "POINT", "POW", "PI", "EULER", "I", 
	"VARIABLE", "QUOTED_STRING", "QUOTE", "ESCAPED_QUOTE", "VALID_ID_START", 
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
	ExpressionLexerCOS = 1
	ExpressionLexerSIN = 2
	ExpressionLexerTAN = 3
	ExpressionLexerACOS = 4
	ExpressionLexerASIN = 5
	ExpressionLexerATAN = 6
	ExpressionLexerLN = 7
	ExpressionLexerLOG = 8
	ExpressionLexerSQRT = 9
	ExpressionLexerIF = 10
	ExpressionLexerLPAREN = 11
	ExpressionLexerRPAREN = 12
	ExpressionLexerPLUS = 13
	ExpressionLexerMINUS = 14
	ExpressionLexerTIMES = 15
	ExpressionLexerDIV = 16
	ExpressionLexerMOD = 17
	ExpressionLexerGT = 18
	ExpressionLexerLT = 19
	ExpressionLexerEQ = 20
	ExpressionLexerOR = 21
	ExpressionLexerAND = 22
	ExpressionLexerXOR = 23
	ExpressionLexerCOMMA = 24
	ExpressionLexerPOINT = 25
	ExpressionLexerPOW = 26
	ExpressionLexerPI = 27
	ExpressionLexerEULER = 28
	ExpressionLexerI = 29
	ExpressionLexerVARIABLE = 30
	ExpressionLexerQUOTED_STRING = 31
	ExpressionLexerQUOTE = 32
	ExpressionLexerSCIENTIFIC_NUMBER = 33
	ExpressionLexerWS = 34
)

