package main

import (
	"fmt"
	"os"

	"github.com/zrs01/sct/internal/generator"

	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v2"
)

var version = "development"

func main() {
	cliapp := cli.NewApp()
	cliapp.Name = "srcode"
	cliapp.Usage = "Source code template"
	cliapp.Version = version
	cliapp.Commands = []*cli.Command{}

	var cfile string
	var option generator.Option

	// global options
	cliapp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "configuration file",
			Value:       "config.yml",
			Required:    false,
			Destination: &cfile,
		},
		&cli.StringFlag{
			Name:        "data",
			Aliases:     []string{"d"},
			Usage:       "user-defined data file in yaml format",
			Required:    false,
			Destination: &option.DataFile,
		},
		&cli.StringFlag{
			Name:        "entity",
			Aliases:     []string{"e"},
			Usage:       "C# file name (separated by comma), wildcard supported, e.g. *entity*.cs",
			Required:    false,
			Destination: &option.Entity,
		},
		&cli.StringFlag{
			Name:        "template",
			Aliases:     []string{"t"},
			Usage:       "template file",
			Required:    true,
			Destination: &option.TemplateFile,
		},
	}

	cliapp.Action = func(ctx *cli.Context) error {
		generator.Output(option)
		return nil
	}

	if err := cliapp.Run(os.Args); err != nil {
		fmt.Println(eris.ToString(err, true))
	}
}
