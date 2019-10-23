package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/eaglesakura/cli/commons/shell"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
)

type GcloudServiceAccount struct {
	ProjectId   string `json:"project_id"`
	ClientEmail string `json:"client_email"`
}

func main() {
	keyFile := flag.String("key-file", "", "Service account file path(*.json). example: 'path/to/service-account.json'")
	accountFromPile := flag.Bool("stdin", false, "Service account file from stdin. example: cat 'path/to/service-account.json' | gcloud-auth --stdin")

	flag.Parse()

	var serviceAccountFileBytes []byte
	if !*accountFromPile {
		if *keyFile == "" {
			log.Fatal("'-key-file' option is invalid\nexample) gcloud-auth -key-file path/to/service-account.json")
		}

		bytes, err := ioutil.ReadFile(*keyFile)
		if err != nil {
			log.Fatalf("read error from '%v'", *keyFile)
		}
		serviceAccountFileBytes = bytes
	} else {
		// read json from stdio
		if terminal.IsTerminal(0) {
			log.Fatal("'stdin' is invalid\nexample) cat 'path/to/service-account.json | gcloud-auth'")
		}
		serviceAccountFileBytes, _ = ioutil.ReadAll(os.Stdin)
	}

	serviceAccount := &GcloudServiceAccount{}

	if err := json.Unmarshal(serviceAccountFileBytes, serviceAccount); err != nil || serviceAccount.ProjectId == "" {
		log.Fatalf("invalid json: %v", string(serviceAccountFileBytes))
	}

	auth(serviceAccountFileBytes, serviceAccount)
}

func auth(jsonBytes []byte, account *GcloudServiceAccount) {
	jsonFile, err := ioutil.TempFile("", "google_service_account_json")
	if err != nil {
		log.Fatalf("temp file creation failed, %v", err)
	}

	if _, err = jsonFile.Write(jsonBytes); err != nil {
		log.Fatalf("temp file write failed, %v", err)
	}
	defer func() {
		_ = os.Remove(jsonFile.Name())
	}()

	if _, stderr, err := (&shell.Shell{
		Commands: []string{
			"gcloud", "auth",
			"activate-service-account",
			"--key-file", jsonFile.Name(),
			"--project", account.ProjectId,
		},
	}).RunStdout(); err != nil {
		log.Fatalf("Invalid gcloud auth\n%v\n%v", err, stderr)
	}

	fmt.Printf("gcloud account(%v, %v)", account.ProjectId, account.ClientEmail)
}
