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
		"FUNCTION", "LET", "TRUE", "FALSE", "IF", "ELSE", "RETURN", "ALPHA",
		"IDENT", "DIGIT", "INTEGER", "STRING_CHAR", "STRING", "EQ_EQ", "BANG_EQ",
		"PLUS", "MINUS", "BANG", "STAR", "SLASH", "LT", "GT", "EQ", "COMMA",
		"SEMICOLON", "COLON", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET",
		"RBRACKET", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 172, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 104, 8, 8, 11, 8, 12, 8, 105, 1, 9, 1, 9,
		1, 10, 4, 10, 111, 8, 10, 11, 10, 12, 10, 112, 1, 11, 1, 11, 1, 12, 1,
		12, 5, 12, 119, 8, 12, 10, 12, 12, 12, 122, 9, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 4, 32, 167,
		8, 32, 11, 32, 12, 32, 168, 1, 32, 1, 32, 0, 0, 33, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 0, 17, 8, 19, 0, 21, 9, 23, 0, 25, 10, 27, 11,
		29, 12, 31, 13, 33, 14, 35, 15, 37, 16, 39, 17, 41, 18, 43, 19, 45, 20,
		47, 21, 49, 22, 51, 23, 53, 24, 55, 25, 57, 26, 59, 27, 61, 28, 63, 29,
		65, 30, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 34,
		34, 3, 0, 9, 10, 13, 13, 32, 32, 172, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 5, 74, 1, 0,
		0, 0, 7, 79, 1, 0, 0, 0, 9, 85, 1, 0, 0, 0, 11, 88, 1, 0, 0, 0, 13, 93,
		1, 0, 0, 0, 15, 100, 1, 0, 0, 0, 17, 103, 1, 0, 0, 0, 19, 107, 1, 0, 0,
		0, 21, 110, 1, 0, 0, 0, 23, 114, 1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 125,
		1, 0, 0, 0, 29, 128, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 133, 1, 0, 0,
		0, 35, 135, 1, 0, 0, 0, 37, 137, 1, 0, 0, 0, 39, 139, 1, 0, 0, 0, 41, 141,
		1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 145, 1, 0, 0, 0, 47, 147, 1, 0, 0,
		0, 49, 149, 1, 0, 0, 0, 51, 151, 1, 0, 0, 0, 53, 153, 1, 0, 0, 0, 55, 155,
		1, 0, 0, 0, 57, 157, 1, 0, 0, 0, 59, 159, 1, 0, 0, 0, 61, 161, 1, 0, 0,
		0, 63, 163, 1, 0, 0, 0, 65, 166, 1, 0, 0, 0, 67, 68, 5, 102, 0, 0, 68,
		69, 5, 110, 0, 0, 69, 2, 1, 0, 0, 0, 70, 71, 5, 108, 0, 0, 71, 72, 5, 101,
		0, 0, 72, 73, 5, 116, 0, 0, 73, 4, 1, 0, 0, 0, 74, 75, 5, 116, 0, 0, 75,
		76, 5, 114, 0, 0, 76, 77, 5, 117, 0, 0, 77, 78, 5, 101, 0, 0, 78, 6, 1,
		0, 0, 0, 79, 80, 5, 102, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5, 108, 0,
		0, 82, 83, 5, 115, 0, 0, 83, 84, 5, 101, 0, 0, 84, 8, 1, 0, 0, 0, 85, 86,
		5, 105, 0, 0, 86, 87, 5, 102, 0, 0, 87, 10, 1, 0, 0, 0, 88, 89, 5, 101,
		0, 0, 89, 90, 5, 108, 0, 0, 90, 91, 5, 115, 0, 0, 91, 92, 5, 101, 0, 0,
		92, 12, 1, 0, 0, 0, 93, 94, 5, 114, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96,
		5, 116, 0, 0, 96, 97, 5, 117, 0, 0, 97, 98, 5, 114, 0, 0, 98, 99, 5, 110,
		0, 0, 99, 14, 1, 0, 0, 0, 100, 101, 7, 0, 0, 0, 101, 16, 1, 0, 0, 0, 102,
		104, 3, 15, 7, 0, 103, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 18, 1, 0, 0, 0, 107, 108, 7, 1,
		0, 0, 108, 20, 1, 0, 0, 0, 109, 111, 3, 19, 9, 0, 110, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		22, 1, 0, 0, 0, 114, 115, 8, 2, 0, 0, 115, 24, 1, 0, 0, 0, 116, 120, 7,
		2, 0, 0, 117, 119, 3, 23, 11, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0,
		0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0,
		122, 120, 1, 0, 0, 0, 123, 124, 7, 2, 0, 0, 124, 26, 1, 0, 0, 0, 125, 126,
		5, 61, 0, 0, 126, 127, 5, 61, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 5, 33,
		0, 0, 129, 130, 5, 61, 0, 0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 43, 0, 0,
		132, 32, 1, 0, 0, 0, 133, 134, 5, 45, 0, 0, 134, 34, 1, 0, 0, 0, 135, 136,
		5, 33, 0, 0, 136, 36, 1, 0, 0, 0, 137, 138, 5, 42, 0, 0, 138, 38, 1, 0,
		0, 0, 139, 140, 5, 47, 0, 0, 140, 40, 1, 0, 0, 0, 141, 142, 5, 60, 0, 0,
		142, 42, 1, 0, 0, 0, 143, 144, 5, 62, 0, 0, 144, 44, 1, 0, 0, 0, 145, 146,
		5, 61, 0, 0, 146, 46, 1, 0, 0, 0, 147, 148, 5, 44, 0, 0, 148, 48, 1, 0,
		0, 0, 149, 150, 5, 59, 0, 0, 150, 50, 1, 0, 0, 0, 151, 152, 5, 58, 0, 0,
		152, 52, 1, 0, 0, 0, 153, 154, 5, 40, 0, 0, 154, 54, 1, 0, 0, 0, 155, 156,
		5, 41, 0, 0, 156, 56, 1, 0, 0, 0, 157, 158, 5, 123, 0, 0, 158, 58, 1, 0,
		0, 0, 159, 160, 5, 125, 0, 0, 160, 60, 1, 0, 0, 0, 161, 162, 5, 91, 0,
		0, 162, 62, 1, 0, 0, 0, 163, 164, 5, 93, 0, 0, 164, 64, 1, 0, 0, 0, 165,
		167, 7, 3, 0, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 6, 32,
		0, 0, 171, 66, 1, 0, 0, 0, 5, 0, 105, 112, 120, 168, 1, 6, 0, 0,
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
