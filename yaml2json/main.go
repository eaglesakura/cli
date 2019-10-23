package main

import (
	"flag"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	stdin := flag.Bool("stdin", true, "Json parse from stdin buffer, example) 'cat path/to/example.yaml | yaml2json -stdin'")
	flag.Parse()

	if !*stdin {
		log.Fatal("invalid '-stdin' option")
	}

	yamlBytes, _ := ioutil.ReadAll(os.Stdin)
	if jsonBytes, err := yaml.YAMLToJSON(yamlBytes); err != nil {
		log.Fatalf("yaml parse failed, %v", err)
	} else {
		_, _ = os.Stdout.Write(jsonBytes)
	}
}
