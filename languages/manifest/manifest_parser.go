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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 53, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 4, 3,
	4, 7, 4, 37, 10, 4, 12, 4, 14, 4, 40, 11, 4, 3, 4, 5, 4, 43, 10, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8,
	10, 12, 2, 2, 2, 50, 2, 17, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 27, 3, 2,
	2, 2, 8, 44, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 50, 3, 2, 2, 2, 14, 16,
	5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2,
	17, 18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 21, 7,
	2, 2, 3, 21, 3, 3, 2, 2, 2, 22, 23, 7, 4, 2, 2, 23, 24, 7, 7, 2, 2, 24,
	25, 7, 42, 2, 2, 25, 26, 5, 6, 4, 2, 26, 5, 3, 2, 2, 2, 27, 32, 7, 5, 2,
	2, 28, 29, 7, 14, 2, 2, 29, 30, 7, 3, 2, 2, 30, 31, 7, 17, 2, 2, 31, 33,
	7, 44, 2, 2, 32, 28, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 38, 3, 2, 2, 2,
	34, 35, 7, 14, 2, 2, 35, 37, 5, 8, 5, 2, 36, 34, 3, 2, 2, 2, 37, 40, 3,
	2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 41, 43, 7, 14, 2, 2, 42, 41, 3, 2, 2, 2, 42, 43, 3, 2,
	2, 2, 43, 7, 3, 2, 2, 2, 44, 45, 5, 10, 6, 2, 45, 46, 7, 40, 2, 2, 46,
	47, 5, 12, 7, 2, 47, 9, 3, 2, 2, 2, 48, 49, 7, 5, 2, 2, 49, 11, 3, 2, 2,
	2, 50, 51, 7, 5, 2, 2, 51, 13, 3, 2, 2, 2, 6, 17, 32, 38, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'version'", "", "", "", "':'", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'", "'=='",
	"'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'", "'-'", "'*'",
	"'/'", "'&'", "'|'", "'^'", "'%'", "'\"'", "':='",
}
var symbolicNames = []string{
	"", "VERSION", "Key", "OTHERS", "ValueText", "COLON", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN",
	"GT", "LT", "BANG", "TILDE", "QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL",
	"AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR",
	"CARET", "MOD", "DQUOTE", "SEQUAL", "Uppercase", "SPACE", "NL", "STRING_LITERAL",
}

var ruleNames = []string{
	"mf", "section", "value", "configAssign", "assignKey", "assignValue",
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
	ManifestParserVERSION        = 1
	ManifestParserKey            = 2
	ManifestParserOTHERS         = 3
	ManifestParserValueText      = 4
	ManifestParserCOLON          = 5
	ManifestParserLPAREN         = 6
	ManifestParserRPAREN         = 7
	ManifestParserLBRACE         = 8
	ManifestParserRBRACE         = 9
	ManifestParserLBRACK         = 10
	ManifestParserRBRACK         = 11
	ManifestParserSEMI           = 12
	ManifestParserCOMMA          = 13
	ManifestParserDOT            = 14
	ManifestParserASSIGN         = 15
	ManifestParserGT             = 16
	ManifestParserLT             = 17
	ManifestParserBANG           = 18
	ManifestParserTILDE          = 19
	ManifestParserQUESTION       = 20
	ManifestParserEQUAL          = 21
	ManifestParserLE             = 22
	ManifestParserGE             = 23
	ManifestParserNOTEQUAL       = 24
	ManifestParserAND            = 25
	ManifestParserOR             = 26
	ManifestParserINC            = 27
	ManifestParserDEC            = 28
	ManifestParserADD            = 29
	ManifestParserSUB            = 30
	ManifestParserMUL            = 31
	ManifestParserDIV            = 32
	ManifestParserBITAND         = 33
	ManifestParserBITOR          = 34
	ManifestParserCARET          = 35
	ManifestParserMOD            = 36
	ManifestParserDQUOTE         = 37
	ManifestParserSEQUAL         = 38
	ManifestParserUppercase      = 39
	ManifestParserSPACE          = 40
	ManifestParserNL             = 41
	ManifestParserSTRING_LITERAL = 42
)

// ManifestParser rules.
const (
	ManifestParserRULE_mf           = 0
	ManifestParserRULE_section      = 1
	ManifestParserRULE_value        = 2
	ManifestParserRULE_configAssign = 3
	ManifestParserRULE_assignKey    = 4
	ManifestParserRULE_assignValue  = 5
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserKey {
		{
			p.SetState(12)
			p.Section()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
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
		p.SetState(20)
		p.Match(ManifestParserKey)
	}
	{
		p.SetState(21)
		p.Match(ManifestParserCOLON)
	}
	{
		p.SetState(22)
		p.Match(ManifestParserSPACE)
	}
	{
		p.SetState(23)
		p.Value()
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

func (s *ValueContext) OTHERS() antlr.TerminalNode {
	return s.GetToken(ManifestParserOTHERS, 0)
}

func (s *ValueContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ManifestParserSEMI)
}

func (s *ValueContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ManifestParserSEMI, i)
}

func (s *ValueContext) VERSION() antlr.TerminalNode {
	return s.GetToken(ManifestParserVERSION, 0)
}

func (s *ValueContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ManifestParserASSIGN, 0)
}

func (s *ValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ManifestParserSTRING_LITERAL, 0)
}

func (s *ValueContext) AllConfigAssign() []IConfigAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConfigAssignContext)(nil)).Elem())
	var tst = make([]IConfigAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConfigAssignContext)
		}
	}

	return tst
}

func (s *ValueContext) ConfigAssign(i int) IConfigAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConfigAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConfigAssignContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(ManifestParserOTHERS)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(26)
			p.Match(ManifestParserSEMI)
		}
		{
			p.SetState(27)
			p.Match(ManifestParserVERSION)
		}
		{
			p.SetState(28)
			p.Match(ManifestParserASSIGN)
		}
		{
			p.SetState(29)
			p.Match(ManifestParserSTRING_LITERAL)
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(32)
				p.Match(ManifestParserSEMI)
			}
			{
				p.SetState(33)
				p.ConfigAssign()
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ManifestParserSEMI {
		{
			p.SetState(39)
			p.Match(ManifestParserSEMI)
		}

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

func (s *ConfigAssignContext) SEQUAL() antlr.TerminalNode {
	return s.GetToken(ManifestParserSEQUAL, 0)
}

func (s *ConfigAssignContext) AssignValue() IAssignValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignValueContext)
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
	p.EnterRule(localctx, 6, ManifestParserRULE_configAssign)

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
		p.SetState(42)
		p.AssignKey()
	}
	{
		p.SetState(43)
		p.Match(ManifestParserSEQUAL)
	}
	{
		p.SetState(44)
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
	p.EnterRule(localctx, 8, ManifestParserRULE_assignKey)

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
		p.SetState(46)
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
	p.EnterRule(localctx, 10, ManifestParserRULE_assignValue)

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
		p.Match(ManifestParserOTHERS)
	}

	return localctx
}
