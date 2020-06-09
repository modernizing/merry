package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var (
	output io.Writer
	rootCmd = &cobra.Command{
		Use:   "igso",
		Short: "A generator for Cobra based Applications",
		Long:  `igso`,
	}
)

func NewRootCmd(out io.Writer) *cobra.Command {
	output = out
	rootCmd.SetOut(out)
	return rootCmd
}
