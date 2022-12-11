# rendergt

[![Go Report Card](https://goreportcard.com/badge/github.com/rawmind0/rendergt)](https://goreportcard.com/report/github.com/rawmind0/rendergt)

rendergt is a simple utility to render go templates.

## Building

`make`

## Running

`./bin/rendergt`

## Usage

rendergt read recursively go template files, from folders passed as arguments, and render them on stdout or output folder if configured. The rendered go templates are writen following the same folder hierarchy. Multiple values files can be used and they are mergered before rendering the go templates.

```
NAME:
   rendergt - rendergt [OPTIONS] <template_folders>

USAGE:
   rendergt [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Rawmind <rawmind@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d               Debug logging (default: false)
   --ext value, -e value     Template files extension (default: "tmpl")
   --help, -h                show help (default: false)
   --output value, -o value  Output folder (default: "stdout")
   --values value            Values yaml file with data to generate templates (default: "values.yaml")  (accepts multiple inputs)
   --version, -v             print the version (default: false)
```
