package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/igso/languages/manifest"
)

type KeyValue struct {
	Key   string
	Value string
}

type JavaPackage struct {
	Name string
}

type IgsoManifest struct {
	ExportPackage []JavaPackage
	Package       []JavaPackage
	KeyValues     []KeyValue
}

type MfIdentListener struct {
	currentKey string
	manifest   IgsoManifest


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

func Analysis(code string, fileName string) IgsoManifest {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Mf()

	listener := NewMfIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetResult()
}

func NewMfIdentListener(fileName string) *MfIdentListener {
	listener := &MfIdentListener{}
	listener.manifest = IgsoManifest{}
	return listener
}

func (s *MfIdentListener) EnterSection(ctx *parser.SectionContext) {
	if ctx.Key() != nil {
		s.currentKey = ctx.Key().GetText()
	}
}

func (s *MfIdentListener) EnterValue(ctx *parser.ValueContext) {
	s.manifest.KeyValues = append(s.manifest.KeyValues, KeyValue{
		Key:   s.currentKey,
		Value: ctx.GetText(),
	})

	if ctx.VERSION() != nil {
		s.manifest.Package = append(s.manifest.Package, JavaPackage{
			Name: ctx.OTHERS().GetText(),
		})
	}
}

func (s *MfIdentListener) GetResult() IgsoManifest {
	return s.manifest
}
