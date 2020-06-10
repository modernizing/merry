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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 76, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 7, 2, 28, 10, 2, 12, 2, 14, 2, 31, 11, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 7, 10, 61, 10, 10, 12, 10, 14, 10, 64, 11,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 5, 12, 71, 10, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 2, 9, 2, 11, 5, 13, 6, 15, 7, 17,
	8, 19, 9, 21, 10, 23, 11, 3, 2, 9, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59,
	3, 2, 67, 92, 3, 2, 99, 124, 9, 2, 36, 36, 42, 43, 46, 48, 61, 61, 63,
	63, 93, 93, 95, 95, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 79,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 42, 3, 2, 2, 2, 7, 44, 3,
	2, 2, 2, 9, 48, 3, 2, 2, 2, 11, 50, 3, 2, 2, 2, 13, 52, 3, 2, 2, 2, 15,
	54, 3, 2, 2, 2, 17, 56, 3, 2, 2, 2, 19, 58, 3, 2, 2, 2, 21, 67, 3, 2, 2,
	2, 23, 70, 3, 2, 2, 2, 25, 29, 5, 11, 6, 2, 26, 28, 5, 7, 4, 2, 27, 26,
	3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2,
	30, 32, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 7, 47, 2, 2, 33, 37, 5,
	11, 6, 2, 34, 36, 5, 7, 4, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37,
	35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 4, 3, 2, 2, 2, 39, 37, 3, 2, 2,
	2, 40, 43, 5, 9, 5, 2, 41, 43, 5, 15, 8, 2, 42, 40, 3, 2, 2, 2, 42, 41,
	3, 2, 2, 2, 43, 6, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2, 45, 8, 3, 2, 2, 2, 46,
	49, 5, 7, 4, 2, 47, 49, 9, 3, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2,
	2, 49, 10, 3, 2, 2, 2, 50, 51, 9, 4, 2, 2, 51, 12, 3, 2, 2, 2, 52, 53,
	9, 5, 2, 2, 53, 14, 3, 2, 2, 2, 54, 55, 9, 6, 2, 2, 55, 16, 3, 2, 2, 2,
	56, 57, 7, 60, 2, 2, 57, 18, 3, 2, 2, 2, 58, 62, 7, 61, 2, 2, 59, 61, 10,
	7, 2, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62,
	63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 8, 10,
	2, 2, 66, 20, 3, 2, 2, 2, 67, 68, 9, 8, 2, 2, 68, 22, 3, 2, 2, 2, 69, 71,
	7, 15, 2, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 73, 7, 12, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 8, 12, 3, 2, 75, 24, 3,
	2, 2, 2, 9, 2, 29, 37, 42, 48, 62, 70, 4, 2, 3, 2, 8, 2, 2,
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
	"", "", "", "", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Key", "ValueText", "Uppercase", "Lowercase", "Symbol", "Colon", "LINE_COMMENT",
	"Space", "NewLine",
}

var lexerRuleNames = []string{
	"Key", "ValueText", "Letter", "LetterOrDigit", "Uppercase", "Lowercase",
	"Symbol", "Colon", "LINE_COMMENT", "Space", "NewLine",
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
	ManifestLexerKey          = 1
	ManifestLexerValueText    = 2
	ManifestLexerUppercase    = 3
	ManifestLexerLowercase    = 4
	ManifestLexerSymbol       = 5
	ManifestLexerColon        = 6
	ManifestLexerLINE_COMMENT = 7
	ManifestLexerSpace        = 8
	ManifestLexerNewLine      = 9
)
