// Code generated from MonkeyLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MonkeyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MonkeyLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monkeylexerLexerInit() {
	staticData := &MonkeyLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'fn'", "'let'", "'true'", "'false'", "'if'", "'else'", "'return'",
		"", "", "", "'=='", "'!='", "'+'", "'-'", "'!'", "'*'", "'/'", "'<'",
		"'>'", "'='", "','", "';'", "':'", "'('", "')'", "'{'", "'}'", "'['",
		"']'",
	}
	staticData.SymbolicNames = []string{
		"", "FUNCTION", "LET", "TRUE", "FALSE", "IF", "ELSE", "RETURN", "IDENT",
		"INTEGER", "STRING", "EQ_EQ", "BANG_EQ", "PLUS", "MINUS", "BANG", "STAR",
		"SLASH", "LT", "GT", "EQ", "COMMA", "SEMICOLON", "COLON", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "WS",
	}
	staticData.RuleNames = []string{
		"FUNCTION", "LET", "TRUE", "FALSE", "IF", "ELSE", "RETURN", "IDENT",
		"INTEGER", "STRING", "STRING_CHAR", "EQ_EQ", "BANG_EQ", "PLUS", "MINUS",
		"BANG", "STAR", "SLASH", "LT", "GT", "EQ", "COMMA", "SEMICOLON", "COLON",
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 164, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 4, 7, 98, 8, 7,
		11, 7, 12, 7, 99, 1, 8, 4, 8, 103, 8, 8, 11, 8, 12, 8, 104, 1, 9, 1, 9,
		5, 9, 109, 8, 9, 10, 9, 12, 9, 112, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 4, 30, 159,
		8, 30, 11, 30, 12, 30, 160, 1, 30, 1, 30, 0, 0, 31, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 11, 25, 12, 27,
		13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45,
		22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 1,
		0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0,
		9, 10, 13, 13, 32, 32, 166, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 1, 63, 1, 0, 0, 0, 3, 66, 1, 0, 0, 0, 5, 70, 1, 0, 0, 0, 7,
		75, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0, 11, 84, 1, 0, 0, 0, 13, 89, 1, 0, 0,
		0, 15, 97, 1, 0, 0, 0, 17, 102, 1, 0, 0, 0, 19, 106, 1, 0, 0, 0, 21, 115,
		1, 0, 0, 0, 23, 117, 1, 0, 0, 0, 25, 120, 1, 0, 0, 0, 27, 123, 1, 0, 0,
		0, 29, 125, 1, 0, 0, 0, 31, 127, 1, 0, 0, 0, 33, 129, 1, 0, 0, 0, 35, 131,
		1, 0, 0, 0, 37, 133, 1, 0, 0, 0, 39, 135, 1, 0, 0, 0, 41, 137, 1, 0, 0,
		0, 43, 139, 1, 0, 0, 0, 45, 141, 1, 0, 0, 0, 47, 143, 1, 0, 0, 0, 49, 145,
		1, 0, 0, 0, 51, 147, 1, 0, 0, 0, 53, 149, 1, 0, 0, 0, 55, 151, 1, 0, 0,
		0, 57, 153, 1, 0, 0, 0, 59, 155, 1, 0, 0, 0, 61, 158, 1, 0, 0, 0, 63, 64,
		5, 102, 0, 0, 64, 65, 5, 110, 0, 0, 65, 2, 1, 0, 0, 0, 66, 67, 5, 108,
		0, 0, 67, 68, 5, 101, 0, 0, 68, 69, 5, 116, 0, 0, 69, 4, 1, 0, 0, 0, 70,
		71, 5, 116, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 117, 0, 0, 73, 74, 5,
		101, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 102, 0, 0, 76, 77, 5, 97, 0, 0,
		77, 78, 5, 108, 0, 0, 78, 79, 5, 115, 0, 0, 79, 80, 5, 101, 0, 0, 80, 8,
		1, 0, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83, 5, 102, 0, 0, 83, 10, 1, 0, 0,
		0, 84, 85, 5, 101, 0, 0, 85, 86, 5, 108, 0, 0, 86, 87, 5, 115, 0, 0, 87,
		88, 5, 101, 0, 0, 88, 12, 1, 0, 0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5,
		101, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 117, 0, 0, 93, 94, 5, 114,
		0, 0, 94, 95, 5, 110, 0, 0, 95, 14, 1, 0, 0, 0, 96, 98, 7, 0, 0, 0, 97,
		96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0,
		0, 0, 100, 16, 1, 0, 0, 0, 101, 103, 7, 1, 0, 0, 102, 101, 1, 0, 0, 0,
		103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105,
		18, 1, 0, 0, 0, 106, 110, 7, 2, 0, 0, 107, 109, 3, 21, 10, 0, 108, 107,
		1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114, 7, 2, 0, 0,
		114, 20, 1, 0, 0, 0, 115, 116, 8, 2, 0, 0, 116, 22, 1, 0, 0, 0, 117, 118,
		5, 61, 0, 0, 118, 119, 5, 61, 0, 0, 119, 24, 1, 0, 0, 0, 120, 121, 5, 33,
		0, 0, 121, 122, 5, 61, 0, 0, 122, 26, 1, 0, 0, 0, 123, 124, 5, 43, 0, 0,
		124, 28, 1, 0, 0, 0, 125, 126, 5, 45, 0, 0, 126, 30, 1, 0, 0, 0, 127, 128,
		5, 33, 0, 0, 128, 32, 1, 0, 0, 0, 129, 130, 5, 42, 0, 0, 130, 34, 1, 0,
		0, 0, 131, 132, 5, 47, 0, 0, 132, 36, 1, 0, 0, 0, 133, 134, 5, 60, 0, 0,
		134, 38, 1, 0, 0, 0, 135, 136, 5, 62, 0, 0, 136, 40, 1, 0, 0, 0, 137, 138,
		5, 61, 0, 0, 138, 42, 1, 0, 0, 0, 139, 140, 5, 44, 0, 0, 140, 44, 1, 0,
		0, 0, 141, 142, 5, 59, 0, 0, 142, 46, 1, 0, 0, 0, 143, 144, 5, 58, 0, 0,
		144, 48, 1, 0, 0, 0, 145, 146, 5, 40, 0, 0, 146, 50, 1, 0, 0, 0, 147, 148,
		5, 41, 0, 0, 148, 52, 1, 0, 0, 0, 149, 150, 5, 123, 0, 0, 150, 54, 1, 0,
		0, 0, 151, 152, 5, 125, 0, 0, 152, 56, 1, 0, 0, 0, 153, 154, 5, 91, 0,
		0, 154, 58, 1, 0, 0, 0, 155, 156, 5, 93, 0, 0, 156, 60, 1, 0, 0, 0, 157,
		159, 7, 3, 0, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 6, 30,
		0, 0, 163, 62, 1, 0, 0, 0, 5, 0, 99, 104, 110, 160, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MonkeyLexerInit initializes any static state used to implement MonkeyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMonkeyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonkeyLexerInit() {
	staticData := &MonkeyLexerLexerStaticData
	staticData.once.Do(monkeylexerLexerInit)
}

// NewMonkeyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMonkeyLexer(input antlr.CharStream) *MonkeyLexer {
	MonkeyLexerInit()
	l := new(MonkeyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MonkeyLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MonkeyLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MonkeyLexer tokens.
const (
	MonkeyLexerFUNCTION  = 1
	MonkeyLexerLET       = 2
	MonkeyLexerTRUE      = 3
	MonkeyLexerFALSE     = 4
	MonkeyLexerIF        = 5
	MonkeyLexerELSE      = 6
	MonkeyLexerRETURN    = 7
	MonkeyLexerIDENT     = 8
	MonkeyLexerINTEGER   = 9
	MonkeyLexerSTRING    = 10
	MonkeyLexerEQ_EQ     = 11
	MonkeyLexerBANG_EQ   = 12
	MonkeyLexerPLUS      = 13
	MonkeyLexerMINUS     = 14
	MonkeyLexerBANG      = 15
	MonkeyLexerSTAR      = 16
	MonkeyLexerSLASH     = 17
	MonkeyLexerLT        = 18
	MonkeyLexerGT        = 19
	MonkeyLexerEQ        = 20
	MonkeyLexerCOMMA     = 21
	MonkeyLexerSEMICOLON = 22
	MonkeyLexerCOLON     = 23
	MonkeyLexerLPAREN    = 24
	MonkeyLexerRPAREN    = 25
	MonkeyLexerLBRACE    = 26
	MonkeyLexerRBRACE    = 27
	MonkeyLexerLBRACKET  = 28
	MonkeyLexerRBRACKET  = 29
	MonkeyLexerWS        = 30
)
