package dependency

import (
	. "github.com/phodal/igso/pkg/domain"
	"regexp"
)

func ByFileName(s string) MavenDependency {
	var dependency MavenDependency
	reg := regexp.MustCompile("([a-z][a-z0-9_-]*)-([.a-z0-9_]+[.0-9a-z_]*).jar")
	result := reg.FindStringSubmatch(s)
	if len(result) >= 2 {
		dependency.ArtifactId = result[1]
		dependency.Version = result[2]
	}

	return dependency
}