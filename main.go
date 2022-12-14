package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-colorable"
	"github.com/rawmind0/rendergt/pkg/rendergt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	VERSION = "dev"
)

func main() {
	log.SetOutput(colorable.NewColorableStdout())

	if err := mainErr(); err != nil {
		log.Fatal(err)
	}
}

func mainErr() error {
	app := cli.NewApp()
	app.Name = "rendergt"
	app.Version = VERSION
	app.Usage = "rendergt [OPTIONS] <template_folders>"
	app.Before = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		if !ctx.Args().Present() {
			return fmt.Errorf("No templates specified")
		}
		if len(ctx.String("values")) == 0 {
			return fmt.Errorf("Values file required")
		}
		if (len(ctx.String("delims")) % 2) != 0 {
			return fmt.Errorf("Bad delims format, try \"%s\"", rendergt.DefaultDelims)
		}
		return nil
	}
	app.Authors = []*cli.Author{
		{
			Name:  "Rawmind",
			Email: "rawmind@gmail.com",
		},
	}
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Debug logging",
		},
		&cli.StringSliceFlag{
			Name:  "values",
			Usage: "Values yaml file with data to generate templates",
			Value: cli.NewStringSlice(rendergt.DefaultValuesFile),
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Output folder",
			Value:   rendergt.DefaultOutputDir,
		},
		&cli.StringFlag{
			Name:    "ext",
			Aliases: []string{"e"},
			Usage:   "Template files extension",
			Value:   rendergt.DefaultFileExt,
		},
		&cli.StringFlag{
			Name:  "delims",
			Usage: "Template delimiters",
			Value: rendergt.DefaultDelims,
		},
	}
	return app.Run(os.Args)
}

func run(cli *cli.Context) error {
	conf := &rendergt.Config{
		InputDir:  filepath.Clean(cli.Args().First()),
		OutputDir: filepath.Clean(cli.String("output")),
		FileExt:   "." + cli.String("ext"),
		Values:    cli.StringSlice("values"),
		Delims:    cli.String("delims"),
	}
	return rendergt.Run(conf)
}
