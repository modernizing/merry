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

// EnterValue is called when production value is entered.
func (s *BaseManifestListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseManifestListener) ExitValue(ctx *ValueContext) {}

// EnterConfigAssign is called when production configAssign is entered.
func (s *BaseManifestListener) EnterConfigAssign(ctx *ConfigAssignContext) {}

// ExitConfigAssign is called when production configAssign is exited.
func (s *BaseManifestListener) ExitConfigAssign(ctx *ConfigAssignContext) {}

// EnterAssignKey is called when production assignKey is entered.
func (s *BaseManifestListener) EnterAssignKey(ctx *AssignKeyContext) {}

// ExitAssignKey is called when production assignKey is exited.
func (s *BaseManifestListener) ExitAssignKey(ctx *AssignKeyContext) {}

// EnterAssignValue is called when production assignValue is entered.
func (s *BaseManifestListener) EnterAssignValue(ctx *AssignValueContext) {}

// ExitAssignValue is called when production assignValue is exited.
func (s *BaseManifestListener) ExitAssignValue(ctx *AssignValueContext) {}
