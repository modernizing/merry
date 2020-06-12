package domain

import (
	"regexp"
	"strings"
)

type MavenDependency struct {
	Version    string
	GroupId    string
	ArtifactId string
}

type MavenProject struct {
	Version      string
	GroupId      string
	ArtifactId   string
	ModelVersion string
	Dependencies []MavenDependency
}

func RemoveDuplicate(deps []MavenDependency) []MavenDependency {
	depMap := make(map[string]MavenDependency)
	for _, dep := range deps {
		depMap[dep.GroupId+"."+dep.ArtifactId] = dep
	}

	var depArray []MavenDependency
	for _, value := range depMap {
		depArray = append(depArray, value)
	}

	return depArray
}

func FromPackage(version string, packages []JavaPackage) []MavenDependency {
	var deps []MavenDependency
	for _, javaPack := range packages {
		dependency := ByPackage(javaPack.Name, 2)
		mavenDep := MavenDependency{
			Version:    version,
			GroupId:    dependency.GroupId,
			ArtifactId: dependency.ArtifactId,
		}

		deps = append(deps, mavenDep)
	}

	return deps
}

func BySlashFileName(s string) MavenDependency {
	var dependency MavenDependency
	reg := regexp.MustCompile("([a-z][a-z0-9_-]*)-([.a-z0-9_]+[.0-9a-z_]*).jar")
	result := reg.FindStringSubmatch(s)
	if len(result) >= 2 {
		dependency.ArtifactId = result[1]
		dependency.Version = result[2]
	}

	return dependency
}

func ByFileName(s string, groupIdLength int) MavenDependency {
	var dependency MavenDependency
	reg := regexp.MustCompile("([.a-z0-9]+[.0-9a-z]*)_([.a-z0-9]+[.0-9a-z]*).jar")
	result := reg.FindStringSubmatch(s)
	if len(result) >= 2 {
		if strings.Contains(result[1], ".") {
			split := strings.Split(result[1], ".")
			dependency.GroupId = strings.Join(split[0:groupIdLength], ".")
			dependency.ArtifactId = strings.Join(split[groupIdLength:], "-")
		} else {
			dependency.ArtifactId = result[1]
		}
		dependency.Version = result[2]
	}

	return dependency
}

func ByPackage(pkg string, groupIdLength int) MavenDependency {
	var dependency MavenDependency

	split := strings.Split(pkg, ".")
	dependency.GroupId = strings.Join(split[0:groupIdLength], ".")
	dependency.ArtifactId = strings.Join(split[groupIdLength:], "-")

	return dependency
}
