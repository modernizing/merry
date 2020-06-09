package cmd

import (
	"fmt"
	"github.com/phodal/igso/cmd/config"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd.SetOut(output)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(output, "Igso Version: " + config.VERSION + " -- HEAD")
	},
}