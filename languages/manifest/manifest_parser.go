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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 7, 2, 11, 10, 2, 12, 2, 14,
	2, 14, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 30, 10, 4, 3, 4, 2, 2, 5, 2, 4, 6,
	2, 2, 2, 32, 2, 12, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8,
	11, 7, 43, 2, 2, 9, 11, 5, 4, 3, 2, 10, 8, 3, 2, 2, 2, 10, 9, 3, 2, 2,
	2, 11, 14, 3, 2, 2, 2, 12, 10, 3, 2, 2, 2, 12, 13, 3, 2, 2, 2, 13, 15,
	3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2, 2,
	17, 18, 7, 4, 2, 2, 18, 20, 7, 9, 2, 2, 19, 21, 7, 44, 2, 2, 20, 19, 3,
	2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 23, 5, 6, 4, 2, 23,
	5, 3, 2, 2, 2, 24, 29, 7, 5, 2, 2, 25, 26, 7, 16, 2, 2, 26, 27, 7, 3, 2,
	2, 27, 28, 7, 19, 2, 2, 28, 30, 7, 46, 2, 2, 29, 25, 3, 2, 2, 2, 29, 30,
	3, 2, 2, 2, 30, 7, 3, 2, 2, 2, 6, 10, 12, 20, 29,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'version'", "", "", "", "", "", "':'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'",
	"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'",
	"'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'\"'",
}
var symbolicNames = []string{
	"", "VERSION", "Key", "OTHERS", "ValueText", "Version", "IDENTIFIER", "COLON",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA",
	"DOT", "ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "DQUOTE", "Uppercase", "LINE_COMMENT",
	"SPACE", "NL", "STRING_LITERAL",
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
	ManifestParserVERSION        = 1
	ManifestParserKey            = 2
	ManifestParserOTHERS         = 3
	ManifestParserValueText      = 4
	ManifestParserVersion        = 5
	ManifestParserIDENTIFIER     = 6
	ManifestParserCOLON          = 7
	ManifestParserLPAREN         = 8
	ManifestParserRPAREN         = 9
	ManifestParserLBRACE         = 10
	ManifestParserRBRACE         = 11
	ManifestParserLBRACK         = 12
	ManifestParserRBRACK         = 13
	ManifestParserSEMI           = 14
	ManifestParserCOMMA          = 15
	ManifestParserDOT            = 16
	ManifestParserASSIGN         = 17
	ManifestParserGT             = 18
	ManifestParserLT             = 19
	ManifestParserBANG           = 20
	ManifestParserTILDE          = 21
	ManifestParserQUESTION       = 22
	ManifestParserEQUAL          = 23
	ManifestParserLE             = 24
	ManifestParserGE             = 25
	ManifestParserNOTEQUAL       = 26
	ManifestParserAND            = 27
	ManifestParserOR             = 28
	ManifestParserINC            = 29
	ManifestParserDEC            = 30
	ManifestParserADD            = 31
	ManifestParserSUB            = 32
	ManifestParserMUL            = 33
	ManifestParserDIV            = 34
	ManifestParserBITAND         = 35
	ManifestParserBITOR          = 36
	ManifestParserCARET          = 37
	ManifestParserMOD            = 38
	ManifestParserDQUOTE         = 39
	ManifestParserUppercase      = 40
	ManifestParserLINE_COMMENT   = 41
	ManifestParserSPACE          = 42
	ManifestParserNL             = 43
	ManifestParserSTRING_LITERAL = 44
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

func (s *SectionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SectionContext) SPACE() antlr.TerminalNode {
	return s.GetToken(ManifestParserSPACE, 0)
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
		p.SetState(15)
		p.Match(ManifestParserKey)
	}
	{
		p.SetState(16)
		p.Match(ManifestParserCOLON)
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ManifestParserSPACE {
		{
			p.SetState(17)
			p.Match(ManifestParserSPACE)
		}

	}
	{
		p.SetState(20)
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

func (s *ValueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ManifestParserSEMI, 0)
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
		p.SetState(22)
		p.Match(ManifestParserOTHERS)
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ManifestParserSEMI {
		{
			p.SetState(23)
			p.Match(ManifestParserSEMI)
		}
		{
			p.SetState(24)
			p.Match(ManifestParserVERSION)
		}
		{
			p.SetState(25)
			p.Match(ManifestParserASSIGN)
		}
		{
			p.SetState(26)
			p.Match(ManifestParserSTRING_LITERAL)
		}

	}

	return localctx
}
