package cmd

import (
	"github.com/phodal/merry/pkg/application/dupsearch"
	"github.com/phodal/merry/pkg/application/maven"
	dependency "github.com/phodal/merry/pkg/domain"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DupSearchConfig struct {
	Path string
}

var (
	dupSearchConfig DupSearchConfig
)

func init() {
	dupSearch.SetOut(output)

	dupSearch.PersistentFlags().StringVarP(&dupSearchConfig.Path, "path", "p", "", "path")

	rootCmd.AddCommand(dupSearch)
}

var dupSearch = &cobra.Command{
	Use:   "dupsearch",
	Short: "build maven pom from all jars file",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		deps := dupsearch.DupSearch(filepath.FromSlash(path))
		deps = dependency.RemoveDuplicate(deps)
		result := maven.BuildByDeps(deps, dependency.MavenProject{
			Version:      "0.0.1",
			GroupId:      "com.merry",
			ArtifactId:   "test",
		})

		_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(result), os.ModePerm)
	},
}
