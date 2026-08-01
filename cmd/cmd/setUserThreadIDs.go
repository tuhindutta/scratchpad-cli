package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cred struct {
	UserID   string `json:"user_id"`
	ThreadID string `json:"thread_id"`
}

func SetUserThreadIDs(userId string, threadID string, credentialPath string) error {
	bytes, err := json.MarshalIndent(
		Cred{
			UserID:   userId,
			ThreadID: threadID,
		},
		"", "  ",
	)
	if err != nil {
		return fmt.Errorf("Could not marshal the credentials: %v", err)
	}

	file, err := os.Create(credentialPath)
	if err != nil {
		return fmt.Errorf("Could not create credential path: %v", err)
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return fmt.Errorf("Could not write user and thread IDs to the credential file: %v", err)
	}

	return nil

}
