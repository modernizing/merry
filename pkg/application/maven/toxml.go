package maven

import "github.com/phodal/igso/pkg/domain"

func FromAntToXml(content string, isExtract bool) string {
	newAntModel := FromAnt(content, isExtract)
	deps := newAntModel.Dependencies

	results := buildDepStr(deps)
	withPom := buildFinaly(newAntModel, results)
	return withPom
}

func buildFinaly(newAntModel dependency.MavenProject, results string) string {
	var withPom = `<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>` + newAntModel.ModelVersion + `</modelVersion>
	<groupId>` + newAntModel.GroupId + `</groupId>        
	<artifactId>` + newAntModel.ArtifactId + `</artifactId>
	<version>` + newAntModel.Version + `</version>

    <dependencies>` + results + `
    </dependencies>

</project>
`
	return withPom
}

func buildDepStr(deps []dependency.MavenDependency) string {
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
