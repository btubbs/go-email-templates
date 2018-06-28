package main

import (
	"fmt"
	"os"

	emailtemplates "github.com/btubbs/go-email-templates"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "eml"
	app.Usage = "Convert txt/html email templates into Go source files."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "ingest",
			Usage: "convert .txt and .html templates to .go file",
			Action: func(c *cli.Context) error {
				dir := c.String("templatedir")
				packageName := c.String("packagename")
				file := c.String("file")
				return emailtemplates.WriteTemplatesToFile(file, dir, packageName, true)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "templatedir",
					Value: ".",
					Usage: "templates directory",
				},
				cli.StringFlag{
					Name:  "packagename",
					Value: "templates",
					Usage: "package name to use in the generated .go file",
				},
				cli.StringFlag{
					Name:  "file",
					Value: "templates.go",
					Usage: "file name to use for the generated .go file",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
