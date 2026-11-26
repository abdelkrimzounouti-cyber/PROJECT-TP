// Code generated from Pcg.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// we don't actually need this

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PcgLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PcgLexerLexerStaticData struct {
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

func pcglexerLexerInit() {
	staticData := &PcgLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'program'", "'begin'", "'end'", "'write'", "'read'",
		"'else'", "'if'", "'var'", "'while'", "'then'", "'do'", "'writeln'",
		"'true'", "'false'", "'not'", "'and'", "'or'", "'<-'", "'<='", "'>='",
		"'='", "'<>'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "':'", "'<'",
		"'>'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "PROGRAM", "BEGIN",
		"END", "WRITE", "READ", "ELSE", "IF", "VAR", "WHILE", "THEN", "DO",
		"WRITELN", "TRUE", "FALSE", "NOT", "AND", "OR", "ASSIGN", "LE", "GE",
		"EQ", "NE", "ADD", "SUB", "MUL", "DIV", "LPAR", "RPAR", "COLON", "LT",
		"GT", "NUMBER", "ID",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT", "PROGRAM", "BEGIN", "END",
		"WRITE", "READ", "ELSE", "IF", "VAR", "WHILE", "THEN", "DO", "WRITELN",
		"TRUE", "FALSE", "NOT", "AND", "OR", "ASSIGN", "LE", "GE", "EQ", "NE",
		"ADD", "SUB", "MUL", "DIV", "LPAR", "RPAR", "COLON", "LT", "GT", "DIGIT",
		"NUMBER", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 242, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 4, 0, 77, 8, 0, 11, 0, 12, 0, 78, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 88, 8, 1, 10, 1, 12, 1, 91, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 97, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 105,
		8, 2, 10, 2, 12, 2, 108, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 4, 35, 232, 8, 35, 11, 35, 12,
		35, 233, 1, 36, 1, 36, 5, 36, 238, 8, 36, 10, 36, 12, 36, 241, 9, 36, 1,
		89, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 0, 71, 35, 73,
		36, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48,
		57, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 247,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 1, 76, 1, 0, 0, 0, 3, 82, 1, 0, 0, 0, 5, 100,
		1, 0, 0, 0, 7, 111, 1, 0, 0, 0, 9, 119, 1, 0, 0, 0, 11, 125, 1, 0, 0, 0,
		13, 129, 1, 0, 0, 0, 15, 135, 1, 0, 0, 0, 17, 140, 1, 0, 0, 0, 19, 145,
		1, 0, 0, 0, 21, 148, 1, 0, 0, 0, 23, 152, 1, 0, 0, 0, 25, 158, 1, 0, 0,
		0, 27, 163, 1, 0, 0, 0, 29, 166, 1, 0, 0, 0, 31, 174, 1, 0, 0, 0, 33, 179,
		1, 0, 0, 0, 35, 185, 1, 0, 0, 0, 37, 189, 1, 0, 0, 0, 39, 193, 1, 0, 0,
		0, 41, 196, 1, 0, 0, 0, 43, 199, 1, 0, 0, 0, 45, 202, 1, 0, 0, 0, 47, 205,
		1, 0, 0, 0, 49, 207, 1, 0, 0, 0, 51, 210, 1, 0, 0, 0, 53, 212, 1, 0, 0,
		0, 55, 214, 1, 0, 0, 0, 57, 216, 1, 0, 0, 0, 59, 218, 1, 0, 0, 0, 61, 220,
		1, 0, 0, 0, 63, 222, 1, 0, 0, 0, 65, 224, 1, 0, 0, 0, 67, 226, 1, 0, 0,
		0, 69, 228, 1, 0, 0, 0, 71, 231, 1, 0, 0, 0, 73, 235, 1, 0, 0, 0, 75, 77,
		7, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 6, 0, 0, 0, 81, 2, 1, 0,
		0, 0, 82, 83, 5, 40, 0, 0, 83, 84, 5, 42, 0, 0, 84, 89, 1, 0, 0, 0, 85,
		88, 3, 3, 1, 0, 86, 88, 9, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0,
		0, 88, 91, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 96,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 42, 0, 0, 93, 97, 5, 41, 0,
		0, 94, 95, 5, 0, 0, 1, 95, 97, 6, 1, 1, 0, 96, 92, 1, 0, 0, 0, 96, 94,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 6, 1, 0, 0, 99, 4, 1, 0, 0, 0,
		100, 101, 5, 45, 0, 0, 101, 102, 5, 45, 0, 0, 102, 106, 1, 0, 0, 0, 103,
		105, 8, 1, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0,
		0, 0, 109, 110, 6, 2, 0, 0, 110, 6, 1, 0, 0, 0, 111, 112, 5, 112, 0, 0,
		112, 113, 5, 114, 0, 0, 113, 114, 5, 111, 0, 0, 114, 115, 5, 103, 0, 0,
		115, 116, 5, 114, 0, 0, 116, 117, 5, 97, 0, 0, 117, 118, 5, 109, 0, 0,
		118, 8, 1, 0, 0, 0, 119, 120, 5, 98, 0, 0, 120, 121, 5, 101, 0, 0, 121,
		122, 5, 103, 0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 110, 0, 0, 124,
		10, 1, 0, 0, 0, 125, 126, 5, 101, 0, 0, 126, 127, 5, 110, 0, 0, 127, 128,
		5, 100, 0, 0, 128, 12, 1, 0, 0, 0, 129, 130, 5, 119, 0, 0, 130, 131, 5,
		114, 0, 0, 131, 132, 5, 105, 0, 0, 132, 133, 5, 116, 0, 0, 133, 134, 5,
		101, 0, 0, 134, 14, 1, 0, 0, 0, 135, 136, 5, 114, 0, 0, 136, 137, 5, 101,
		0, 0, 137, 138, 5, 97, 0, 0, 138, 139, 5, 100, 0, 0, 139, 16, 1, 0, 0,
		0, 140, 141, 5, 101, 0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 115, 0,
		0, 143, 144, 5, 101, 0, 0, 144, 18, 1, 0, 0, 0, 145, 146, 5, 105, 0, 0,
		146, 147, 5, 102, 0, 0, 147, 20, 1, 0, 0, 0, 148, 149, 5, 118, 0, 0, 149,
		150, 5, 97, 0, 0, 150, 151, 5, 114, 0, 0, 151, 22, 1, 0, 0, 0, 152, 153,
		5, 119, 0, 0, 153, 154, 5, 104, 0, 0, 154, 155, 5, 105, 0, 0, 155, 156,
		5, 108, 0, 0, 156, 157, 5, 101, 0, 0, 157, 24, 1, 0, 0, 0, 158, 159, 5,
		116, 0, 0, 159, 160, 5, 104, 0, 0, 160, 161, 5, 101, 0, 0, 161, 162, 5,
		110, 0, 0, 162, 26, 1, 0, 0, 0, 163, 164, 5, 100, 0, 0, 164, 165, 5, 111,
		0, 0, 165, 28, 1, 0, 0, 0, 166, 167, 5, 119, 0, 0, 167, 168, 5, 114, 0,
		0, 168, 169, 5, 105, 0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 101, 0,
		0, 171, 172, 5, 108, 0, 0, 172, 173, 5, 110, 0, 0, 173, 30, 1, 0, 0, 0,
		174, 175, 5, 116, 0, 0, 175, 176, 5, 114, 0, 0, 176, 177, 5, 117, 0, 0,
		177, 178, 5, 101, 0, 0, 178, 32, 1, 0, 0, 0, 179, 180, 5, 102, 0, 0, 180,
		181, 5, 97, 0, 0, 181, 182, 5, 108, 0, 0, 182, 183, 5, 115, 0, 0, 183,
		184, 5, 101, 0, 0, 184, 34, 1, 0, 0, 0, 185, 186, 5, 110, 0, 0, 186, 187,
		5, 111, 0, 0, 187, 188, 5, 116, 0, 0, 188, 36, 1, 0, 0, 0, 189, 190, 5,
		97, 0, 0, 190, 191, 5, 110, 0, 0, 191, 192, 5, 100, 0, 0, 192, 38, 1, 0,
		0, 0, 193, 194, 5, 111, 0, 0, 194, 195, 5, 114, 0, 0, 195, 40, 1, 0, 0,
		0, 196, 197, 5, 60, 0, 0, 197, 198, 5, 45, 0, 0, 198, 42, 1, 0, 0, 0, 199,
		200, 5, 60, 0, 0, 200, 201, 5, 61, 0, 0, 201, 44, 1, 0, 0, 0, 202, 203,
		5, 62, 0, 0, 203, 204, 5, 61, 0, 0, 204, 46, 1, 0, 0, 0, 205, 206, 5, 61,
		0, 0, 206, 48, 1, 0, 0, 0, 207, 208, 5, 60, 0, 0, 208, 209, 5, 62, 0, 0,
		209, 50, 1, 0, 0, 0, 210, 211, 5, 43, 0, 0, 211, 52, 1, 0, 0, 0, 212, 213,
		5, 45, 0, 0, 213, 54, 1, 0, 0, 0, 214, 215, 5, 42, 0, 0, 215, 56, 1, 0,
		0, 0, 216, 217, 5, 47, 0, 0, 217, 58, 1, 0, 0, 0, 218, 219, 5, 40, 0, 0,
		219, 60, 1, 0, 0, 0, 220, 221, 5, 41, 0, 0, 221, 62, 1, 0, 0, 0, 222, 223,
		5, 58, 0, 0, 223, 64, 1, 0, 0, 0, 224, 225, 5, 60, 0, 0, 225, 66, 1, 0,
		0, 0, 226, 227, 5, 62, 0, 0, 227, 68, 1, 0, 0, 0, 228, 229, 7, 2, 0, 0,
		229, 70, 1, 0, 0, 0, 230, 232, 3, 69, 34, 0, 231, 230, 1, 0, 0, 0, 232,
		233, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 72, 1,
		0, 0, 0, 235, 239, 7, 3, 0, 0, 236, 238, 7, 4, 0, 0, 237, 236, 1, 0, 0,
		0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240,
		74, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 8, 0, 78, 87, 89, 96, 106, 233, 239,
		2, 6, 0, 0, 1, 1, 0,
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

// PcgLexerInit initializes any static state used to implement PcgLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPcgLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PcgLexerInit() {
	staticData := &PcgLexerLexerStaticData
	staticData.once.Do(pcglexerLexerInit)
}

// NewPcgLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPcgLexer(input antlr.CharStream) *PcgLexer {
	PcgLexerInit()
	l := new(PcgLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PcgLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Pcg.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PcgLexer tokens.
const (
	PcgLexerWHITESPACE    = 1
	PcgLexerBLOCK_COMMENT = 2
	PcgLexerLINE_COMMENT  = 3
	PcgLexerPROGRAM       = 4
	PcgLexerBEGIN         = 5
	PcgLexerEND           = 6
	PcgLexerWRITE         = 7
	PcgLexerREAD          = 8
	PcgLexerELSE          = 9
	PcgLexerIF            = 10
	PcgLexerVAR           = 11
	PcgLexerWHILE         = 12
	PcgLexerTHEN          = 13
	PcgLexerDO            = 14
	PcgLexerWRITELN       = 15
	PcgLexerTRUE          = 16
	PcgLexerFALSE         = 17
	PcgLexerNOT           = 18
	PcgLexerAND           = 19
	PcgLexerOR            = 20
	PcgLexerASSIGN        = 21
	PcgLexerLE            = 22
	PcgLexerGE            = 23
	PcgLexerEQ            = 24
	PcgLexerNE            = 25
	PcgLexerADD           = 26
	PcgLexerSUB           = 27
	PcgLexerMUL           = 28
	PcgLexerDIV           = 29
	PcgLexerLPAR          = 30
	PcgLexerRPAR          = 31
	PcgLexerCOLON         = 32
	PcgLexerLT            = 33
	PcgLexerGT            = 34
	PcgLexerNUMBER        = 35
	PcgLexerID            = 36
)

func (l *PcgLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 1:
		l.BLOCK_COMMENT_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PcgLexer) BLOCK_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		lexical_error("Unclosed block comment")

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
