package cmd

import (
	"bufio"
	"encoding/json"
	"github.com/phodal/igso/pkg/application/manifest"
	"github.com/phodal/igso/pkg/infrastructure"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ManifestConfig struct {
	Path      string
	IsExtract bool
	IsScan    bool
}

var (
	manifestConfig ManifestConfig
)

func init() {
	manifestCmd.SetOut(output)
	manifestCmd.PersistentFlags().StringVarP(&manifestConfig.Path, "path", "p", ".", "path")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.IsExtract, "extract", "x", false, "extract manifest file from jar")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.IsScan, "scan", "s", false, "scan manifest file to graphviz")

	rootCmd.AddCommand(manifestCmd)
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "manifest query & map tools",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if manifestConfig.IsExtract {
			ExtractManifest(path)
		}
		if manifestConfig.IsScan {
			scanManifest := manifest.ScanManifest(path)
			res, _ := json.MarshalIndent(scanManifest, "", "\t")
			ioutil.WriteFile("manifest-map.json", res, os.ModePerm)

			//fmt.Println(scanManifest)
			result := manifest.BuildFullGraph(scanManifest)

			ignores := strings.Split("", ",")
			var nodeFilter = func(key string) bool {
				for _, f := range ignores {
					if key == f {
						return true
					}
				}
				return false
			}

			graph := result.ToDot(".", nodeFilter)
			f, _ := os.Create("dep.dot")
			w := bufio.NewWriter(f)
			_, _ = w.WriteString("di" + graph.String())
			_ = w.Flush()

			exec.Command("dot", "-Tsvg","dep.dot", "-o", "dep.svg")
		}
	},
}

func ExtractManifest(ppath string) {
	var jarFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	}

	jarPaths := infrastructure.GetFilesByFilter(ppath, jarFileFilter)
	for _, path := range jarPaths {
		_, content, _ := bundle.GetFileFromJar(path, "MANIFEST.MF")
		pureName := path[len(ppath)+1 : len(path)-len(".jar")]

		filePath := "_m/" + pureName + "/META-INF/"

		_ = os.MkdirAll(filepath.FromSlash(filePath), os.ModePerm)
		_ = ioutil.WriteFile(filepath.FromSlash(filePath+"MANIFEST.MF"), []byte(content), os.ModePerm)
	}
}
