package cmd

import (
	"fmt"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/spf13/cobra"
)

type ChecksumConfig struct {
	Path    string
	WithGit bool
}

var (
	checksumConfig ChecksumConfig
)


func init() {
	checksumCmd.SetOut(output)

	checksumCmd.PersistentFlags().StringVarP(&checksumConfig.Path, "path", "p", "", "path")

	rootCmd.AddCommand(checksumCmd)
}

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "c",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		result, err := bundle.HashFileMD5(path)
		fmt.Fprintf(output, result)
		if err != nil {
			fmt.Println(err)
		}
	},
}
