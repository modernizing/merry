package dependency

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
		depMap[dep.GroupId + "." + dep.ArtifactId] = dep
	}

	var depArray []MavenDependency
	for _, value := range depMap {
		depArray = append(depArray, value)
	}

	return depArray;
}
