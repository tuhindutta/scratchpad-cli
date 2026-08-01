package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tuhindutta/scratchpad-cli/cmd/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentialPath := os.Getenv("CREDENTIALS_PATH")
	var userId string
	var threadId string

	data, err := os.ReadFile(credentialPath)
	if err != nil {
		log.Printf("Error credential reading file: %v \nCreating credentials.", err)
		userId = "temporary_user001"
		threadId = "temporary_thread001"
	} else {
		var cred cli.Cred
		err = json.Unmarshal(data, &cred)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		userId = cred.UserID
		threadId = cred.ThreadID
	}

	fmt.Println(userId, "  ", threadId)

	app := cli.App{
		Url:            "http://0.0.0.0",
		Port:           8081,
		UserId:         userId,
		ThreadId:       threadId,
		CredentialPath: credentialPath,
	}

	app.Execute()
}

// prompts, errs := apirequests.Assistant("http://localhost:8081", "u1", "t1", "Hello!")

// for prompt := range prompts {
// 	fmt.Print(prompt)
// }

// if err := <-errs; err != nil {
// 	log.Fatalf("Assistant stream failed: %v", err)
// }
