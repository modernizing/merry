// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Manifest

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 44, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 7, 2, 13, 10,
	2, 12, 2, 14, 2, 16, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 23, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 5, 3, 33, 10, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 39, 10, 5, 12, 5, 14, 5, 42, 11, 5, 3, 5,
	2, 2, 6, 2, 4, 6, 8, 2, 3, 4, 2, 5, 5, 11, 11, 2, 45, 2, 14, 3, 2, 2, 2,
	4, 32, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 13, 7, 10,
	2, 2, 11, 13, 5, 4, 3, 2, 12, 10, 3, 2, 2, 2, 12, 11, 3, 2, 2, 2, 13, 16,
	3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 17, 3, 2, 2, 2,
	16, 14, 3, 2, 2, 2, 17, 18, 7, 2, 2, 3, 18, 3, 3, 2, 2, 2, 19, 20, 5, 6,
	4, 2, 20, 22, 7, 9, 2, 2, 21, 23, 7, 11, 2, 2, 22, 21, 3, 2, 2, 2, 22,
	23, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 5, 8, 5, 2, 25, 33, 3, 2, 2,
	2, 26, 27, 7, 4, 2, 2, 27, 29, 7, 9, 2, 2, 28, 30, 7, 11, 2, 2, 29, 28,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 5, 8, 5, 2,
	32, 19, 3, 2, 2, 2, 32, 26, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 35, 7, 3,
	2, 2, 35, 7, 3, 2, 2, 2, 36, 40, 7, 5, 2, 2, 37, 39, 9, 2, 2, 2, 38, 37,
	3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 9, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 8, 12, 14, 22, 29, 32, 40,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'Import-Package'", "", "", "", "", "", "':'",
}
var symbolicNames = []string{
	"", "", "Key", "ValueText", "Uppercase", "Lowercase", "Symbol", "Colon",
	"LINE_COMMENT", "Space", "NewLine",
}

var ruleNames = []string{
	"mf", "section", "isImport", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ManifestParser struct {
	*antlr.BaseParser
}

func NewManifestParser(input antlr.TokenStream) *ManifestParser {
	this := new(ManifestParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Manifest.g4"

	return this
}

// ManifestParser tokens.
const (
	ManifestParserEOF          = antlr.TokenEOF
	ManifestParserT__0         = 1
	ManifestParserKey          = 2
	ManifestParserValueText    = 3
	ManifestParserUppercase    = 4
	ManifestParserLowercase    = 5
	ManifestParserSymbol       = 6
	ManifestParserColon        = 7
	ManifestParserLINE_COMMENT = 8
	ManifestParserSpace        = 9
	ManifestParserNewLine      = 10
)

// ManifestParser rules.
const (
	ManifestParserRULE_mf       = 0
	ManifestParserRULE_section  = 1
	ManifestParserRULE_isImport = 2
	ManifestParserRULE_value    = 3
)

// IMfContext is an interface to support dynamic dispatch.
type IMfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMfContext differentiates from other interfaces.
	IsMfContext()
}

type MfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMfContext() *MfContext {
	var p = new(MfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_mf
	return p
}

func (*MfContext) IsMfContext() {}

func NewMfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MfContext {
	var p = new(MfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_mf

	return p
}

func (s *MfContext) GetParser() antlr.Parser { return s.parser }

func (s *MfContext) EOF() antlr.TerminalNode {
	return s.GetToken(ManifestParserEOF, 0)
}

func (s *MfContext) AllLINE_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserLINE_COMMENT)
}

func (s *MfContext) LINE_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserLINE_COMMENT, i)
}

func (s *MfContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *MfContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *MfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterMf(s)
	}
}

func (s *MfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitMf(s)
	}
}

func (p *ManifestParser) Mf() (localctx IMfContext) {
	localctx = NewMfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ManifestParserRULE_mf)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ManifestParserT__0)|(1<<ManifestParserKey)|(1<<ManifestParserLINE_COMMENT))) != 0 {
		p.SetState(10)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ManifestParserLINE_COMMENT:
			{
				p.SetState(8)
				p.Match(ManifestParserLINE_COMMENT)
			}

		case ManifestParserT__0, ManifestParserKey:
			{
				p.SetState(9)
				p.Section()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(15)
		p.Match(ManifestParserEOF)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) IsImport() IIsImportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsImportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsImportContext)
}

func (s *SectionContext) Colon() antlr.TerminalNode {
	return s.GetToken(ManifestParserColon, 0)
}

func (s *SectionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SectionContext) Space() antlr.TerminalNode {
	return s.GetToken(ManifestParserSpace, 0)
}

func (s *SectionContext) Key() antlr.TerminalNode {
	return s.GetToken(ManifestParserKey, 0)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *ManifestParser) Section() (localctx ISectionContext) {
	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ManifestParserRULE_section)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ManifestParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.IsImport()
		}
		{
			p.SetState(18)
			p.Match(ManifestParserColon)
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ManifestParserSpace {
			{
				p.SetState(19)
				p.Match(ManifestParserSpace)
			}

		}
		{
			p.SetState(22)
			p.Value()
		}

	case ManifestParserKey:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(ManifestParserKey)
		}
		{
			p.SetState(25)
			p.Match(ManifestParserColon)
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ManifestParserSpace {
			{
				p.SetState(26)
				p.Match(ManifestParserSpace)
			}

		}
		{
			p.SetState(29)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIsImportContext is an interface to support dynamic dispatch.
type IIsImportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsImportContext differentiates from other interfaces.
	IsIsImportContext()
}

type IsImportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsImportContext() *IsImportContext {
	var p = new(IsImportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_isImport
	return p
}

func (*IsImportContext) IsIsImportContext() {}

func NewIsImportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsImportContext {
	var p = new(IsImportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_isImport

	return p
}

func (s *IsImportContext) GetParser() antlr.Parser { return s.parser }
func (s *IsImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsImportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterIsImport(s)
	}
}

func (s *IsImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitIsImport(s)
	}
}

func (p *ManifestParser) IsImport() (localctx IIsImportContext) {
	localctx = NewIsImportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ManifestParserRULE_isImport)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(ManifestParserT__0)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) AllValueText() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserValueText)
}

func (s *ValueContext) ValueText(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserValueText, i)
}

func (s *ValueContext) AllSpace() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserSpace)
}

func (s *ValueContext) Space(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserSpace, i)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ManifestParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ManifestParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(ManifestParserValueText)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserValueText || _la == ManifestParserSpace {
		{
			p.SetState(35)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ManifestParserValueText || _la == ManifestParserSpace) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
