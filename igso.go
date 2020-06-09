package main

import (
	. "github.com/phodal/igso/cmd"
	"github.com/phodal/igso/pkg/application/dupsearch"
	"github.com/phodal/igso/pkg/application/maven"
	dependency "github.com/phodal/igso/pkg/domain"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	output := os.Stdout
	rootCmd := NewRootCmd(output)
	_ = rootCmd.Execute()
}

func main2() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
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

				deps := dupsearch.DupSearch(path)
				deps = dependency.RemoveDuplicate(deps)
				result := maven.BuildByDeps(deps, dependency.MavenProject{"0.0.1", "com.igso", "test", "4.0.0", nil})

				_ = ioutil.WriteFile(filepath.FromSlash(path+"/pom.xml"), []byte(result), os.ModePerm)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
