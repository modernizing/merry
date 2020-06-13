package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/igso/languages/manifest"
	"github.com/phodal/igso/pkg/domain"
	"regexp"
	"strings"
)

type MfIdentListener struct {
	currentKey string
	manifest   domain.IgsoManifest

	parser.BaseManifestListener
}

func processStream(is antlr.CharStream) *parser.ManifestParser {
	lexer := parser.NewManifestLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewManifestParser(stream)
}

func ProcessString(code string) *parser.ManifestParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

func Analysis(code string, fileName string) domain.IgsoManifest {
	re := regexp.MustCompile(`\r?\n `)
	code = re.ReplaceAllString(code, "")

	scriptParser := ProcessString(code)
	context := scriptParser.Mf()

	listener := NewMfIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetResult()
}

func NewMfIdentListener() *MfIdentListener {
	listener := &MfIdentListener{}
	listener.manifest = domain.IgsoManifest{}
	return listener
}

func (s *MfIdentListener) EnterSection(ctx *parser.SectionContext) {
	if ctx.Key() != nil {
		s.currentKey = ctx.Key().GetText()
	}
}

func (s *MfIdentListener) EnterValue(ctx *parser.ValueContext) {
	s.manifest.KeyValues = append(s.manifest.KeyValues, domain.KeyValue{
		Key:   s.currentKey,
		Value: ctx.GetText(),
	})

	for _, pkg := range ctx.AllPkg() {
		pkgContext := (pkg).(*parser.PkgContext)
		javaPackage := domain.JavaPackage{
			Name: pkgContext.OTHERS().GetText(),
		}

		if len(pkgContext.AllConfigAssign()) > 0 {
			for _, config := range pkgContext.AllConfigAssign() {
				configAssign := (config).(*parser.ConfigAssignContext)
				assignText := configAssign.AssignKey().GetText()
				s.buildPackageVersion(assignText, configAssign, &javaPackage)

				text := ""
				if configAssign.AssignValue() != nil {
					text = configAssign.AssignValue().GetText()
				}

				javaPackage.Config = append(javaPackage.Config, domain.KeyValue{
					Key:   assignText,
					Value: text,
				})
			}
		}

		buildPackageName(ctx, s)

		buildPackageVersion(ctx, s)

		if s.currentKey == "Import-Package" {
			s.manifest.ImportPackage = append(s.manifest.ImportPackage, javaPackage)
		}
		if s.currentKey == "Export-Package" {
			s.manifest.ExportPackage = append(s.manifest.ExportPackage, javaPackage)
		}
	}
}

func (s *MfIdentListener) buildPackageVersion(assignText string, configAssign *parser.ConfigAssignContext, javaPackage *domain.JavaPackage) {
	if assignText == "version" {
		if configAssign.AssignValue() == nil {
			fmt.Println(configAssign.GetText())
			return
		}

		versionText := configAssign.AssignValue().GetText()
		if strings.Contains(versionText, "[") {
			versionInfo := versionText[2 : len(versionText)-2]
			split := strings.Split(versionInfo, ",")
			javaPackage.VersionInfo = versionInfo
			if len(split) == 2 {
				javaPackage.StartVersion = split[0]
				javaPackage.EndVersion = split[1]
			}
		} else {
			if s.currentKey == "Import-Package" {
				javaPackage.StartVersion = versionText[1 : len(versionText)-1]
			}
			if s.currentKey == "Export-Package" {
				javaPackage.ExportVersion = versionText[1 : len(versionText)-1]
			}
		}
	}
}

func buildPackageVersion(ctx *parser.ValueContext, s *MfIdentListener) {
	if s.currentKey == "Implementation-Version" {
		s.manifest.Version = ctx.GetText()
	}
	if s.manifest.Version == "" && s.currentKey == "Bundle-Version" {
		s.manifest.Version = ctx.GetText()
	}
}

func buildPackageName(ctx *parser.ValueContext, s *MfIdentListener) {
	if s.currentKey == "Private-Package" {
		s.manifest.PackageName = ctx.GetText()
	}
	if s.currentKey == "Bundle-SymbolicName" {
		s.manifest.PackageName = ctx.GetText()
	}
}

func (s *MfIdentListener) GetResult() domain.IgsoManifest {
	return s.manifest
}
