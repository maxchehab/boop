// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	6, 3, 47, 10, 3, 13, 3, 14, 3, 48, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5,
	56, 10, 5, 13, 5, 14, 5, 57, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6,
	14, 77, 10, 14, 13, 14, 14, 14, 78, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 5, 4, 2, 12, 12, 15, 15, 3, 2,
	50, 59, 4, 2, 67, 92, 99, 124, 2, 111, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2,
	5, 46, 3, 2, 2, 2, 7, 50, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 59, 3, 2,
	2, 2, 13, 61, 3, 2, 2, 2, 15, 63, 3, 2, 2, 2, 17, 65, 3, 2, 2, 2, 19, 67,
	3, 2, 2, 2, 21, 69, 3, 2, 2, 2, 23, 71, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2,
	27, 76, 3, 2, 2, 2, 29, 80, 3, 2, 2, 2, 31, 83, 3, 2, 2, 2, 33, 86, 3,
	2, 2, 2, 35, 90, 3, 2, 2, 2, 37, 95, 3, 2, 2, 2, 39, 102, 3, 2, 2, 2, 41,
	106, 3, 2, 2, 2, 43, 44, 7, 38, 2, 2, 44, 4, 3, 2, 2, 2, 45, 47, 9, 2,
	2, 2, 46, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 6, 3, 2, 2, 2, 50, 51, 7, 34, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 53, 8, 4, 2, 2, 53, 8, 3, 2, 2, 2, 54, 56, 9, 3, 2, 2, 55, 54, 3, 2,
	2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 10,
	3, 2, 2, 2, 59, 60, 7, 45, 2, 2, 60, 12, 3, 2, 2, 2, 61, 62, 7, 47, 2,
	2, 62, 14, 3, 2, 2, 2, 63, 64, 7, 44, 2, 2, 64, 16, 3, 2, 2, 2, 65, 66,
	7, 49, 2, 2, 66, 18, 3, 2, 2, 2, 67, 68, 7, 63, 2, 2, 68, 20, 3, 2, 2,
	2, 69, 70, 7, 126, 2, 2, 70, 22, 3, 2, 2, 2, 71, 72, 7, 42, 2, 2, 72, 24,
	3, 2, 2, 2, 73, 74, 7, 43, 2, 2, 74, 26, 3, 2, 2, 2, 75, 77, 9, 4, 2, 2,
	76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3,
	2, 2, 2, 79, 28, 3, 2, 2, 2, 80, 81, 7, 107, 2, 2, 81, 82, 7, 104, 2, 2,
	82, 30, 3, 2, 2, 2, 83, 84, 7, 102, 2, 2, 84, 85, 7, 113, 2, 2, 85, 32,
	3, 2, 2, 2, 86, 87, 7, 103, 2, 2, 87, 88, 7, 112, 2, 2, 88, 89, 7, 102,
	2, 2, 89, 34, 3, 2, 2, 2, 90, 91, 7, 104, 2, 2, 91, 92, 7, 119, 2, 2, 92,
	93, 7, 112, 2, 2, 93, 94, 7, 101, 2, 2, 94, 36, 3, 2, 2, 2, 95, 96, 7,
	116, 2, 2, 96, 97, 7, 103, 2, 2, 97, 98, 7, 118, 2, 2, 98, 99, 7, 119,
	2, 2, 99, 100, 7, 116, 2, 2, 100, 101, 7, 112, 2, 2, 101, 38, 3, 2, 2,
	2, 102, 103, 7, 104, 2, 2, 103, 104, 7, 113, 2, 2, 104, 105, 7, 116, 2,
	2, 105, 40, 3, 2, 2, 2, 106, 107, 7, 107, 2, 2, 107, 108, 7, 117, 2, 2,
	108, 42, 3, 2, 2, 2, 6, 2, 48, 57, 78, 3, 8, 2, 2,
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
	"", "'$'", "", "' '", "", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('",
	"')'", "", "'if'", "'do'", "'end'", "'func'", "'return'", "'for'", "'is'",
}

var lexerSymbolicNames = []string{
	"", "", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT", "DIVIDE",
	"ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END", "FUNC",
	"RETURN", "FOR", "IS",
}

var lexerRuleNames = []string{
	"T__0", "NEWLINE", "WHITESPACE", "INT", "PLUS", "MINUS", "MULT", "DIVIDE",
	"ASSIGN", "PIPE", "LPAREN", "RPAREN", "STRING", "IF", "DO", "END", "FUNC",
	"RETURN", "FOR", "IS",
}

type BeepBoopLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBeepBoopLexer(input antlr.CharStream) *BeepBoopLexer {

	l := new(BeepBoopLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeepBoop.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeepBoopLexer tokens.
const (
	BeepBoopLexerT__0       = 1
	BeepBoopLexerNEWLINE    = 2
	BeepBoopLexerWHITESPACE = 3
	BeepBoopLexerINT        = 4
	BeepBoopLexerPLUS       = 5
	BeepBoopLexerMINUS      = 6
	BeepBoopLexerMULT       = 7
	BeepBoopLexerDIVIDE     = 8
	BeepBoopLexerASSIGN     = 9
	BeepBoopLexerPIPE       = 10
	BeepBoopLexerLPAREN     = 11
	BeepBoopLexerRPAREN     = 12
	BeepBoopLexerSTRING     = 13
	BeepBoopLexerIF         = 14
	BeepBoopLexerDO         = 15
	BeepBoopLexerEND        = 16
	BeepBoopLexerFUNC       = 17
	BeepBoopLexerRETURN     = 18
	BeepBoopLexerFOR        = 19
	BeepBoopLexerIS         = 20
)
