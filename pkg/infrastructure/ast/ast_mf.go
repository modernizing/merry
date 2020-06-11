package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/igso/languages/manifest"
	"github.com/phodal/igso/pkg/domain"
	"regexp"
	"strings"
)

type MfIdentListener struct {
	currentKey string
	manifest   dependency.IgsoManifest

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

func Analysis(code string, fileName string) dependency.IgsoManifest {
	re := regexp.MustCompile(`\r?\n `)
	code = re.ReplaceAllString(code, "")

	scriptParser := ProcessTsString(code)
	context := scriptParser.Mf()

	listener := NewMfIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetResult()
}

func NewMfIdentListener(fileName string) *MfIdentListener {
	listener := &MfIdentListener{}
	listener.manifest = dependency.IgsoManifest{}
	return listener
}

func (s *MfIdentListener) EnterSection(ctx *parser.SectionContext) {
	if ctx.Key() != nil {
		s.currentKey = ctx.Key().GetText()
	}
}

func (s *MfIdentListener) EnterValue(ctx *parser.ValueContext) {
	s.manifest.KeyValues = append(s.manifest.KeyValues, dependency.KeyValue{
		Key:   s.currentKey,
		Value: ctx.GetText(),
	})

	for _, pkg := range ctx.AllPkg() {
		pkgContext := (pkg).(*parser.PkgContext)
		javaPackage := dependency.JavaPackage{
			Name:         pkgContext.OTHERS().GetText(),
		}

		if len(pkgContext.AllConfigAssign()) > 0 {
			for _, config := range pkgContext.AllConfigAssign() {
				configAssign := (config).(*parser.ConfigAssignContext)
				assignText := configAssign.AssignKey().GetText()
				if assignText == "version" {
					versionText := configAssign.AssignValue().GetText()
					versionInfo := versionText[2 : len(versionText)-2]
					split := strings.Split(versionInfo, ",")
					javaPackage.VersionInfo = versionInfo
					javaPackage.StartVersion = split[0]
					javaPackage.EndVersion = split[1]
				}
				javaPackage.Config = append(javaPackage.Config, dependency.KeyValue{
					Key:   assignText,
					Value: configAssign.AssignValue().GetText(),
				})
			}
		}

		s.manifest.Package = append(s.manifest.Package, javaPackage)
	}
}

func (s *MfIdentListener) GetResult() dependency.IgsoManifest {
	return s.manifest
}
