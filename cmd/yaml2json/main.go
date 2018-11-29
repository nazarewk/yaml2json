package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/nazarewk/yaml2json"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		_, _ = fmt.Fprintf(os.Stderr, `USAGE: %s < in.json > out.yaml`, os.Args[0])
		os.Exit(1)
	}

	if err := yaml2json.Convert(yaml.YAMLToJSON, os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(2)
	}
}
