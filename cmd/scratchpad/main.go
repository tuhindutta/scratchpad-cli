package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tuhindutta/scratchpad-cli/internals/cli"
	cliCreds "github.com/tuhindutta/scratchpad-cli/internals/cli/creds"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentialPath := os.Getenv("CREDENTIALS_PATH")
	userId, threadId := cliCreds.ReadUserThreadIDs(credentialPath)

	app := cli.App{
		Url:            "http://0.0.0.0",
		Port:           8081,
		UserId:         userId,
		ThreadId:       threadId,
		CredentialPath: credentialPath,
	}

	app.Execute()
}
