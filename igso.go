package main

import (
	"fmt"
	"github.com/phodal/igso/pkg/application/maven"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			var contents []byte
			var err error
			var path = "."
			if c.Args().Len() > 1 {
				path = c.Args().Get(1)
				contents, err = ioutil.ReadFile(filepath.FromSlash(path + "/" + "build.xml"))
			} else {
				contents, err = ioutil.ReadFile("build.xml")
			}
			if err != nil {
				fmt.Println(err)
				_ = fmt.Errorf("Failed read file: %s ", err)
				return nil
			}

			deps := maven.FromAnt(string(contents))

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

			var withPom = `<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>

    <parent>
        <groupId></groupId>        
        <artifactId></artifactId>
        <version>1.0-SNAPSHOT</version>
    </parent>

    <artifactId></artifactId>

    <dependencies>` + results + `   
</dependencies>

</project>
`
			fmt.Println(withPom)
			_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(withPom), os.ModePerm)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
