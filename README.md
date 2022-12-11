# rendergt

[![Go Report Card](https://goreportcard.com/badge/github.com/rawmind0/rendergt)](https://goreportcard.com/report/github.com/rawmind0/rendergt)

rendergt is a simple utility to render go templates.

## Building

`make`

## Running

`./bin/rendergt`

## Usage

Rendergt read recursively go template files from input folder passed as arguments, rendering them on stdout or output folder if configured. The rendered go templates are writen following the same input folder hierarchy. Values can be replaced on templates using multiple values files. If multiple values are provided, they will; be merged before rendering the go templates.

```
NAME:
   rendergt - rendergt [OPTIONS] <template_folders>

USAGE:
   rendergt [global options] command [command options] [arguments...]

VERSION:
   0.0.2

AUTHOR:
   Rawmind <rawmind@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d               Debug logging (default: false)
   --delims value            Template delimiters (default: "{{}}")
   --ext value, -e value     Template files extension (default: "tmpl")
   --help, -h                show help (default: false)
   --output value, -o value  Output folder (default: "stdout")
   --values value            Values yaml file with data to generate templates (default: "values.yaml")  (accepts multiple inputs)
   --version, -v             print the version (default: false)
```
