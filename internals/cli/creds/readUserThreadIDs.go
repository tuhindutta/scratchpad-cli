package cliCreds

import (
	"encoding/json"
	"log"
	"os"
)

func ReadUserThreadIDsPort(credentialPath string) (string, string, int) {

	var userId string
	var threadId string
	var port int

	data, err := os.ReadFile(credentialPath)
	if err != nil {
		log.Printf(
			"[WARN] Error credential reading file: %v\n"+
				"Note: This will be a volatile session and all conversation records will be forgotten once the session is closed if you have not set the credentials first.\n"+
				"To avoid this, please create credentials first using `cred` command.",
			err,
		)

		userId = "temporary_user001"
		threadId = "temporary_thread001"
		port = 8080
	} else {
		var cred Cred
		err = json.Unmarshal(data, &cred)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
		userId = cred.UserID
		threadId = cred.ThreadID
		port = cred.Port
	}

	return userId, threadId, port

}
