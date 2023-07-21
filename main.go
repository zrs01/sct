package main

import (
	"fmt"
	"os"

	"srcode/internal/generator"

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
			Name:        "namespace",
			Aliases:     []string{"n"},
			Usage:       "namespace, e.g. MyProject.Service",
			Required:    false,
			Destination: &option.Namespace,
		},
		&cli.StringFlag{
			Name:        "module",
			Aliases:     []string{"m"},
			Usage:       "module name, e.g. User",
			Required:    false,
			Destination: &option.ModuleName,
		},
		&cli.StringFlag{
			Name:        "entity",
			Aliases:     []string{"e"},
			Usage:       "C# file name (separated by comma), wildcard supported, e.g. *entity*.cs",
			Required:    false,
			Destination: &option.CsharpFile,
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

	// outputFlag := func(file *string) *cli.StringFlag {
	// 	return &cli.StringFlag{
	// 		Name:        "output",
	// 		Aliases:     []string{"o"},
	// 		Usage:       "output file",
	// 		Required:    false,
	// 		Destination: file,
	// 	}
	// }

	// templateFlag := func(file *string) *cli.StringFlag {
	// 	return &cli.StringFlag{
	// 		Name:        "template",
	// 		Aliases:     []string{"t"},
	// 		Usage:       "template file",
	// 		Required:    true,
	// 		Destination: file,
	// 	}
	// }

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var ns, ifile, tfile string
	// 	return &cli.Command{
	// 		Name:  "cs",
	// 		Usage: "Tempalte using C# class",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "namespace",
	// 				Aliases:     []string{"n"},
	// 				Usage:       "namespace, e.g. MyProject.Service",
	// 				Required:    true,
	// 				Destination: &ns,
	// 			},
	// 			&cli.StringFlag{
	// 				Name:        "input",
	// 				Aliases:     []string{"i"},
	// 				Usage:       "entity file name (separated by comma), wildcard supported, e.g. *entity*.cs",
	// 				Required:    true,
	// 				Destination: &ifile,
	// 			},
	// 			templateFlag(&tfile),
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			return generator.ByCs(ns, ifile, tfile)
	// 		},
	// 	}
	// }())

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var ns, ifile, tfile string
	// 	return &cli.Command{
	// 		Name:  "dto",
	// 		Usage: "Convert entity to DTO",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "namespace",
	// 				Aliases:     []string{"n"},
	// 				Usage:       "namespace, e.g. MyProject.Service",
	// 				Required:    true,
	// 				Destination: &ns,
	// 			},
	// 			&cli.StringFlag{
	// 				Name:        "input",
	// 				Aliases:     []string{"i"},
	// 				Usage:       "entity file name (separated by comma), wildcard supported, e.g. *entity*.cs",
	// 				Required:    true,
	// 				Destination: &ifile,
	// 			},
	// 			templateFlag(&tfile),
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			return generator.ConvertCsDto(ns, ifile, tfile)
	// 		},
	// 	}
	// }())

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var ifile, tfile string
	// 	return &cli.Command{
	// 		Name:  "model",
	// 		Usage: "Convert entity to typescript model",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "input",
	// 				Aliases:     []string{"i"},
	// 				Usage:       "entity file name (separated by comma), wildcard supported, e.g. *entity*.cs",
	// 				Required:    true,
	// 				Destination: &ifile,
	// 			},
	// 			templateFlag(&tfile),
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			return generator.ConvertTsModel(ifile, tfile)
	// 		},
	// 	}
	// }())

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var ifile, tfile string
	// 	return &cli.Command{
	// 		Name:  "formgroup",
	// 		Usage: "Convert entity to typescript FormGroup",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "input",
	// 				Aliases:     []string{"i"},
	// 				Usage:       "entity file name (separated by comma), wildcard supported, e.g. *entity*.cs",
	// 				Required:    true,
	// 				Destination: &ifile,
	// 			},
	// 			templateFlag(&tfile),
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			return generator.ConvertTsFormGroup(ifile, tfile)
	// 		},
	// 	}
	// }())

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var ns, mname, tfile string
	// 	return &cli.Command{
	// 		Name:  "module",
	// 		Usage: "Generate project module",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "namespace",
	// 				Aliases:     []string{"n"},
	// 				Usage:       "namespace, e.g. MyProject.Service",
	// 				Required:    true,
	// 				Destination: &ns,
	// 			},
	// 			&cli.StringFlag{
	// 				Name:        "module",
	// 				Aliases:     []string{"m"},
	// 				Usage:       "module name, e.g. User",
	// 				Required:    true,
	// 				Destination: &mname,
	// 			},
	// 			templateFlag(&tfile),
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			return generator.ConvertModule(ns, mname, tfile)
	// 		},
	// 	}
	// }())

	// cliapp.Commands = append(cliapp.Commands, func() *cli.Command {
	// 	var moduleName string
	// 	return &cli.Command{
	// 		Name:  "print",
	// 		Usage: "Generate print service",
	// 		Flags: []cli.Flag{
	// 			&cli.StringFlag{
	// 				Name:        "name",
	// 				Aliases:     []string{"n"},
	// 				Usage:       "module name",
	// 				Required:    true,
	// 				Destination: &moduleName,
	// 			},
	// 		},
	// 		Action: func(c *cli.Context) error {
	// 			config.InitConfig(cfile)
	// 			if err := generator.ConvertReport(moduleName); err != nil {
	// 				return eris.Wrapf(err, "failed to generate report with %s", moduleName)
	// 			}
	// 			return nil
	// 		},
	// 	}
	// }())

	if err := cliapp.Run(os.Args); err != nil {
		fmt.Println(eris.ToString(err, true))
	}
}
