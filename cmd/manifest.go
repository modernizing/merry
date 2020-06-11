package cmd

import (
	"github.com/phodal/igso/pkg/application/manifest"
	"github.com/spf13/cobra"
)

type ManifestConfig struct {
	Path    string
	Extract bool
}

var (
	manifestConfig ManifestConfig
)

func init() {
	manifestCmd.SetOut(output)
	manifestCmd.PersistentFlags().StringVarP(&manifestConfig.Path, "path", "p", ".", "path")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.Extract, "extract", "x", false, "extract file")

	rootCmd.AddCommand(manifestCmd)
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "manifest query & map tools",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if manifestConfig.Extract {
			manifest.ExtractManifest(path)
		}
	},
}
