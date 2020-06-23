package cmd

import (
	"fmt"
	"github.com/phodal/merry/pkg/application/manifestapp"
	"github.com/phodal/merry/pkg/infrastructure"
	"github.com/spf13/cobra"
)

type FixConfig struct {
	Path string
}

func init() {
	fixCmd.SetOut(output)
	fixCmd.PersistentFlags().StringVarP(&manifestConfig.Path, "path", "p", ".", "path")
	rootCmd.AddCommand(fixCmd)
}

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "fix jar naming issue",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		jarPaths := infrastructure.GetJarFilesByPath(path, false)
		for _, jarPath := range jarPaths {
			manifest := manifestapp.ScanByFile(jarPath)
			if manifest.HasValidJarInfo() {
				newPkgName := manifest.BuildJarFileName()
				basePath := infrastructure.GetJarPath(jarPath)
				dst := basePath + newPkgName
				if jarPath != dst {
					_, err := infrastructure.Copy(jarPath, dst)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Fprintln(output, "rename: "+jarPath+" -> "+dst)
				}
			}
		}
	},
}
