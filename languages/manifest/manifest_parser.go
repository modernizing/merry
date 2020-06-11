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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 62, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 33, 10, 3, 3, 4, 3,
	4, 3, 4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3, 4, 5, 4, 44, 10,
	4, 3, 5, 3, 5, 3, 5, 7, 5, 49, 10, 5, 12, 5, 14, 5, 52, 11, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12,
	14, 2, 4, 4, 2, 18, 18, 41, 41, 4, 2, 6, 6, 45, 45, 2, 59, 2, 19, 3, 2,
	2, 2, 4, 32, 3, 2, 2, 2, 6, 43, 3, 2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 53,
	3, 2, 2, 2, 12, 57, 3, 2, 2, 2, 14, 59, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2,
	17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3,
	2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 2, 2, 3, 23,
	3, 3, 2, 2, 2, 24, 25, 7, 3, 2, 2, 25, 26, 7, 8, 2, 2, 26, 27, 7, 43, 2,
	2, 27, 33, 7, 4, 2, 2, 28, 29, 7, 5, 2, 2, 29, 30, 7, 8, 2, 2, 30, 31,
	7, 43, 2, 2, 31, 33, 5, 6, 4, 2, 32, 24, 3, 2, 2, 2, 32, 28, 3, 2, 2, 2,
	33, 5, 3, 2, 2, 2, 34, 39, 5, 8, 5, 2, 35, 36, 7, 16, 2, 2, 36, 38, 5,
	8, 5, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39,
	40, 3, 2, 2, 2, 40, 44, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 44, 7, 45,
	2, 2, 43, 34, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 7, 3, 2, 2, 2, 45, 50,
	7, 6, 2, 2, 46, 47, 7, 15, 2, 2, 47, 49, 5, 10, 6, 2, 48, 46, 3, 2, 2,
	2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 9, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2, 54, 55, 9, 2, 2, 2, 55,
	56, 5, 14, 8, 2, 56, 11, 3, 2, 2, 2, 57, 58, 7, 6, 2, 2, 58, 13, 3, 2,
	2, 2, 59, 60, 9, 3, 2, 2, 60, 15, 3, 2, 2, 2, 7, 19, 32, 39, 43, 50,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'SHA1-Digest'", "", "", "", "", "':'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'",
	"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'",
	"'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'\"'", "':='",
}
var symbolicNames = []string{
	"", "", "Hash", "Key", "OTHERS", "ValueText", "COLON", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN",
	"GT", "LT", "BANG", "TILDE", "QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL",
	"AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR",
	"CARET", "MOD", "DQUOTE", "SEQUAL", "Uppercase", "SPACE", "NL", "STRING_LITERAL",
	"IDENTIFIER",
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
	ManifestParserT__0           = 1
	ManifestParserHash           = 2
	ManifestParserKey            = 3
	ManifestParserOTHERS         = 4
	ManifestParserValueText      = 5
	ManifestParserCOLON          = 6
	ManifestParserLPAREN         = 7
	ManifestParserRPAREN         = 8
	ManifestParserLBRACE         = 9
	ManifestParserRBRACE         = 10
	ManifestParserLBRACK         = 11
	ManifestParserRBRACK         = 12
	ManifestParserSEMI           = 13
	ManifestParserCOMMA          = 14
	ManifestParserDOT            = 15
	ManifestParserASSIGN         = 16
	ManifestParserGT             = 17
	ManifestParserLT             = 18
	ManifestParserBANG           = 19
	ManifestParserTILDE          = 20
	ManifestParserQUESTION       = 21
	ManifestParserEQUAL          = 22
	ManifestParserLE             = 23
	ManifestParserGE             = 24
	ManifestParserNOTEQUAL       = 25
	ManifestParserAND            = 26
	ManifestParserOR             = 27
	ManifestParserINC            = 28
	ManifestParserDEC            = 29
	ManifestParserADD            = 30
	ManifestParserSUB            = 31
	ManifestParserMUL            = 32
	ManifestParserDIV            = 33
	ManifestParserBITAND         = 34
	ManifestParserBITOR          = 35
	ManifestParserCARET          = 36
	ManifestParserMOD            = 37
	ManifestParserDQUOTE         = 38
	ManifestParserSEQUAL         = 39
	ManifestParserUppercase      = 40
	ManifestParserSPACE          = 41
	ManifestParserNL             = 42
	ManifestParserSTRING_LITERAL = 43
	ManifestParserIDENTIFIER     = 44
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

	for _la == ManifestParserT__0 || _la == ManifestParserKey {
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

func (s *SectionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ManifestParserCOLON, 0)
}

func (s *SectionContext) SPACE() antlr.TerminalNode {
	return s.GetToken(ManifestParserSPACE, 0)
}

func (s *SectionContext) Hash() antlr.TerminalNode {
	return s.GetToken(ManifestParserHash, 0)
}

func (s *SectionContext) Key() antlr.TerminalNode {
	return s.GetToken(ManifestParserKey, 0)
}

func (s *SectionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ManifestParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(ManifestParserT__0)
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
			p.Match(ManifestParserHash)
		}

	case ManifestParserKey:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(ManifestParserKey)
		}
		{
			p.SetState(27)
			p.Match(ManifestParserCOLON)
		}
		{
			p.SetState(28)
			p.Match(ManifestParserSPACE)
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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ManifestParserOTHERS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Pkg()
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ManifestParserCOMMA {
			{
				p.SetState(33)
				p.Match(ManifestParserCOMMA)
			}
			{
				p.SetState(34)
				p.Pkg()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ManifestParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
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
		p.SetState(43)
		p.Match(ManifestParserOTHERS)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserSEMI {
		{
			p.SetState(44)
			p.Match(ManifestParserSEMI)
		}
		{
			p.SetState(45)
			p.ConfigAssign()
		}

		p.SetState(50)
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
		p.SetState(51)
		p.AssignKey()
	}
	{
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ManifestParserASSIGN || _la == ManifestParserSEQUAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(53)
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
		p.SetState(55)
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
		p.SetState(57)
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
