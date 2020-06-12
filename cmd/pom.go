package cmd

import (
	"fmt"
	"github.com/phodal/igso/pkg/application/manifest"
	"github.com/phodal/igso/pkg/application/maven"
	. "github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure"
	"github.com/phodal/igso/pkg/infrastructure/bundle"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PomConfig struct {
	Path string
}

var (
	pomConfig PomConfig
)

func init() {
	pomCmd.SetOut(output)
	pomCmd.PersistentFlags().StringVarP(&pomConfig.Path, "path", "p", ".", "path")
	rootCmd.AddCommand(pomCmd)
}

var pomCmd = &cobra.Command{
	Use:   "pom",
	Short: "generate pom file from jar file",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		var jarFileFilter = func(path string) bool {
			return strings.HasSuffix(path, ".jar")
		}

		jarPaths := infrastructure.GetFilesByFilter(path, jarFileFilter)
		for _, jarPath := range jarPaths {
			pathSplit := strings.Split(jarPath, string(os.PathSeparator))
			jarName := pathSplit[len(pathSplit)-1]
			pomName := strings.ReplaceAll(jarName, ".jar", ".pom")

			_, content, _ := bundle.GetFileFromJar(jarPath, "MANIFEST.MF")
			dep := ByFileName(jarName, 2)
			manifest := manifest.ProcessManifest(content, "MANIFEST.MF")

			//all := strings.ReplaceAll(dep.GroupId, ".", "/")
			//basePkgPath := "maven/" + all
			//_ = os.MkdirAll(filepath.FromSlash(basePkgPath), os.ModePerm)

			importDeps := FromPackage(manifest.ImportPackage)
			pomfile := maven.BuildByDeps(importDeps, MavenProject{
				Version:    dep.Version,
				GroupId:    dep.GroupId,
				ArtifactId: dep.ArtifactId,
			})

			err := ioutil.WriteFile(filepath.FromSlash(path+ "/" + pomName), []byte(pomfile), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

//
//func generateMavenDataXml(dep dependency2.MavenDependency, path string) {
//	xml := `<?xml version="1.0" encoding="UTF-8"?>
//<metadata modelVersion="1.1.0">
//    <groupId>` + dep.GroupId + `</groupId>
//    <artifactId>` + dep.ArtifactId + `</artifactId>
//    <versioning>
//        <latest>` + dep.Version + `</latest>
//        <release>` + dep.Version + `</release>
//        <versions>
//            <version>` + dep.Version + `</version>
//        </versions>
//        <lastUpdated>20200612091552</lastUpdated>
//    </versioning>
//</metadata>`
//
//	ioutil.WriteFile("mavendata.xml", []byte(xml), os.ModePerm)
//}
