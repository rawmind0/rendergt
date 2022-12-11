# rendergt

[![Go Report Card](https://goreportcard.com/badge/github.com/rawmind0/tq)](https://goreportcard.com/report/github.com/rawmind0/tq)

rendergt is a simple utility to render go templates.

## Building

`make`

## Running

`./bin/rendergt`

## Usage

rendergt read recursively go template files, from folders passed as arguments, and render them on stdout or output folder if configured. The rendered go templates are writen following the same folder hierarchy. Multiple values files can be used and they are mergered before rendering the go templates.

```
NAME:
   tq - tq [OPTIONS] templates

USAGE:
   tq [global options] command [command options] [arguments...]

VERSION:
   dev

AUTHOR:
   Rawmind <rawmind@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d               Debug logging (default: false)
   --ext value, -e value     Template files extension (default: "tmpl")
   --help, -h                show help (default: false)
   --output value, -o value  Output folder (default: "stdout")
   --values value            Values yaml file with data to generate templates (default: "values.yanl")  (accepts multiple inputs)
   --version, -v             print the version (default: false)
```