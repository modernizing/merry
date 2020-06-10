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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 7, 2, 13, 10,
	2, 12, 2, 14, 2, 16, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 6, 4, 26, 10, 4, 13, 4, 14, 4, 27, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4,
	6, 8, 2, 2, 2, 30, 2, 14, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 23, 3, 2, 2,
	2, 8, 29, 3, 2, 2, 2, 10, 13, 7, 7, 2, 2, 11, 13, 5, 4, 3, 2, 12, 10, 3,
	2, 2, 2, 12, 11, 3, 2, 2, 2, 13, 16, 3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 14,
	15, 3, 2, 2, 2, 15, 17, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 17, 18, 7, 2, 2,
	3, 18, 3, 3, 2, 2, 2, 19, 20, 5, 6, 4, 2, 20, 21, 7, 3, 2, 2, 21, 22, 5,
	8, 5, 2, 22, 5, 3, 2, 2, 2, 23, 25, 7, 4, 2, 2, 24, 26, 7, 5, 2, 2, 25,
	24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2,
	2, 28, 7, 3, 2, 2, 2, 29, 30, 7, 6, 2, 2, 30, 9, 3, 2, 2, 2, 5, 12, 14,
	27,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'",
}
var symbolicNames = []string{
	"", "", "START_HEAD", "HEAD_TEXT", "TEXT", "LINE_COMMENT",
}

var ruleNames = []string{
	"mf", "section", "section_header", "key_values",
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
	ManifestParserSTART_HEAD   = 2
	ManifestParserHEAD_TEXT    = 3
	ManifestParserTEXT         = 4
	ManifestParserLINE_COMMENT = 5
)

// ManifestParser rules.
const (
	ManifestParserRULE_mf             = 0
	ManifestParserRULE_section        = 1
	ManifestParserRULE_section_header = 2
	ManifestParserRULE_key_values     = 3
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

	for _la == ManifestParserSTART_HEAD || _la == ManifestParserLINE_COMMENT {
		p.SetState(10)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ManifestParserLINE_COMMENT:
			{
				p.SetState(8)
				p.Match(ManifestParserLINE_COMMENT)
			}

		case ManifestParserSTART_HEAD:
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

func (s *SectionContext) Section_header() ISection_headerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_headerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_headerContext)
}

func (s *SectionContext) Key_values() IKey_valuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKey_valuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKey_valuesContext)
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
		p.SetState(17)
		p.Section_header()
	}
	{
		p.SetState(18)
		p.Match(ManifestParserT__0)
	}
	{
		p.SetState(19)
		p.Key_values()
	}

	return localctx
}

// ISection_headerContext is an interface to support dynamic dispatch.
type ISection_headerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_headerContext differentiates from other interfaces.
	IsSection_headerContext()
}

type Section_headerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_headerContext() *Section_headerContext {
	var p = new(Section_headerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_section_header
	return p
}

func (*Section_headerContext) IsSection_headerContext() {}

func NewSection_headerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_headerContext {
	var p = new(Section_headerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_section_header

	return p
}

func (s *Section_headerContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_headerContext) START_HEAD() antlr.TerminalNode {
	return s.GetToken(ManifestParserSTART_HEAD, 0)
}

func (s *Section_headerContext) AllHEAD_TEXT() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserHEAD_TEXT)
}

func (s *Section_headerContext) HEAD_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserHEAD_TEXT, i)
}

func (s *Section_headerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_headerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_headerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterSection_header(s)
	}
}

func (s *Section_headerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitSection_header(s)
	}
}

func (p *ManifestParser) Section_header() (localctx ISection_headerContext) {
	localctx = NewSection_headerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ManifestParserRULE_section_header)
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
		p.SetState(21)
		p.Match(ManifestParserSTART_HEAD)
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ManifestParserHEAD_TEXT {
		{
			p.SetState(22)
			p.Match(ManifestParserHEAD_TEXT)
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKey_valuesContext is an interface to support dynamic dispatch.
type IKey_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKey_valuesContext differentiates from other interfaces.
	IsKey_valuesContext()
}

type Key_valuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_valuesContext() *Key_valuesContext {
	var p = new(Key_valuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_key_values
	return p
}

func (*Key_valuesContext) IsKey_valuesContext() {}

func NewKey_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_valuesContext {
	var p = new(Key_valuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_key_values

	return p
}

func (s *Key_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_valuesContext) TEXT() antlr.TerminalNode {
	return s.GetToken(ManifestParserTEXT, 0)
}

func (s *Key_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterKey_values(s)
	}
}

func (s *Key_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitKey_values(s)
	}
}

func (p *ManifestParser) Key_values() (localctx IKey_valuesContext) {
	localctx = NewKey_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ManifestParserRULE_key_values)

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
		p.SetState(27)
		p.Match(ManifestParserTEXT)
	}

	return localctx
}
