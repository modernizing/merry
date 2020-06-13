package domain

import "fmt"

func FromCSV(csv [][]string) []MavenDependency {
	var results []MavenDependency
	header := csv[0]
	if len(header) != 4 {
		fmt.Println("not valid igso map csv")
		return nil
	}
	if header[0] != "origin" {
		fmt.Println("not valid igso map csv")
		return nil
	}
	body := csv[1:]
	for _, item := range body {
		dep := MavenDependency{
			GroupId:    item[1],
			ArtifactId: item[2],
			Version:    item[3],
		}
		results = append(results, dep)
	}

	return results
}
