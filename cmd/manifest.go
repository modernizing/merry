package cmd

import (
	"fmt"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/phodal/igso/pkg/instrastructure"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
			ExtractManifest(path)
		}
	},
}

func ExtractManifest(ppath string) {
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	jarPaths := instrastructure.GetFilesByFilter(ppath, jarFileFilter)
	fmt.Println(ppath)
	for _, path := range jarPaths {
		_, content, _ := bundle.GetFileFromJar(path, "MANIFEST.MF")
		pureName := path[len(ppath)+1 : len(path)-len(".jar")]

		filePath := "_m/" + pureName + "/META-INF/"

		_ = os.MkdirAll(filepath.FromSlash(filePath), os.ModePerm)
		_ = ioutil.WriteFile(filepath.FromSlash(filePath+"MANIFEST.MF"), []byte(content), os.ModePerm)
	}
}
