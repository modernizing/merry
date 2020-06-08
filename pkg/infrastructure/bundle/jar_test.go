package bundle

import (
	"fmt"
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"testing"
)

func TestShouldReadZipList(t *testing.T) {
	g := NewGomegaWithT(t)
	path := `../../../_fixtures/bundle/sample/felix-tut-8-1.0-SNAPSHOT.jar`
	absPath := filepath.FromSlash(path)

	result := GetJarFiles(absPath)
	g.Expect(len(result)).To(Equal(11))

	var names []string
	for _, file := range result {
		names = append(names, file.Name)
	}
	g.Expect(names).To(Equal([]string{
		"META-INF/MANIFEST.MF",
		"META-INF/",
		"tutorial/",
		"tutorial/example8/",
		"META-INF/maven/",
		"META-INF/maven/org.vaadin.example.osgi/",
		"META-INF/maven/org.vaadin.example.osgi/felix-tut-8/",
		"tutorial/example8/SpellCheckerImpl.class",
		"tutorial/example8/Activator.class",
		"META-INF/maven/org.vaadin.example.osgi/felix-tut-8/pom.xml",
		"META-INF/maven/org.vaadin.example.osgi/felix-tut-8/pom.properties",
	}))
}

func TestShouldUnzipJar(t *testing.T) {
	g := NewGomegaWithT(t)

	path := `../../../_fixtures/bundle/sample/felix-tut-8-1.0-SNAPSHOT.jar`
	absPath := filepath.FromSlash(path)
	err := Unzip(absPath, filepath.FromSlash("./tmp"))
	if err != nil {
		fmt.Println(err)
	}

	_, e := os.Stat(filepath.FromSlash("./tmp/META-INF/MANIFEST.MF"))
	g.Expect(e).To(BeNil())
}
