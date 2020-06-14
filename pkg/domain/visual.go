package domain

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

func VisualFromManifest(manifests []IgsoManifest) DData {
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
