package main

import (
	"fmt"
	"github.com/phodal/igso/pkg/application/dupsearch"
	"github.com/phodal/igso/pkg/application/maven"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "boom",
			Usage: "make an explosive entrance",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "path", Aliases: []string{"p"},
				},
				&cli.BoolFlag{
					Name: "extract", Aliases: []string{"x"},
				},
			},
			Action: func(c *cli.Context) error {
				err, done := outputResult(c)
				if done {
					return err
				}
				return nil
			},
		},
		{
			Name:  "dupsearch",
			Usage: "search unkonw jar source",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "path", Aliases: []string{"p"},
				},
			},
			Action: func(c *cli.Context) error {
				var path = "."
				if c.String("path") != "" {
					path = c.String("path")
				}

				dupsearch.DupSearch(path)

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func outputResult(c *cli.Context) (error, bool) {
	var contents []byte
	var err error
	var path = "."
	if c.String("path") != "" {
		path = c.String("path")
	}

	isExtract := c.Bool("extract")
	if c.Args().Len() > 1 {
		contents, err = ioutil.ReadFile(filepath.FromSlash(path + "/" + "build.xml"))
	} else {
		contents, err = ioutil.ReadFile("build.xml")
	}
	if err != nil {
		fmt.Println(err)
		_ = fmt.Errorf("Failed read file: %s ", err)
		return nil, true
	}

	withPom := maven.FromAntToXml(string(contents), isExtract)

	_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(withPom), os.ModePerm)
	return nil, false
}
