// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Manifest

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ManifestListener is a complete listener for a parse tree produced by ManifestParser.
type ManifestListener interface {
	antlr.ParseTreeListener

	// EnterMf is called when entering the mf production.
	EnterMf(c *MfContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterPkg is called when entering the pkg production.
	EnterPkg(c *PkgContext)

	// EnterConfigAssign is called when entering the configAssign production.
	EnterConfigAssign(c *ConfigAssignContext)

	// EnterAssignKey is called when entering the assignKey production.
	EnterAssignKey(c *AssignKeyContext)

	// EnterAssignValue is called when entering the assignValue production.
	EnterAssignValue(c *AssignValueContext)

	// ExitMf is called when exiting the mf production.
	ExitMf(c *MfContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitPkg is called when exiting the pkg production.
	ExitPkg(c *PkgContext)

	// ExitConfigAssign is called when exiting the configAssign production.
	ExitConfigAssign(c *ConfigAssignContext)

	// ExitAssignKey is called when exiting the assignKey production.
	ExitAssignKey(c *AssignKeyContext)

	// ExitAssignValue is called when exiting the assignValue production.
	ExitAssignValue(c *AssignValueContext)
}
