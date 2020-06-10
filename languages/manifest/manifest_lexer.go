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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 60, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 6, 2, 17, 10, 2, 13, 2, 14, 2, 18, 3, 2, 3, 2, 6, 2, 23, 10, 2,
	13, 2, 14, 2, 24, 6, 2, 27, 10, 2, 13, 2, 14, 2, 28, 3, 3, 3, 3, 6, 3,
	33, 10, 3, 13, 3, 14, 3, 34, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6,
	43, 10, 6, 12, 6, 14, 6, 46, 11, 6, 3, 6, 3, 6, 3, 7, 5, 7, 51, 10, 7,
	3, 7, 3, 7, 6, 7, 55, 10, 7, 13, 7, 14, 7, 56, 3, 7, 3, 7, 2, 2, 8, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 8, 3, 2, 50, 59, 3, 2, 67, 92,
	5, 2, 47, 47, 67, 92, 99, 124, 4, 2, 47, 47, 99, 124, 4, 2, 12, 12, 15,
	15, 4, 2, 11, 11, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	3, 16, 3, 2, 2, 2, 5, 30, 3, 2, 2, 2, 7, 36, 3, 2, 2, 2, 9, 38, 3, 2, 2,
	2, 11, 40, 3, 2, 2, 2, 13, 54, 3, 2, 2, 2, 15, 17, 9, 2, 2, 2, 16, 15,
	3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2,
	19, 26, 3, 2, 2, 2, 20, 27, 7, 48, 2, 2, 21, 23, 9, 2, 2, 2, 22, 21, 3,
	2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25,
	27, 3, 2, 2, 2, 26, 20, 3, 2, 2, 2, 26, 22, 3, 2, 2, 2, 27, 28, 3, 2, 2,
	2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 4, 3, 2, 2, 2, 30, 32, 9,
	3, 2, 2, 31, 33, 9, 4, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 6, 3, 2, 2, 2, 36, 37, 9, 5, 2,
	2, 37, 8, 3, 2, 2, 2, 38, 39, 7, 60, 2, 2, 39, 10, 3, 2, 2, 2, 40, 44,
	7, 61, 2, 2, 41, 43, 10, 6, 2, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2,
	2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44,
	3, 2, 2, 2, 47, 48, 8, 6, 2, 2, 48, 12, 3, 2, 2, 2, 49, 51, 7, 15, 2, 2,
	50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 55, 7,
	12, 2, 2, 53, 55, 9, 7, 2, 2, 54, 50, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55,
	56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 3, 2, 2,
	2, 58, 59, 8, 7, 3, 2, 59, 14, 3, 2, 2, 2, 13, 2, 18, 24, 26, 28, 32, 34,
	44, 50, 54, 56, 4, 2, 3, 2, 8, 2, 2,
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
	"", "", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Version", "START_HEAD", "HEAD_TEXT", "Colon", "LINE_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"Version", "START_HEAD", "HEAD_TEXT", "Colon", "LINE_COMMENT", "WS",
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
	ManifestLexerVersion      = 1
	ManifestLexerSTART_HEAD   = 2
	ManifestLexerHEAD_TEXT    = 3
	ManifestLexerColon        = 4
	ManifestLexerLINE_COMMENT = 5
	ManifestLexerWS           = 6
)
