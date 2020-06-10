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

	// EnterKey_values is called when entering the key_values production.
	EnterKey_values(c *Key_valuesContext)

	// ExitMf is called when exiting the mf production.
	ExitMf(c *MfContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitKey_values is called when exiting the key_values production.
	ExitKey_values(c *Key_valuesContext)
}
