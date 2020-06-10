package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/igso/languages/manifest"
)

type Manifest struct {
	Key   string
	Value string
}

type MfIdentListener struct {
	currentKey string
	manifest   []Manifest

	parser.BaseManifestListener
}

func processStream(is antlr.CharStream) *parser.ManifestParser {
	lexer := parser.NewManifestLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewManifestParser(stream)
}

func ProcessTsString(code string) *parser.ManifestParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

func Analysis(code string, fileName string) []Manifest {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Mf()

	listener := NewMfIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetResult()
}

func NewMfIdentListener(fileName string) *MfIdentListener {
	listener := &MfIdentListener{}
	return listener
}

func (s *MfIdentListener) EnterSection(ctx *parser.SectionContext) {
	s.currentKey = ctx.Key().GetText()
}

func (s *MfIdentListener) EnterValue(ctx *parser.ValueContext) {
	s.manifest = append(s.manifest, Manifest{
		Key:   s.currentKey,
		Value: ctx.GetText(),
	})
}

func (s *MfIdentListener) GetResult() []Manifest {
	return s.manifest
}
