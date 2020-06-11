package ast

import (
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
	code := `Import-Package: org.osgi.framework;version="[1.3,2.0)"

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
	code := `Import-Package: org.osgi.framework;version="[1.3,2.0)"

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.ImportPackage[0].Name).To(Equal(`org.osgi.framework`))
	g.Expect(results.ImportPackage[0].VersionInfo).To(Equal(`1.3,2.0`))
	g.Expect(results.ImportPackage[0].StartVersion).To(Equal(`1.3`))
	g.Expect(results.ImportPackage[0].EndVersion).To(Equal(`2.0`))
}

func Test_ShouldGetCompileSuccessForInline(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Import-Package: javax.persistence.spi;version="[1.0.0,2.0.0)"

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.ImportPackage[0].Name).To(Equal(`javax.persistence.spi`))
}

func Test_ShouldGetCompileSuccessForMorePackageInfo(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Import-Package: javax.persistence.spi;version="[1.0.0,2.0.0)";resolution:=optional

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.ImportPackage[0].Config[0].Key).To(Equal("version"))
	g.Expect(results.ImportPackage[0].Config[1].Key).To(Equal("resolution"))
	g.Expect(results.ImportPackage[0].Config[1].Value).To(Equal("optional"))
}

func Test_ShouldSuccessParseExportPackageVersion(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Export-Package: org.springframework.asm;version="2.5.6.SEC01"

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.ExportPackage[0].ExportVersion).To(Equal("2.5.6.SEC01"))
}

func Test_ShouldGetCompileComplexImport(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `
Import-Package: edu.emory.mathcs.backport.java.util.concurrent;version
 ="[3.0.0, 4.0.0)";resolution:=optional,javax.xml.transform;resolution
 :=optional,org.apache.commons.attributes;version="[2.2.0, 3.0.0)";res
 olution:=optional,org.apache.commons.collections;version="[3.2.0, 4.0
 .0)";resolution:=optional,org.apache.commons.collections.map;version=
 "[3.2.0, 4.0.0)";resolution:=optional,org.apache.commons.logging;vers
 ion="[1.0.4, 2.0.0)",org.apache.log4j;version="[1.2.15, 2.0.0)";resol
 ution:=optional,org.apache.log4j.xml;version="[1.2.15, 2.0.0)";resolu
 tion:=optional,org.aspectj.bridge;version="[1.5.4, 2.0.0)";resolution
 :=optional,org.aspectj.weaver;version="[1.5.4, 2.0.0)";resolution:=op
 tional,org.aspectj.weaver.bcel;version="[1.5.4, 2.0.0)";resolution:=o
 ptional,org.aspectj.weaver.patterns;version="[1.5.4, 2.0.0)";resoluti
 on:=optional,org.eclipse.core.runtime;common="split";resolution:=opti
 onal,org.w3c.dom;resolution:=optional,org.xml.sax;resolution:=optiona
 l

`
	results := Analysis(code, "hello.mf")
	g.Expect(len(results.ImportPackage)).To(Equal(15))
}


func Test_ShouldBuildOsgiPackageName(t *testing.T) {
	g := NewGomegaWithT(t)
	code := `Private-Package: felix-tut-8

`
	results := Analysis(code, "hello.mf")
	g.Expect(results.PackageName).To(Equal("felix-tut-8"))
}
