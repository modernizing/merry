package manifest

import (
	"github.com/phodal/igso/pkg/adapter/tequila"
	"github.com/phodal/igso/pkg/domain"
)

func BuildFullGraph(scanManifest []dependency.IgsoManifest) *tequila.FullGraph {
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

