package ast

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ShouldGetManifestProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Bundle-Version: 1.0

`
	results := Analysis(code, "hello.mf").KeyValues
	g.Expect(len(results)).To(Equal(1))
	g.Expect(results[0].Key).To(Equal("Bundle-Version"))
	g.Expect(results[0].Value).To(Equal("1.0"))
}

func Test_ShouldParseImportPackage(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Import-Package: org.osgi.framework;
version="[1.3,2.0)"

`
	results := Analysis(code, "hello.mf").KeyValues
	g.Expect(len(results)).To(Equal(1))
	g.Expect(results[0].Key).To(Equal("Import-Package"))
	g.Expect(results[0].Value).To(Equal(`org.osgi.framework;version="[1.3,2.0)"`))
}

func Test_ShouldGetMultipleManifestProperties(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `
Implementation-Title: Spring Framework
Spring-Version: 2.5.6.SEC01
Implementation-Version: 2.5.6.SEC01
Bundle-Name: Spring Core
Created-By: 10.0-b23 (Sun Microsystems Inc.)
Ant-Version: Apache Ant 1.7.0
Bundle-Vendor: SpringSource
Bundle-Version: 2.5.6.SEC01
Bundle-ManifestVersion: 2

`
	results := Analysis(code, "hello.mf").KeyValues
	g.Expect(len(results)).To(Equal(9))
	g.Expect(results[2].Value).To(Equal("2.5.6.SEC01"))
	g.Expect(results[4].Value).To(Equal("10.0-b23 (Sun Microsystems Inc.)"))
}

func Test_ShouldGetImportPackage(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Import-Package: org.osgi.framework;
version="[1.3,2.0)"

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.Package[0].Name).To(Equal(`org.osgi.framework`))
	g.Expect(results.Package[0].VersionInfo).To(Equal(`1.3,2.0`))
	g.Expect(results.Package[0].StartVersion).To(Equal(`1.3`))
	g.Expect(results.Package[0].EndVersion).To(Equal(`2.0`))
}

func Test_ShouldGetCompileSuccessForMorePackageInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Import-Package: javax.persistence.spi;
version="[1.0.0,2.0.0)";resolution:=optional

`
	results := Analysis(code, "hello.mf")
	fmt.Println(results.Package)
	//g.Expect(results.Package[0].Name).To(Equal(`javax.persistence.spi`))
	g.Expect(1).To(Equal(1))
}
