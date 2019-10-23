package main

import (
	"github.com/ghodss/yaml"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// read json from stdio
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		log.Fatal("'stdin' is invalid\nexample) cat 'path/to/yaml | yaml2json'")
	}

	yamlBytes, _ := ioutil.ReadAll(os.Stdin)
	if jsonBytes, err := yaml.YAMLToJSON(yamlBytes); err != nil {
		log.Fatalf("yaml parse failed, %v", err)
	} else {
		_, _ = os.Stdout.Write(jsonBytes)
	}
}
