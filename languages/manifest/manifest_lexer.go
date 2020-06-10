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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 93, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 45, 10, 3, 12, 3, 14, 3,
	48, 11, 3, 3, 3, 3, 3, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14, 3, 56, 11, 3,
	3, 4, 3, 4, 5, 4, 60, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 66, 10, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 78,
	10, 11, 12, 11, 14, 11, 81, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	5, 13, 88, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7,
	5, 9, 2, 11, 2, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 3,
	2, 9, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 3, 2, 67, 92, 3, 2, 99, 124,
	9, 2, 36, 36, 42, 43, 46, 48, 61, 61, 63, 63, 93, 93, 95, 95, 4, 2, 12,
	12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 96, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 42, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9,
	61, 3, 2, 2, 2, 11, 65, 3, 2, 2, 2, 13, 67, 3, 2, 2, 2, 15, 69, 3, 2, 2,
	2, 17, 71, 3, 2, 2, 2, 19, 73, 3, 2, 2, 2, 21, 75, 3, 2, 2, 2, 23, 84,
	3, 2, 2, 2, 25, 87, 3, 2, 2, 2, 27, 28, 7, 75, 2, 2, 28, 29, 7, 111, 2,
	2, 29, 30, 7, 114, 2, 2, 30, 31, 7, 113, 2, 2, 31, 32, 7, 116, 2, 2, 32,
	33, 7, 118, 2, 2, 33, 34, 7, 47, 2, 2, 34, 35, 7, 82, 2, 2, 35, 36, 7,
	99, 2, 2, 36, 37, 7, 101, 2, 2, 37, 38, 7, 109, 2, 2, 38, 39, 7, 99, 2,
	2, 39, 40, 7, 105, 2, 2, 40, 41, 7, 103, 2, 2, 41, 4, 3, 2, 2, 2, 42, 46,
	5, 13, 7, 2, 43, 45, 5, 9, 5, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2,
	46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3,
	2, 2, 2, 49, 50, 7, 47, 2, 2, 50, 54, 5, 13, 7, 2, 51, 53, 5, 9, 5, 2,
	52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3,
	2, 2, 2, 55, 6, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 60, 5, 11, 6, 2, 58,
	60, 5, 17, 9, 2, 59, 57, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 8, 3, 2, 2,
	2, 61, 62, 9, 2, 2, 2, 62, 10, 3, 2, 2, 2, 63, 66, 5, 9, 5, 2, 64, 66,
	9, 3, 2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 12, 3, 2, 2, 2,
	67, 68, 9, 4, 2, 2, 68, 14, 3, 2, 2, 2, 69, 70, 9, 5, 2, 2, 70, 16, 3,
	2, 2, 2, 71, 72, 9, 6, 2, 2, 72, 18, 3, 2, 2, 2, 73, 74, 7, 60, 2, 2, 74,
	20, 3, 2, 2, 2, 75, 79, 7, 61, 2, 2, 76, 78, 10, 7, 2, 2, 77, 76, 3, 2,
	2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 8, 11, 2, 2, 83, 22, 3, 2, 2, 2,
	84, 85, 9, 8, 2, 2, 85, 24, 3, 2, 2, 2, 86, 88, 7, 15, 2, 2, 87, 86, 3,
	2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 12, 2, 2, 90,
	91, 3, 2, 2, 2, 91, 92, 8, 13, 3, 2, 92, 26, 3, 2, 2, 2, 9, 2, 46, 54,
	59, 65, 79, 87, 4, 2, 3, 2, 8, 2, 2,
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
	"", "'Import-Package'", "", "", "", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "Key", "ValueText", "Uppercase", "Lowercase", "Symbol", "Colon",
	"LINE_COMMENT", "Space", "NewLine",
}

var lexerRuleNames = []string{
	"T__0", "Key", "ValueText", "Letter", "LetterOrDigit", "Uppercase", "Lowercase",
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
	ManifestLexerT__0         = 1
	ManifestLexerKey          = 2
	ManifestLexerValueText    = 3
	ManifestLexerUppercase    = 4
	ManifestLexerLowercase    = 5
	ManifestLexerSymbol       = 6
	ManifestLexerColon        = 7
	ManifestLexerLINE_COMMENT = 8
	ManifestLexerSpace        = 9
	ManifestLexerNewLine      = 10
)
