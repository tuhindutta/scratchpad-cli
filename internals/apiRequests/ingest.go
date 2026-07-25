package apirequests

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IngestPayload struct {
	UserId   string `json:"user_id"`
	ThreadID string `json:"thread_id"`
}

func Ingest(url string, user_id string, thread_id string) error {

	payload := IngestPayload{
		UserId:   user_id,
		ThreadID: thread_id,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Unable to marshal the payload: %s", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("Unable to create http request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Unable to do request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Server returned %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	buff := make([]byte, 0, 64*1024)
	scanner.Buffer(buff, 10*1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println(line)
	}

	return scanner.Err()
}
