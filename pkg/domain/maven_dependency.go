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

func FromPackage(packages []JavaPackage, depmap map[string]MavenDependency) []MavenDependency {
	var deps []MavenDependency
	for _, javaPack := range packages {
		if val, ok := depmap[javaPack.Name]; ok {
			deps = append(deps, val)
			continue;
		}

		checkSplit := strings.Split(javaPack.Name, ".")
		var dependency MavenDependency
		if len(checkSplit) <= 1 {
			dependency = ByPackage(javaPack.Name, 0)
			dependency.GroupId = javaPack.Name
		} else {
			dependency = ByPackage(javaPack.Name, 2)
		}
		mavenDep := MavenDependency{
			Version:    javaPack.StartVersion,
			GroupId:    dependency.GroupId,
			ArtifactId: dependency.ArtifactId,
		}

		deps = append(deps, mavenDep)
	}

	return RemoveDuplicate(deps)
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
	reg := regexp.MustCompile("([.a-z0-9]+[.0-9a-z]*)[_-]([.a-z0-9]+[.0-9a-z]*).jar")
	result := reg.FindStringSubmatch(s)
	if len(result) >= 2 {
		if strings.Contains(result[1], ".") {
			split := strings.Split(result[1], ".")
			dependency.GroupId = strings.Join(split[0:groupIdLength], ".")
			dependency.ArtifactId = strings.Join(split[groupIdLength:], "-")
		} else {
			dependency.GroupId = result[1]
			dependency.ArtifactId = result[1]
		}
		dependency.Version = result[2]
	} else {
		dependency.GroupId = s
	}

	return dependency
}

func ByPackage(pkg string, groupIdLength int) MavenDependency {
	var dependency MavenDependency
	split := strings.Split(pkg, ".")

	if len(split) <= 2 {
		dependency.GroupId = pkg
		dependency.ArtifactId = pkg
		return dependency
	}
	dependency.GroupId = strings.Join(split[0:groupIdLength], ".")
	dependency.ArtifactId = strings.Join(split[groupIdLength:], "-")

	return dependency
}

