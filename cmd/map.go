package cmd

import (
	"fmt"
	"github.com/phodal/igso/pkg/application/manifest"
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type MapConfig struct {
	Path string
}

var (
	mapConfig MapConfig
)

func init() {
	mapCmd.SetOut(output)
	mapCmd.PersistentFlags().StringVarP(&mapConfig.Path, "path", "p", ".", "path")
	rootCmd.AddCommand(mapCmd)
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "generate map.csv from jar",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		jarPaths := infrastructure.GetJarFilesByPath(path, false)
		var results  = "origin,groupid,artifactId,version" + "\n"
		for _, jarPath := range jarPaths {
			manifest := manifest.ScanByFile(jarPath)
			for _, pkg := range manifest.ExportPackage {
				dep := domain.ByPackage(manifest.PackageName, 2)
				var pkgStr = pkg.Name + "," + dep.GroupId + "," + dep.ArtifactId + "," + manifest.Version
				results += pkgStr + "\n"
			}
		}
		fmt.Fprintf(output, results)
		ioutil.WriteFile("map.csv", []byte(results), os.ModePerm)
	},
}
