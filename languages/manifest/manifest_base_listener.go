// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Manifest

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseManifestListener is a complete listener for a parse tree produced by ManifestParser.
type BaseManifestListener struct{}

var _ ManifestListener = &BaseManifestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseManifestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseManifestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseManifestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseManifestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMf is called when production mf is entered.
func (s *BaseManifestListener) EnterMf(ctx *MfContext) {}

// ExitMf is called when production mf is exited.
func (s *BaseManifestListener) ExitMf(ctx *MfContext) {}

// EnterSection is called when production section is entered.
func (s *BaseManifestListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseManifestListener) ExitSection(ctx *SectionContext) {}

// EnterSection_header is called when production section_header is entered.
func (s *BaseManifestListener) EnterSection_header(ctx *Section_headerContext) {}

// ExitSection_header is called when production section_header is exited.
func (s *BaseManifestListener) ExitSection_header(ctx *Section_headerContext) {}

// EnterKey_values is called when production key_values is entered.
func (s *BaseManifestListener) EnterKey_values(ctx *Key_valuesContext) {}

// ExitKey_values is called when production key_values is exited.
func (s *BaseManifestListener) ExitKey_values(ctx *Key_valuesContext) {}
