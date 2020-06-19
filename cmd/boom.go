package cmd

import (
	"fmt"
	"github.com/phodal/merry/pkg/application/maven"
	"github.com/phodal/merry/pkg/domain"
	"github.com/phodal/merry/pkg/infrastructure/csvconv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BoomConfig struct {
	Path    string
	MapFile string
	WithGit bool
}

var (
	boomConfig BoomConfig
)

func init() {
	boomCmd.SetOut(output)
	boomCmd.PersistentFlags().StringVarP(&boomConfig.Path, "path", "p", ".", "path")
	boomCmd.PersistentFlags().StringVarP(&boomConfig.MapFile, "map", "m", "", "-m map.csv")
	boomCmd.PersistentFlags().BoolVarP(&boomConfig.WithGit, "extract", "x", false, "extract file")
	rootCmd.AddCommand(boomCmd)
}

var boomCmd = &cobra.Command{
	Use:   "boom",
	Short: "generate pom.xml from build.xml",
	Run: func(cmd *cobra.Command, args []string) {
		var contents []byte
		var err error
		path := cmd.Flag("path").Value.String()

		contents, err = ioutil.ReadFile(filepath.FromSlash(path + "/" + "build.xml"))
		if err != nil {
			fmt.Println(err)
			_ = fmt.Errorf("Failed read file: %s ", err)
		}

		var depmap = make(map[string]domain.MavenDependency)
		if boomConfig.MapFile != "" {
			csv := csvconv.ParseCSV(filepath.FromSlash(boomConfig.MapFile))
			_, depmap = domain.MapFromCSV(csv)
		}

		withPom := maven.FromAntToXml(string(contents), boomConfig.WithGit, depmap)

		_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(withPom), os.ModePerm)
	},
}
