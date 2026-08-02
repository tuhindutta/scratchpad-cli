package cliCreds

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func GetFreePort() (int, error) {
	// Resolve a local TCP address on port 0
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	// Open a temporary listener
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listener.Close() // Free it immediately so your server can use it

	// Return the assigned port number
	return listener.Addr().(*net.TCPAddr).Port, nil
}

type Cred struct {
	UserID   string `json:"user_id"`
	ThreadID string `json:"thread_id"`
	Port     int    `json:"port"`
}

func SetUserThreadIDsPort(userId string, threadID string, credentialPath string) error {

	defaultPort := 8080
	port, err := GetFreePort()
	if err != nil {
		log.Printf("Could not get a free port: %v \nDrfaulting to %d", err, defaultPort)
		port = defaultPort
	}

	bytes, err := json.MarshalIndent(
		Cred{
			UserID:   userId,
			ThreadID: threadID,
			Port:     port,
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
