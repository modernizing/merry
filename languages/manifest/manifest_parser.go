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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 59, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 35,
	10, 4, 12, 4, 14, 4, 38, 11, 4, 3, 4, 5, 4, 41, 10, 4, 3, 5, 3, 5, 3, 5,
	7, 5, 46, 10, 5, 12, 5, 14, 5, 49, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 4, 2, 16,
	16, 39, 39, 4, 2, 4, 4, 43, 43, 2, 56, 2, 19, 3, 2, 2, 2, 4, 24, 3, 2,
	2, 2, 6, 40, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 54,
	3, 2, 2, 2, 14, 56, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2,
	18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22, 3,
	2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 2, 2, 3, 23, 3, 3, 2, 2, 2, 24,
	25, 7, 3, 2, 2, 25, 26, 7, 6, 2, 2, 26, 27, 7, 41, 2, 2, 27, 29, 5, 6,
	4, 2, 28, 30, 7, 16, 2, 2, 29, 28, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	5, 3, 2, 2, 2, 31, 36, 5, 8, 5, 2, 32, 33, 7, 14, 2, 2, 33, 35, 5, 8, 5,
	2, 34, 32, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37,
	3, 2, 2, 2, 37, 41, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 41, 7, 43, 2, 2,
	40, 31, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 47, 7, 4,
	2, 2, 43, 44, 7, 13, 2, 2, 44, 46, 5, 10, 6, 2, 45, 43, 3, 2, 2, 2, 46,
	49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 9, 3, 2, 2,
	2, 49, 47, 3, 2, 2, 2, 50, 51, 5, 12, 7, 2, 51, 52, 9, 2, 2, 2, 52, 53,
	5, 14, 8, 2, 53, 11, 3, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 13, 3, 2, 2, 2,
	56, 57, 9, 3, 2, 2, 57, 15, 3, 2, 2, 2, 7, 19, 29, 36, 40, 47,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "':'", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'",
	"','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'", "'=='", "'<='",
	"'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'",
	"'&'", "'|'", "'^'", "'%'", "'\"'", "':='",
}
var symbolicNames = []string{
	"", "Key", "OTHERS", "ValueText", "COLON", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT",
	"BANG", "TILDE", "QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR",
	"INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET", "MOD",
	"DQUOTE", "SEQUAL", "Uppercase", "SPACE", "NL", "STRING_LITERAL", "IDENTIFIER",
}

var ruleNames = []string{
	"mf", "section", "value", "pkg", "configAssign", "assignKey", "assignValue",
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
	ManifestParserEOF            = antlr.TokenEOF
	ManifestParserKey            = 1
	ManifestParserOTHERS         = 2
	ManifestParserValueText      = 3
	ManifestParserCOLON          = 4
	ManifestParserLPAREN         = 5
	ManifestParserRPAREN         = 6
	ManifestParserLBRACE         = 7
	ManifestParserRBRACE         = 8
	ManifestParserLBRACK         = 9
	ManifestParserRBRACK         = 10
	ManifestParserSEMI           = 11
	ManifestParserCOMMA          = 12
	ManifestParserDOT            = 13
	ManifestParserASSIGN         = 14
	ManifestParserGT             = 15
	ManifestParserLT             = 16
	ManifestParserBANG           = 17
	ManifestParserTILDE          = 18
	ManifestParserQUESTION       = 19
	ManifestParserEQUAL          = 20
	ManifestParserLE             = 21
	ManifestParserGE             = 22
	ManifestParserNOTEQUAL       = 23
	ManifestParserAND            = 24
	ManifestParserOR             = 25
	ManifestParserINC            = 26
	ManifestParserDEC            = 27
	ManifestParserADD            = 28
	ManifestParserSUB            = 29
	ManifestParserMUL            = 30
	ManifestParserDIV            = 31
	ManifestParserBITAND         = 32
	ManifestParserBITOR          = 33
	ManifestParserCARET          = 34
	ManifestParserMOD            = 35
	ManifestParserDQUOTE         = 36
	ManifestParserSEQUAL         = 37
	ManifestParserUppercase      = 38
	ManifestParserSPACE          = 39
	ManifestParserNL             = 40
	ManifestParserSTRING_LITERAL = 41
	ManifestParserIDENTIFIER     = 42
)

// ManifestParser rules.
const (
	ManifestParserRULE_mf           = 0
	ManifestParserRULE_section      = 1
	ManifestParserRULE_value        = 2
	ManifestParserRULE_pkg          = 3
	ManifestParserRULE_configAssign = 4
	ManifestParserRULE_assignKey    = 5
	ManifestParserRULE_assignValue  = 6
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserKey {
		{
			p.SetState(14)
			p.Section()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
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

func (s *SectionContext) Key() antlr.TerminalNode {
	return s.GetToken(ManifestParserKey, 0)
}

func (s *SectionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ManifestParserCOLON, 0)
}

func (s *SectionContext) SPACE() antlr.TerminalNode {
	return s.GetToken(ManifestParserSPACE, 0)
}

func (s *SectionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SectionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ManifestParserASSIGN, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(ManifestParserKey)
	}
	{
		p.SetState(23)
		p.Match(ManifestParserCOLON)
	}
	{
		p.SetState(24)
		p.Match(ManifestParserSPACE)
	}
	{
		p.SetState(25)
		p.Value()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ManifestParserASSIGN {
		{
			p.SetState(26)
			p.Match(ManifestParserASSIGN)
		}

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

func (s *ValueContext) AllPkg() []IPkgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPkgContext)(nil)).Elem())
	var tst = make([]IPkgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPkgContext)
		}
	}

	return tst
}

func (s *ValueContext) Pkg(i int) IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

func (s *ValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserCOMMA)
}

func (s *ValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserCOMMA, i)
}

func (s *ValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ManifestParserSTRING_LITERAL, 0)
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
	p.EnterRule(localctx, 4, ManifestParserRULE_value)
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

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ManifestParserOTHERS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Pkg()
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ManifestParserCOMMA {
			{
				p.SetState(30)
				p.Match(ManifestParserCOMMA)
			}
			{
				p.SetState(31)
				p.Pkg()
			}

			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ManifestParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(ManifestParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) OTHERS() antlr.TerminalNode {
	return s.GetToken(ManifestParserOTHERS, 0)
}

func (s *PkgContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserSEMI)
}

func (s *PkgContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserSEMI, i)
}

func (s *PkgContext) AllConfigAssign() []IConfigAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigAssignContext)(nil)).Elem())
	var tst = make([]IConfigAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigAssignContext)
		}
	}

	return tst
}

func (s *PkgContext) ConfigAssign(i int) IConfigAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigAssignContext)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *ManifestParser) Pkg() (localctx IPkgContext) {
	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ManifestParserRULE_pkg)
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
		p.SetState(40)
		p.Match(ManifestParserOTHERS)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserSEMI {
		{
			p.SetState(41)
			p.Match(ManifestParserSEMI)
		}
		{
			p.SetState(42)
			p.ConfigAssign()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConfigAssignContext is an interface to support dynamic dispatch.
type IConfigAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigAssignContext differentiates from other interfaces.
	IsConfigAssignContext()
}

type ConfigAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigAssignContext() *ConfigAssignContext {
	var p = new(ConfigAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_configAssign
	return p
}

func (*ConfigAssignContext) IsConfigAssignContext() {}

func NewConfigAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigAssignContext {
	var p = new(ConfigAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_configAssign

	return p
}

func (s *ConfigAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigAssignContext) AssignKey() IAssignKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignKeyContext)
}

func (s *ConfigAssignContext) AssignValue() IAssignValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignValueContext)
}

func (s *ConfigAssignContext) SEQUAL() antlr.TerminalNode {
	return s.GetToken(ManifestParserSEQUAL, 0)
}

func (s *ConfigAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ManifestParserASSIGN, 0)
}

func (s *ConfigAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterConfigAssign(s)
	}
}

func (s *ConfigAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitConfigAssign(s)
	}
}

func (p *ManifestParser) ConfigAssign() (localctx IConfigAssignContext) {
	localctx = NewConfigAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ManifestParserRULE_configAssign)
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
		p.SetState(48)
		p.AssignKey()
	}
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ManifestParserASSIGN || _la == ManifestParserSEQUAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(50)
		p.AssignValue()
	}

	return localctx
}

// IAssignKeyContext is an interface to support dynamic dispatch.
type IAssignKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignKeyContext differentiates from other interfaces.
	IsAssignKeyContext()
}

type AssignKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignKeyContext() *AssignKeyContext {
	var p = new(AssignKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_assignKey
	return p
}

func (*AssignKeyContext) IsAssignKeyContext() {}

func NewAssignKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignKeyContext {
	var p = new(AssignKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_assignKey

	return p
}

func (s *AssignKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignKeyContext) OTHERS() antlr.TerminalNode {
	return s.GetToken(ManifestParserOTHERS, 0)
}

func (s *AssignKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterAssignKey(s)
	}
}

func (s *AssignKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitAssignKey(s)
	}
}

func (p *ManifestParser) AssignKey() (localctx IAssignKeyContext) {
	localctx = NewAssignKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ManifestParserRULE_assignKey)

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
		p.SetState(52)
		p.Match(ManifestParserOTHERS)
	}

	return localctx
}

// IAssignValueContext is an interface to support dynamic dispatch.
type IAssignValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignValueContext differentiates from other interfaces.
	IsAssignValueContext()
}

type AssignValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignValueContext() *AssignValueContext {
	var p = new(AssignValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManifestParserRULE_assignValue
	return p
}

func (*AssignValueContext) IsAssignValueContext() {}

func NewAssignValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignValueContext {
	var p = new(AssignValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManifestParserRULE_assignValue

	return p
}

func (s *AssignValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ManifestParserSTRING_LITERAL, 0)
}

func (s *AssignValueContext) OTHERS() antlr.TerminalNode {
	return s.GetToken(ManifestParserOTHERS, 0)
}

func (s *AssignValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.EnterAssignValue(s)
	}
}

func (s *AssignValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManifestListener); ok {
		listenerT.ExitAssignValue(s)
	}
}

func (p *ManifestParser) AssignValue() (localctx IAssignValueContext) {
	localctx = NewAssignValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ManifestParserRULE_assignValue)
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
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ManifestParserOTHERS || _la == ManifestParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
