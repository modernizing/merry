package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/merry/cmd/cmd_util"
	"github.com/phodal/merry/pkg/adapter/tequila"
	"github.com/phodal/merry/pkg/application/manifest"
	"github.com/phodal/merry/pkg/infrastructure"
	"github.com/phodal/merry/pkg/infrastructure/bundle"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ManifestConfig struct {
	Path            string
	FilterName      string
	IsExtract       bool
	IsScan          bool
	ManifestVersion bool
	IsMergePackage  bool
	ExcludeSource   bool
}

var (
	manifestConfig ManifestConfig
)

func init() {
	manifestCmd.SetOut(output)
	manifestCmd.PersistentFlags().StringVarP(&manifestConfig.Path, "path", "p", ".", "path")
	manifestCmd.PersistentFlags().StringVarP(&manifestConfig.FilterName, "filter", "f", "", "filter")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.IsExtract, "extract", "x", false, "extract manifest file from jar")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.IsScan, "scan", "s", false, "scan manifest file to graphviz")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.ManifestVersion, "version", "v", false, "show manifest version info of jar ")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.IsMergePackage, "merge", "m", false, "is merge package")
	manifestCmd.PersistentFlags().BoolVarP(&manifestConfig.ExcludeSource, "excludeSource", "e", false, "is with exclude source file")

	rootCmd.AddCommand(manifestCmd)
}

var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "manifest query & map tools",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if manifestConfig.IsExtract {
			ExtractManifest(path, manifestConfig.FilterName)
		}
		if manifestConfig.ManifestVersion {
			if !strings.HasSuffix(path, ".jar") {
				fmt.Fprintf(output, "path: "+path+" lost jar files")
				return
			}
			merryManifest := manifest.ScanByFile(path)
			table := cmd_util.NewOutput(output)

			table.SetHeader([]string{"Name", "Version Info", "Export Version", "Start Version", "EndVersion", "Config"})

			for _, pkg := range merryManifest.ImportPackage {
				results := ""
				for _, conf := range pkg.Config {
					if conf.Key != "version" {
						results += conf.Key + ": " + conf.Value + " "
					}
				}
				table.Append([]string{pkg.Name, pkg.VersionInfo, pkg.ExportVersion, pkg.StartVersion, pkg.EndVersion, results})
			}
			table.Render()
		}
		if manifestConfig.IsScan {
			scanManifest := manifest.ScanByPath(path)
			res, _ := json.MarshalIndent(scanManifest, "", "\t")
			cmd_util.WriteToCocaFile("manifest-map.json", string(res))

			result := manifest.BuildFullGraph(scanManifest, nil)

			if manifestConfig.IsMergePackage {
				result = result.MergeHeaderFile(tequila.MergePackageFunc)
			}

			graph := result.ToDot(".", func(s string) bool {
				return false
			})

			cmd_util.WriteToCocaFile("dep.dot", "di" + graph.String())
			cmd_util.ConvertToSvg("dep")
		}
	},
}

func ExtractManifest(path string, filter string) {
	jarPaths := infrastructure.GetJarFilesByPath(path, manifestConfig.ExcludeSource)
	for _, path := range jarPaths {
		if !strings.Contains(path, filter) {
			continue
		}

		_, content, _ := bundle.GetFileFromJar(path, "MANIFEST.MF")
		pureName := path[0 : len(path)-len(".jar")]

		_ = os.MkdirAll(filepath.FromSlash(pureName), os.ModePerm)
		_ = ioutil.WriteFile(filepath.FromSlash(pureName+"/"+"MANIFEST.MF"), []byte(content), os.ModePerm)
	}
}
