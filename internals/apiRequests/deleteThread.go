package apirequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeleteThreadPayload struct {
	ThreadID string `json:"thread_id"`
}

func DeleteChatThread(url string, threadID string) error {
	fmt.Println("\n===== DELETE THREAD =====")

	// 1. Prepare payload
	payload := DeleteThreadPayload{ThreadID: threadID}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("unable to marshal payload: %w", err)
	}

	// 2. Perform POST request
	resp, err := http.Post(url+"/assistant/chats/delete/thread", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("unable to perform post request: %w", err)
	}
	defer resp.Body.Close()

	// 3. Verify success status codes (raise_for_status equivalent)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned error status: %s", resp.Status)
	}

	// 4. Read body response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body: %w", err)
	}

	// 5. Unmarshal to map and print (response.json() equivalent)
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(respBytes, &jsonResponse); err != nil {
		return fmt.Errorf("unable to unmarshal response json: %w", err)
	}

	fmt.Printf("%+v\n", jsonResponse)
	return nil
}
