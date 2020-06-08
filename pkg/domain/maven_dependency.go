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
