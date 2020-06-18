package cmd

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/phodal/merry/cmd/cmd_util"
	"github.com/phodal/merry/cmd/config"
	"github.com/phodal/merry/pkg/application/manifest"
	visual2 "github.com/phodal/merry/pkg/application/visual"
	"github.com/phodal/merry/pkg/domain"
	"github.com/phodal/merry/pkg/infrastructure/csvconv"
	"github.com/spf13/cobra"
	"net/http"
	"path/filepath"
)

type CallConfig struct {
	Path    string
	MapFile string
	Server  bool
}

var (
	callConfig CallConfig
)

func init() {
	callCmd.SetOut(output)
	callCmd.PersistentFlags().StringVarP(&callConfig.Path, "path", "p", ".", "path")
	callCmd.PersistentFlags().StringVarP(&callConfig.MapFile, "map", "m", "", "map file")
	callCmd.PersistentFlags().BoolVarP(&callConfig.Server, "server", "s", false, "with server")
	rootCmd.AddCommand(callCmd)
}

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "show call graph for packages",
	Run: func(cmd *cobra.Command, args []string) {
		if callConfig.Server {
			box := packr.NewBox("../static")

			dContent := visual2.ManifestFileToDContent(config.CmdConfig.ReporterPath + "/manifest-map.json")

			cmd_util.WriteToCocaFile("output.json", string(dContent))

			http.HandleFunc("/output.json", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write(dContent)
			})

			http.Handle("/", http.FileServer(box))
			fmt.Fprintf(output, "localserver started: http://localhost:3000/\n")

			err := http.ListenAndServe(":3000", nil)
			if err != nil {
				fmt.Println(err)
			}

			return
		}

		path := cmd.Flag("path").Value.String()

		scanManifest := manifest.ScanByPath(path)
		var depmap = make(map[string]domain.MavenDependency)
		if callConfig.MapFile != "" {
			csv := csvconv.ParseCSV(filepath.FromSlash(callConfig.MapFile))
			_, depmap = domain.MapFromCSV(csv)
		}

		result := manifest.BuildFullGraph(scanManifest, depmap)

		graph := result.ToDot(".", func(s string) bool {
			return false
		})

		cmd_util.WriteToCocaFile("call.dot", graph.String())
		cmd_util.ConvertToSvg("call")
	},
}

