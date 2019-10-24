package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input := flag.String("input", "", "original files directory path, example) 'path/to/origin/images/res'")
	output := flag.String("output", "", "output files directory path, example) 'path/to/src/main/res'")
	example := flag.String("example", "", "generate example settings file, example) 'path/to/example/file.yaml'")

	flag.Parse()

	if len(*example) > 0 {
		_ = ioutil.WriteFile(*example, []byte(exampleYaml), os.ModePerm)
		return
	}

	if len(*input) == 0 {
		log.Fatal("invalid '-input' option, see '-help' option.")
	}

	if len(*output) == 0 {
		log.Fatal("invalid '-output' option, see '-help' option.")
	}

	var configure *Configure
	if config, err := parseConfigure(*input); err != nil {
		log.Fatalf("config file parse error, %v", err.Error())
	} else {
		configure = config
	}

	if len(configure.Requests) == 0 {
		log.Fatal("requests is empty")
	}

	for _, req := range configure.Requests {
		if err := req.Generate(*input, *output); err != nil {
			log.Fatalf("generate failed, %v", err)
		}
	}
}

const exampleYaml = `
#
# 1. copy original files.
#    root/
#        config.yaml
#        drawable/
#                xxhdpi/example.png
#                ...
#        mipmap/
#                xxhdpi/example.png
#                ...
# 2. Edit config.yaml file to your project.
# 3. generate mipmap.
#    $ genmipmap -input path/to/original/files/dir -output path/to/generation/dir
#
requests:
  - path: android/drawable
    platform: android
    type: drawable
    outpath: dst/Resources
    format: webp
    convert_args:
      - "-quality"
      - "100"
  - path: android/mipmap
    platform: android
    type: mipmap
    outpath: ../dst/Resources
    format: jpg
    convert_args:
      - "-quality"
      - "50"
`
