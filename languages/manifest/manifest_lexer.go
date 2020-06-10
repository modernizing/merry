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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 137,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2,
	36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 2, 3, 2, 3, 2, 7, 2, 44, 10, 2,
	12, 2, 14, 2, 47, 11, 2, 5, 2, 49, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 6, 5, 83, 10, 5, 13, 5, 14, 5, 84, 6, 5, 87, 10, 5, 13,
	5, 14, 5, 88, 3, 5, 5, 5, 92, 10, 5, 3, 6, 3, 6, 3, 6, 6, 6, 97, 10, 6,
	13, 6, 14, 6, 98, 3, 7, 3, 7, 7, 7, 103, 10, 7, 12, 7, 14, 7, 106, 11,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 112, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 7, 13, 122, 10, 13, 12, 13, 14, 13, 125, 11,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 5, 15, 132, 10, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 2, 2, 16, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2, 17,
	2, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 14, 3, 2, 8, 4, 2, 67, 92,
	99, 124, 3, 2, 50, 59, 3, 2, 67, 92, 9, 2, 36, 36, 42, 43, 46, 47, 61,
	61, 63, 63, 93, 93, 95, 95, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34,
	2, 147, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 3, 48, 3, 2, 2, 2, 5, 50, 3, 2, 2, 2, 7, 65, 3, 2, 2, 2,
	9, 91, 3, 2, 2, 2, 11, 93, 3, 2, 2, 2, 13, 100, 3, 2, 2, 2, 15, 107, 3,
	2, 2, 2, 17, 111, 3, 2, 2, 2, 19, 113, 3, 2, 2, 2, 21, 115, 3, 2, 2, 2,
	23, 117, 3, 2, 2, 2, 25, 119, 3, 2, 2, 2, 27, 128, 3, 2, 2, 2, 29, 131,
	3, 2, 2, 2, 31, 49, 5, 5, 3, 2, 32, 49, 5, 7, 4, 2, 33, 37, 5, 21, 11,
	2, 34, 36, 5, 15, 8, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35,
	3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2,
	40, 41, 7, 47, 2, 2, 41, 45, 5, 21, 11, 2, 42, 44, 5, 15, 8, 2, 43, 42,
	3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2,
	46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 31, 3, 2, 2, 2, 48, 32, 3,
	2, 2, 2, 48, 33, 3, 2, 2, 2, 49, 4, 3, 2, 2, 2, 50, 51, 7, 75, 2, 2, 51,
	52, 7, 111, 2, 2, 52, 53, 7, 114, 2, 2, 53, 54, 7, 113, 2, 2, 54, 55, 7,
	116, 2, 2, 55, 56, 7, 118, 2, 2, 56, 57, 7, 47, 2, 2, 57, 58, 7, 82, 2,
	2, 58, 59, 7, 99, 2, 2, 59, 60, 7, 101, 2, 2, 60, 61, 7, 109, 2, 2, 61,
	62, 7, 99, 2, 2, 62, 63, 7, 105, 2, 2, 63, 64, 7, 103, 2, 2, 64, 6, 3,
	2, 2, 2, 65, 66, 7, 71, 2, 2, 66, 67, 7, 122, 2, 2, 67, 68, 7, 114, 2,
	2, 68, 69, 7, 113, 2, 2, 69, 70, 7, 116, 2, 2, 70, 71, 7, 118, 2, 2, 71,
	72, 7, 47, 2, 2, 72, 73, 7, 82, 2, 2, 73, 74, 7, 99, 2, 2, 74, 75, 7, 101,
	2, 2, 75, 76, 7, 109, 2, 2, 76, 77, 7, 99, 2, 2, 77, 78, 7, 105, 2, 2,
	78, 79, 7, 103, 2, 2, 79, 8, 3, 2, 2, 2, 80, 87, 7, 48, 2, 2, 81, 83, 5,
	17, 9, 2, 82, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 82, 3, 2, 2,
	2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 92,
	3, 2, 2, 2, 90, 92, 5, 23, 12, 2, 91, 86, 3, 2, 2, 2, 91, 90, 3, 2, 2,
	2, 92, 10, 3, 2, 2, 2, 93, 96, 5, 13, 7, 2, 94, 95, 7, 48, 2, 2, 95, 97,
	5, 13, 7, 2, 96, 94, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2,
	98, 99, 3, 2, 2, 2, 99, 12, 3, 2, 2, 2, 100, 104, 5, 15, 8, 2, 101, 103,
	5, 17, 9, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 14, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	107, 108, 9, 2, 2, 2, 108, 16, 3, 2, 2, 2, 109, 112, 5, 15, 8, 2, 110,
	112, 9, 3, 2, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 18, 3,
	2, 2, 2, 113, 114, 7, 60, 2, 2, 114, 20, 3, 2, 2, 2, 115, 116, 9, 4, 2,
	2, 116, 22, 3, 2, 2, 2, 117, 118, 9, 5, 2, 2, 118, 24, 3, 2, 2, 2, 119,
	123, 7, 61, 2, 2, 120, 122, 10, 6, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125,
	3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2,
	2, 2, 125, 123, 3, 2, 2, 2, 126, 127, 8, 13, 2, 2, 127, 26, 3, 2, 2, 2,
	128, 129, 9, 7, 2, 2, 129, 28, 3, 2, 2, 2, 130, 132, 7, 15, 2, 2, 131,
	130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134,
	7, 12, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 8, 15, 3, 2, 136, 30, 3, 2,
	2, 2, 15, 2, 37, 45, 48, 84, 86, 88, 91, 98, 104, 111, 123, 131, 4, 2,
	3, 2, 8, 2, 2,
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
	"", "", "'Import-Package'", "'Export-Package'", "", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Key", "IsImport", "IsExport", "ValueText", "QualifiedName", "IDENTIFIER",
	"Colon", "Uppercase", "Symbol", "LINE_COMMENT", "SPACE", "NL",
}

var lexerRuleNames = []string{
	"Key", "IsImport", "IsExport", "ValueText", "QualifiedName", "IDENTIFIER",
	"Letter", "LetterOrDigit", "Colon", "Uppercase", "Symbol", "LINE_COMMENT",
	"SPACE", "NL",
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
	ManifestLexerIsImport      = 2
	ManifestLexerIsExport      = 3
	ManifestLexerValueText     = 4
	ManifestLexerQualifiedName = 5
	ManifestLexerIDENTIFIER    = 6
	ManifestLexerColon         = 7
	ManifestLexerUppercase     = 8
	ManifestLexerSymbol        = 9
	ManifestLexerLINE_COMMENT  = 10
	ManifestLexerSPACE         = 11
	ManifestLexerNL            = 12
)
