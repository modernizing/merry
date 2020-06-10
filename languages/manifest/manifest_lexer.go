// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 33, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 21, 10, 5, 13, 5, 14, 5, 22, 3,
	6, 3, 6, 7, 6, 27, 10, 6, 12, 6, 14, 6, 30, 11, 6, 3, 6, 3, 6, 2, 2, 7,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 5, 4, 2, 47, 47, 99, 124, 10, 2, 34,
	34, 44, 44, 46, 46, 48, 60, 66, 92, 94, 94, 97, 97, 99, 124, 4, 2, 12,
	12, 15, 15, 2, 34, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15, 3, 2, 2,
	2, 7, 17, 3, 2, 2, 2, 9, 20, 3, 2, 2, 2, 11, 24, 3, 2, 2, 2, 13, 14, 7,
	60, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 4, 67, 92, 2, 16, 6, 3, 2, 2, 2, 17,
	18, 9, 2, 2, 2, 18, 8, 3, 2, 2, 2, 19, 21, 9, 3, 2, 2, 20, 19, 3, 2, 2,
	2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 10,
	3, 2, 2, 2, 24, 28, 7, 61, 2, 2, 25, 27, 10, 4, 2, 2, 26, 25, 3, 2, 2,
	2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31,
	3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 8, 6, 2, 2, 32, 12, 3, 2, 2, 2,
	5, 2, 22, 28, 3, 2, 3, 2,
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
	"", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "START_HEAD", "HEAD_TEXT", "TEXT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "START_HEAD", "HEAD_TEXT", "TEXT", "LINE_COMMENT",
}

type ManifestLexer struct {
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

func NewManifestLexer(input antlr.CharStream) *ManifestLexer {

	l := new(ManifestLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Manifest.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ManifestLexer tokens.
const (
	ManifestLexerT__0         = 1
	ManifestLexerSTART_HEAD   = 2
	ManifestLexerHEAD_TEXT    = 3
	ManifestLexerTEXT         = 4
	ManifestLexerLINE_COMMENT = 5
)
