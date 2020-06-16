package manifest

import (
	"fmt"
	"github.com/phodal/merry/pkg/adapter/tequila"
	"github.com/phodal/merry/pkg/domain"
)

func BuildFullGraph(scanManifest []domain.IgsoManifest, depmap map[string]domain.MavenDependency) *tequila.FullGraph {
	fullGraph := &tequila.FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*tequila.Relation),
	}
	for _, mani := range scanManifest {
		if mani.ExportPackage != nil {
			for _, exp := range mani.ExportPackage {
				src := exp.Name
				fullGraph.NodeList[src] = src

				for _, impl := range mani.ImportPackage {
					implName := impl.Name
					fullGraph.NodeList[implName] = implName
					relation := &tequila.Relation{
						From:  src,
						To:    implName,
						Style: "\"solid\"",
					}

					fullGraph.RelationList[relation.From+"->"+relation.To] = relation
				}
			}
		} else {
			src := mani.PackageName
			fullGraph.NodeList[src] = src

			for _, impl := range mani.ImportPackage {
				if depmap != nil {
					fmt.Println(depmap[impl.Name])
					if _, ok := depmap[impl.Name]; ok {
						continue
					}
				}

				implName := impl.Name
				fullGraph.NodeList[implName] = implName
				relation := &tequila.Relation{
					From:  src,
					To:    implName,
					Style: "\"solid\"",
				}

				fullGraph.RelationList[relation.From+"->"+relation.To] = relation
			}
		}
	}

	return fullGraph
}

