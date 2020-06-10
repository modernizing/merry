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

	// ExitMf is called when exiting the mf production.
	ExitMf(c *MfContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
