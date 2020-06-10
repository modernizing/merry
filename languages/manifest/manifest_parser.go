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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 36, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 7, 2, 11, 10, 2, 12, 2, 14,
	2, 14, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 29, 10, 4, 3, 4, 5, 4, 32, 10, 4, 5, 4, 34,
	10, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 37, 2, 12, 3, 2, 2, 2, 4, 17, 3,
	2, 2, 2, 6, 22, 3, 2, 2, 2, 8, 11, 7, 45, 2, 2, 9, 11, 5, 4, 3, 2, 10,
	8, 3, 2, 2, 2, 10, 9, 3, 2, 2, 2, 11, 14, 3, 2, 2, 2, 12, 10, 3, 2, 2,
	2, 12, 13, 3, 2, 2, 2, 13, 15, 3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 15, 16,
	7, 2, 2, 3, 16, 3, 3, 2, 2, 2, 17, 18, 7, 5, 2, 2, 18, 19, 7, 10, 2, 2,
	19, 20, 7, 46, 2, 2, 20, 21, 5, 6, 4, 2, 21, 5, 3, 2, 2, 2, 22, 33, 7,
	6, 2, 2, 23, 24, 7, 17, 2, 2, 24, 25, 7, 4, 2, 2, 25, 26, 7, 20, 2, 2,
	26, 28, 7, 48, 2, 2, 27, 29, 7, 17, 2, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3,
	2, 2, 2, 29, 31, 3, 2, 2, 2, 30, 32, 7, 3, 2, 2, 31, 30, 3, 2, 2, 2, 31,
	32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 23, 3, 2, 2, 2, 33, 34, 3, 2, 2,
	2, 34, 7, 3, 2, 2, 2, 7, 10, 12, 28, 31, 33,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'version'", "", "", "", "", "", "':'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'",
	"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'",
	"'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'\"'", "':='",
}
var symbolicNames = []string{
	"", "ConfigAssign", "VERSION", "Key", "OTHERS", "ValueText", "Version",
	"IDENTIFIER", "COLON", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
	"RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT", "BANG", "TILDE",
	"QUESTION", "EQUAL", "LE", "GE", "NOTEQUAL", "AND", "OR", "INC", "DEC",
	"ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET", "MOD", "DQUOTE",
	"SEQUAL", "Uppercase", "LINE_COMMENT", "SPACE", "NL", "STRING_LITERAL",
}

var ruleNames = []string{
	"mf", "section", "value",
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
	ManifestParserConfigAssign   = 1
	ManifestParserVERSION        = 2
	ManifestParserKey            = 3
	ManifestParserOTHERS         = 4
	ManifestParserValueText      = 5
	ManifestParserVersion        = 6
	ManifestParserIDENTIFIER     = 7
	ManifestParserCOLON          = 8
	ManifestParserLPAREN         = 9
	ManifestParserRPAREN         = 10
	ManifestParserLBRACE         = 11
	ManifestParserRBRACE         = 12
	ManifestParserLBRACK         = 13
	ManifestParserRBRACK         = 14
	ManifestParserSEMI           = 15
	ManifestParserCOMMA          = 16
	ManifestParserDOT            = 17
	ManifestParserASSIGN         = 18
	ManifestParserGT             = 19
	ManifestParserLT             = 20
	ManifestParserBANG           = 21
	ManifestParserTILDE          = 22
	ManifestParserQUESTION       = 23
	ManifestParserEQUAL          = 24
	ManifestParserLE             = 25
	ManifestParserGE             = 26
	ManifestParserNOTEQUAL       = 27
	ManifestParserAND            = 28
	ManifestParserOR             = 29
	ManifestParserINC            = 30
	ManifestParserDEC            = 31
	ManifestParserADD            = 32
	ManifestParserSUB            = 33
	ManifestParserMUL            = 34
	ManifestParserDIV            = 35
	ManifestParserBITAND         = 36
	ManifestParserBITOR          = 37
	ManifestParserCARET          = 38
	ManifestParserMOD            = 39
	ManifestParserDQUOTE         = 40
	ManifestParserSEQUAL         = 41
	ManifestParserUppercase      = 42
	ManifestParserLINE_COMMENT   = 43
	ManifestParserSPACE          = 44
	ManifestParserNL             = 45
	ManifestParserSTRING_LITERAL = 46
)

// ManifestParser rules.
const (
	ManifestParserRULE_mf      = 0
	ManifestParserRULE_section = 1
	ManifestParserRULE_value   = 2
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
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ManifestParserKey || _la == ManifestParserLINE_COMMENT {
		p.SetState(8)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ManifestParserLINE_COMMENT:
			{
				p.SetState(6)
				p.Match(ManifestParserLINE_COMMENT)
			}

		case ManifestParserKey:
			{
				p.SetState(7)
				p.Section()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
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
		p.SetState(15)
		p.Match(ManifestParserKey)
	}
	{
		p.SetState(16)
		p.Match(ManifestParserCOLON)
	}
	{
		p.SetState(17)
		p.Match(ManifestParserSPACE)
	}
	{
		p.SetState(18)
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

func (s *ValueContext) ConfigAssign() antlr.TerminalNode {
	return s.GetToken(ManifestParserConfigAssign, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(ManifestParserOTHERS)
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ManifestParserSEMI {
		{
			p.SetState(21)
			p.Match(ManifestParserSEMI)
		}
		{
			p.SetState(22)
			p.Match(ManifestParserVERSION)
		}
		{
			p.SetState(23)
			p.Match(ManifestParserASSIGN)
		}
		{
			p.SetState(24)
			p.Match(ManifestParserSTRING_LITERAL)
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ManifestParserSEMI {
			{
				p.SetState(25)
				p.Match(ManifestParserSEMI)
			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ManifestParserConfigAssign {
			{
				p.SetState(28)
				p.Match(ManifestParserConfigAssign)
			}

		}

	}

	return localctx
}
