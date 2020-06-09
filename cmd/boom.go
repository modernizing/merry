package cmd

import (
	"fmt"
	"github.com/phodal/igso/pkg/application/maven"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

type BoomConfig struct {
	Path    string
	WithGit bool
}

var (
	boomConfig BoomConfig
)


func init() {
	boomCmd.SetOut(output)

	boomCmd.PersistentFlags().StringVarP(&boomConfig.Path, "path", "p", ".", "path")
	boomCmd.PersistentFlags().BoolVarP(&boomConfig.WithGit, "extract", "x", false, "extract file")

	rootCmd.AddCommand(boomCmd)
}

var boomCmd = &cobra.Command{
	Use:   "boom",
	Short: "b",
	Run: func(cmd *cobra.Command, args []string) {
		var contents []byte
		var err error
		path := cmd.Flag("path").Value.String()

		contents, err = ioutil.ReadFile(filepath.FromSlash(path + "/" + "build.xml"))
		if err != nil {
			fmt.Println(err)
			_ = fmt.Errorf("Failed read file: %s ", err)
		}

		withPom := maven.FromAntToXml(string(contents), boomConfig.WithGit)

		_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(withPom), os.ModePerm)
	},
}
