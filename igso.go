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
		Name: "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			var contents []byte
			var err error
			if c.Args().Len() > 1 {
				path := c.Args().Get(1)
				fmt.Println(path)
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

			fmt.Println(results)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
