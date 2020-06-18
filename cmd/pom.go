package cmd

import (
	"fmt"
	"github.com/phodal/merry/pkg/application/manifest"
	"github.com/phodal/merry/pkg/application/maven"
	domain "github.com/phodal/merry/pkg/domain"
	"github.com/phodal/merry/pkg/infrastructure"
	"github.com/phodal/merry/pkg/infrastructure/bundle"
	"github.com/phodal/merry/pkg/infrastructure/csvconv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PomConfig struct {
	Path    string
	MapFile string
}

var (
	pomConfig PomConfig
)

func init() {
	pomCmd.SetOut(output)
	pomCmd.PersistentFlags().StringVarP(&pomConfig.Path, "path", "p", ".", "path")
	pomCmd.PersistentFlags().StringVarP(&pomConfig.MapFile, "map", "m", "", "map file")
	rootCmd.AddCommand(pomCmd)
}

var pomCmd = &cobra.Command{
	Use:   "pom",
	Short: "generate pom file from jar file",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()

		jarPaths := infrastructure.GetJarFilesByPath(path, false)
		for _, jarPath := range jarPaths {
			pathSplit := strings.Split(jarPath, string(os.PathSeparator))
			jarName := pathSplit[len(pathSplit)-1]

			_, content, _ := bundle.GetFileFromJar(jarPath, "MANIFEST.MF")
			dep := domain.ByFileName(jarName, 2)
			manifest := manifest.ProcessManifest(content, "MANIFEST.MF")

			var depmap = make(map[string]domain.MavenDependency)
			if pomConfig.MapFile != "" {
				csv := csvconv.ParseCSV(filepath.FromSlash(pomConfig.MapFile))
				_, depmap = domain.MapFromCSV(csv)
			}

			importDeps := domain.FromPackage(manifest.ImportPackage, depmap)
			pomfile := maven.BuildByDeps(importDeps, domain.MavenProject{
				Version:    dep.Version,
				GroupId:    dep.GroupId,
				ArtifactId: dep.ArtifactId,
			})

			err := ioutil.WriteFile(filepath.FromSlash(strings.ReplaceAll(jarPath, ".jar", ".pom")), []byte(pomfile), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}
