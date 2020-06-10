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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 99, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 2,
	3, 2, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 3, 3, 3, 6, 3,
	45, 10, 3, 13, 3, 14, 3, 46, 6, 3, 49, 10, 3, 13, 3, 14, 3, 50, 3, 3, 5,
	3, 54, 10, 3, 3, 4, 3, 4, 3, 4, 6, 4, 59, 10, 4, 13, 4, 14, 4, 60, 3, 5,
	3, 5, 7, 5, 65, 10, 5, 12, 5, 14, 5, 68, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	5, 7, 74, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7,
	11, 84, 10, 11, 12, 11, 14, 11, 87, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 5, 13, 94, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 2, 13, 2, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12,
	3, 2, 8, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 3, 2, 67, 92, 9, 2, 36, 36,
	42, 43, 46, 47, 61, 61, 63, 63, 93, 93, 95, 95, 4, 2, 12, 12, 15, 15, 4,
	2, 11, 11, 34, 34, 2, 107, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3,
	27, 3, 2, 2, 2, 5, 53, 3, 2, 2, 2, 7, 55, 3, 2, 2, 2, 9, 62, 3, 2, 2, 2,
	11, 69, 3, 2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 75, 3, 2, 2, 2, 17, 77, 3,
	2, 2, 2, 19, 79, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 90, 3, 2, 2, 2, 25,
	93, 3, 2, 2, 2, 27, 31, 5, 17, 9, 2, 28, 30, 5, 11, 6, 2, 29, 28, 3, 2,
	2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34,
	3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 35, 7, 47, 2, 2, 35, 39, 5, 17, 9,
	2, 36, 38, 5, 11, 6, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 4, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	42, 49, 7, 48, 2, 2, 43, 45, 5, 13, 7, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3,
	2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48,
	42, 3, 2, 2, 2, 48, 44, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 54, 5, 19, 10, 2, 53, 48,
	3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 6, 3, 2, 2, 2, 55, 58, 5, 9, 5, 2,
	56, 57, 7, 48, 2, 2, 57, 59, 5, 9, 5, 2, 58, 56, 3, 2, 2, 2, 59, 60, 3,
	2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 8, 3, 2, 2, 2, 62,
	66, 5, 11, 6, 2, 63, 65, 5, 13, 7, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2,
	2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 10, 3, 2, 2, 2, 68, 66,
	3, 2, 2, 2, 69, 70, 9, 2, 2, 2, 70, 12, 3, 2, 2, 2, 71, 74, 5, 11, 6, 2,
	72, 74, 9, 3, 2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 14, 3,
	2, 2, 2, 75, 76, 7, 60, 2, 2, 76, 16, 3, 2, 2, 2, 77, 78, 9, 4, 2, 2, 78,
	18, 3, 2, 2, 2, 79, 80, 9, 5, 2, 2, 80, 20, 3, 2, 2, 2, 81, 85, 7, 61,
	2, 2, 82, 84, 10, 6, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85,
	83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2,
	2, 88, 89, 8, 11, 2, 2, 89, 22, 3, 2, 2, 2, 90, 91, 9, 7, 2, 2, 91, 24,
	3, 2, 2, 2, 92, 94, 7, 15, 2, 2, 93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2,
	94, 95, 3, 2, 2, 2, 95, 96, 7, 12, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 8,
	13, 3, 2, 98, 26, 3, 2, 2, 2, 14, 2, 31, 39, 46, 48, 50, 53, 60, 66, 73,
	85, 93, 4, 2, 3, 2, 8, 2, 2,
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
	"", "", "", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Key", "ValueText", "QualifiedName", "IDENTIFIER", "Colon", "Uppercase",
	"Symbol", "LINE_COMMENT", "SPACE", "NL",
}

var lexerRuleNames = []string{
	"Key", "ValueText", "QualifiedName", "IDENTIFIER", "Letter", "LetterOrDigit",
	"Colon", "Uppercase", "Symbol", "LINE_COMMENT", "SPACE", "NL",
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
	ManifestLexerKey           = 1
	ManifestLexerValueText     = 2
	ManifestLexerQualifiedName = 3
	ManifestLexerIDENTIFIER    = 4
	ManifestLexerColon         = 5
	ManifestLexerUppercase     = 6
	ManifestLexerSymbol        = 7
	ManifestLexerLINE_COMMENT  = 8
	ManifestLexerSPACE         = 9
	ManifestLexerNL            = 10
)
