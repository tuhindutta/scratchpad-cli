package apirequests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DeleteFullChatHistory(url string) error {
	fmt.Println("\n===== DELETE FULL CHAT HISTORY =====")

	// 1. Perform POST request with nil body payload
	resp, err := http.Post(url+"/assistant/chats/delete/full", "application/json", nil)
	if err != nil {
		return fmt.Errorf("unable to perform post request: %w", err)
	}
	defer resp.Body.Close()

	// 2. Verify response status code (raise_for_status equivalent)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned error status: %s", resp.Status)
	}

	// 3. Read body data stream
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body: %w", err)
	}

	// 4. Unmarshal and print (response.json() equivalent)
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(respBytes, &jsonResponse); err != nil {
		return fmt.Errorf("unable to unmarshal response json: %w", err)
	}

	fmt.Printf("%+v\n", jsonResponse)
	return nil
}
