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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 46, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6,
	2, 15, 10, 2, 13, 2, 14, 2, 16, 3, 3, 3, 3, 6, 3, 21, 10, 3, 13, 3, 14,
	3, 22, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 29, 10, 5, 12, 5, 14, 5, 32, 11, 5,
	3, 5, 3, 5, 3, 6, 5, 6, 37, 10, 6, 3, 6, 3, 6, 6, 6, 41, 10, 6, 13, 6,
	14, 6, 42, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 7,
	7, 2, 42, 44, 48, 48, 50, 59, 67, 92, 99, 124, 3, 2, 67, 92, 5, 2, 47,
	47, 67, 92, 99, 124, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 51,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 3, 14, 3, 2, 2, 2, 5, 18, 3, 2, 2, 2, 7, 24, 3, 2, 2,
	2, 9, 26, 3, 2, 2, 2, 11, 40, 3, 2, 2, 2, 13, 15, 9, 2, 2, 2, 14, 13, 3,
	2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17,
	4, 3, 2, 2, 2, 18, 20, 9, 3, 2, 2, 19, 21, 9, 4, 2, 2, 20, 19, 3, 2, 2,
	2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 6, 3,
	2, 2, 2, 24, 25, 7, 60, 2, 2, 25, 8, 3, 2, 2, 2, 26, 30, 7, 61, 2, 2, 27,
	29, 10, 5, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2,
	2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 34,
	8, 5, 2, 2, 34, 10, 3, 2, 2, 2, 35, 37, 7, 15, 2, 2, 36, 35, 3, 2, 2, 2,
	36, 37, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 41, 7, 12, 2, 2, 39, 41, 9,
	6, 2, 2, 40, 36, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42,
	40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 8, 6, 3,
	2, 45, 12, 3, 2, 2, 2, 11, 2, 14, 16, 20, 22, 30, 36, 40, 42, 4, 2, 3,
	2, 8, 2, 2,
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
	"", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Value", "Key", "Colon", "LINE_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"Value", "Key", "Colon", "LINE_COMMENT", "WS",
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
	ManifestLexerValue        = 1
	ManifestLexerKey          = 2
	ManifestLexerColon        = 3
	ManifestLexerLINE_COMMENT = 4
	ManifestLexerWS           = 5
)
