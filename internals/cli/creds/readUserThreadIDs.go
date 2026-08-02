package cliCreds

import (
	"encoding/json"
	"log"
	"os"
)

func ReadUserThreadIDs(credentialPath string) (string, string) {

	var userId string
	var threadId string

	data, err := os.ReadFile(credentialPath)
	if err != nil {
		log.Printf("Error credential reading file: %v \nCreating credentials.", err)
		userId = "temporary_user001"
		threadId = "temporary_thread001"
	} else {
		var cred Cred
		err = json.Unmarshal(data, &cred)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		userId = cred.UserID
		threadId = cred.ThreadID
	}

	return userId, threadId

}
