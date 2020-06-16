package maven

import (
	"github.com/phodal/merry/pkg/domain"
)

func FromAntToXml(content string, isExtract bool, depmap map[string]domain.MavenDependency) string {
	newAntModel := FromAnt(content, isExtract)
	deps := newAntModel.Dependencies

	var javaPkgs []domain.JavaPackage
	for _, dep := range deps {
		pkg := domain.JavaPackage{
			Name: dep.GroupId + "." + dep.ArtifactId,
			StartVersion: dep.Version,
		}
		javaPkgs = append(javaPkgs, pkg)
	}

	importDeps := domain.FromPackage(javaPkgs, depmap)

	withPom := BuildByDeps(importDeps, newAntModel)
	return withPom
}

func BuildByDeps(deps []domain.MavenDependency, newAntModel domain.MavenProject) string {
	results := buildDepStr(deps)
	withPom := buildFinal(newAntModel, results)
	return withPom
}

func buildFinal(newAntModel domain.MavenProject, results string) string {
	var withPom = `<?xml version="1.0"?>
<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
         xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>` + newAntModel.GroupId + `</groupId>        
	<artifactId>` + newAntModel.ArtifactId + `</artifactId>
	<version>` + newAntModel.Version + `</version>
	<packaging>jar</packaging>

    <dependencies>` + results + `
    </dependencies>

</project>
`
	return withPom
}

func buildDepStr(deps []domain.MavenDependency) string {
	var results = ""
	for _, dep := range deps {
		var tmpl = `
      <dependency>
          <groupId>` + dep.GroupId + `</groupId>
          <artifactId>` + dep.ArtifactId + `</artifactId>
          <version>` + dep.Version + `</version>
      </dependency>
`
		results += tmpl
	}
	return results
}
