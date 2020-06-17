package domain

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetVersionByDepName(t *testing.T) {
	g := NewGomegaWithT(t)

	var deps []MavenDependency
	dependency := MavenDependency{
		Version:    "1.0.0",
		GroupId:    "com.phodal",
		ArtifactId: "merry",
	}
	deps = append(deps, dependency)

	result := RemoveDuplicate(deps)
	g.Expect(len(result)).To(Equal(1))
}

func TestShouldGetVersionByName(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BySlashFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.Version).To(Equal(`1.1.0`))
}

func TestShouldGetArtifactId(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BySlashFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.ArtifactId).To(Equal(`pax-exam`))
}

func Test_ShouldSupportCustomPackageNameLength(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 0)
	g.Expect(result.Version).To(Equal(`6.0.0.20121109`))
}

func Test_ShouldGetTwoLengthGroupIdWhen2(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 2)
	g.Expect(result.GroupId).To(Equal(`com.csii`))
	g.Expect(result.ArtifactId).To(Equal(`pe-http`))
}

func Test_ShouldHandleOneLengthGroupIdWhen2(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`rxjava-3.0.4.jar`, 2)
	g.Expect(result.GroupId).To(Equal(`rxjava`))
	g.Expect(result.ArtifactId).To(Equal(`rxjava`))
}

func Test_ShouldGetOneLengthGroupIdWhen2(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 1)
	g.Expect(result.GroupId).To(Equal(`com`))
	g.Expect(result.ArtifactId).To(Equal(`csii-pe-http`))
}

func Test_ShouldBuildVersionFromManifestPackage(t *testing.T) {
	g := NewGomegaWithT(t)

	var pkgs []JavaPackage
	javaPackage := JavaPackage{
		Name:         "com.csii.pe.http",
		StartVersion: "5.0.0_201",
	}
	pkgs = append(pkgs, javaPackage)

	dependencies := FromPackage(pkgs, nil)
	g.Expect(dependencies[0].Version).To(Equal("5.0.0_201"))
	g.Expect(dependencies[0].GroupId).To(Equal("com.csii"))
	g.Expect(dependencies[0].ArtifactId).To(Equal("pe-http"))
}

func Test_ShouldBuildVersionFromManifestPackageForSingleeName(t *testing.T) {
	g := NewGomegaWithT(t)

	var pkgs []JavaPackage
	javaPackage := JavaPackage{
		Name:         "ognl",
		StartVersion: "5.0.0_201",
	}
	pkgs = append(pkgs, javaPackage)

	dependencies := FromPackage(pkgs, nil)
	g.Expect(dependencies[0].Version).To(Equal("5.0.0_201"))
	g.Expect(dependencies[0].GroupId).To(Equal("ognl"))
	g.Expect(dependencies[0].ArtifactId).To(Equal("ognl"))
}

func Test_ShouldBuildDepWithMapFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	ognlDep := MavenDependency{GroupId: "ognl", ArtifactId: "ognl", Version: "5.0"}

	var depmap = make(map[string]MavenDependency)
	depmap["ognl"] = ognlDep

	var pkgs []JavaPackage
	pkgs = append(pkgs, JavaPackage{Name: "ognl", StartVersion: "5.0.0_201"})

	dependencies := FromPackage(pkgs, depmap)
	g.Expect(dependencies[0].GroupId).To(Equal("ognl"))
	g.Expect(dependencies[0].GroupId).To(Equal("ognl"))
	g.Expect(dependencies[0].Version).To(Equal("5.0"))
}

func Test_ShouldBuildDepWithMapFilterRemoveDep(t *testing.T) {
	g := NewGomegaWithT(t)

	ognlDep := MavenDependency{GroupId: "ognl", ArtifactId: "ognl", Version: "5.0"}

	var depmap = make(map[string]MavenDependency)
	depmap["ognl.bean"] = ognlDep
	depmap["ognl.factory"] = ognlDep

	var pkgs []JavaPackage
	pkgs = append(pkgs, JavaPackage{Name: "ognl.bean", StartVersion: "5.0.0_201"})
	pkgs = append(pkgs, JavaPackage{Name: "ognl.factory", StartVersion: "5.0.0_201"})
	pkgs = append(pkgs, JavaPackage{Name: "factory", StartVersion: "5.0.0_201"})

	dependencies := FromPackage(pkgs, depmap)
	g.Expect(len(dependencies)).To(Equal(2))
	g.Expect(dependencies[0].GroupId).To(Equal("factory"))
	g.Expect(dependencies[1].GroupId).To(Equal("ognl"))
	g.Expect(dependencies[1].Version).To(Equal("5.0"))
}

func Test_ShouldBuildByPackageWhenFourDot(t *testing.T) {
	g := NewGomegaWithT(t)

	dep := ByPackage("com.phodal.chapi.zero", 2)

	g.Expect(dep.GroupId).To(Equal("com.phodal"))
	g.Expect(dep.ArtifactId).To(Equal("chapi-zero"))
}

func Test_ShouldBuildByPackageWhenTwoDot(t *testing.T) {
	g := NewGomegaWithT(t)

	dep := ByPackage("log4j", 2)

	g.Expect(dep.GroupId).To(Equal("log4j"))
	g.Expect(dep.ArtifactId).To(Equal("log4j"))
}
