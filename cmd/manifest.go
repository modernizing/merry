package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/phodal/igso/cmd/cmd_util"
	"github.com/phodal/igso/pkg/adapter/tequila"
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
	Path            string
	FilterName      string
	IsExtract       bool
	IsScan          bool
	ManifestVersion bool
	IsMergePackage  bool
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
			igsoManifest := manifest.ScanByFile(path)
			table := cmd_util.NewOutput(output)

			table.SetHeader([]string{"Name", "Version Info", "Export Version", "Start Version", "EndVersion", "Config"})

			for _, pkg := range igsoManifest.ImportPackage {
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
			ioutil.WriteFile("manifest-map.json", res, os.ModePerm)

			result := manifest.BuildFullGraph(scanManifest)

			if manifestConfig.IsMergePackage {
				//fans := result.SortedByFan(tequila.EmptyMergePackageFunc)
				//for _, fan := range fans {
				//	fmt.Println(fan, fan.FanIn, fan.FanOut)
				//}
				result = result.MergeHeaderFile(tequila.MergePackageFunc)
			}

			graph := result.ToDot(".", func(s string) bool {
				return false
			})
			f, _ := os.Create("dep.dot")
			w := bufio.NewWriter(f)
			_, _ = w.WriteString("di" + graph.String())
			_ = w.Flush()

			exec.Command("dot", "-Tsvg", "dep.dot", "-o", "dep.svg")
		}
	},
}

func ExtractManifest(ppath string, filter string) {
	jarPaths := infrastructure.GetJarFilesByPath(ppath)
	for _, path := range jarPaths {
		if !strings.Contains(path, filter) {
			continue
		}

		_, content, _ := bundle.GetFileFromJar(path, "MANIFEST.MF")
		pureName := path[0 : len(path)-len(".jar")]

		//filePath := path + pureName + "/META-INF/"

		_ = os.MkdirAll(filepath.FromSlash(pureName), os.ModePerm)
		_ = ioutil.WriteFile(filepath.FromSlash(pureName+"/"+"MANIFEST.MF"), []byte(content), os.ModePerm)
	}
}
