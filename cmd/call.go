package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/phodal/igso/pkg/application/manifest"
	"github.com/phodal/igso/pkg/domain"
	"github.com/phodal/igso/pkg/infrastructure/csvconv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

			abs, _ := filepath.Abs("./manifest-map.json")
			content, err := ioutil.ReadFile(abs)

			var manifests []domain.IgsoManifest
			_ = json.Unmarshal(content, &manifests)
			dData := manifestForD3(manifests)
			dContent, err := json.Marshal(dData)

			ioutil.WriteFile("output.json", []byte(dContent), os.ModePerm)

			http.HandleFunc("/manifest-map.json", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(dContent)
			})

			http.Handle("/", http.FileServer(box))
			http.ListenAndServe(":3000", nil)

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
		f, _ := os.Create("call.dot")
		w := bufio.NewWriter(f)
		_, _ = w.WriteString("di" + graph.String())
		_ = w.Flush()

		acmd := exec.Command("dot", "-Tsvg", "call.dot", "-o", "call.svg")
		out, err := acmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			log.Fatalf("Cmd.Run() failed with %s\n", err)
		}

	},
}

type DData struct {
	Nodes []DNode `json:"nodes,omitempty"`
	Links []DLink `json:"links,omitempty"`
}

type DNode struct {
	ID    string `json:"id,omitempty"`
	Group int    `json:"group,omitempty"`
}

type DLink struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	Value  int    `json:"value,omitempty"`
}

func manifestForD3(manifests []domain.IgsoManifest) DData {
	var data DData
	nodeMap := make(map[string]DNode)
	sourceTargetMap := make(map[string]int)
	var links []DLink
	for _, mani := range manifests {
		nodeMap[mani.PackageName] = DNode{
			ID:    mani.PackageName,
			Group: 1,
		}
		for _, pkg := range mani.ImportPackage {
			if pkg.Name != "" {
				nodeMap[pkg.Name] = DNode{
					ID:    pkg.Name,
					Group: 1,
				}
				links = append(links, DLink{
					Source: mani.PackageName,
					Target: pkg.Name,
					Value:  1,
				})
				sourceTargetMap[mani.PackageName+".igso."+pkg.Name]++
			}
		}
	}

	for _, value := range nodeMap {
		data.Nodes = append(data.Nodes, value)
	}
	for _, link := range links {
		link.Value = sourceTargetMap[link.Source+".igso."+link.Target]
	}

	data.Links = links
	return data
}
