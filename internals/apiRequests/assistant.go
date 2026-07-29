package apirequests

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Credential struct {
	UserId   string `json:"user_id"`
	ThreadID string `json:"thread_id"`
}

type AssistantPayload struct {
	Message    string     `json:"message"`
	Credential Credential `json:"credential"`
}

type AssistantEvent struct {
	Node     string `json:"node"`
	Type     string `json:"type"`
	ToolName string `json:"tool_name"`
	Content  string `json:"content"`
}

func trimContent(content string, length int) string {
	if len(content) > int(1.1*float64(length)) {
		return content[:length] + "......"
	}
	return content
}

// Assistant processes response streams and yields formatted strings line-by-line
func Assistant(url string, userId string, threadId string, message string) (<-chan string, <-chan error) {
	out := make(chan string)
	errChan := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errChan)

		payload := AssistantPayload{
			Message:    message,
			Credential: Credential{UserId: userId, ThreadID: threadId},
		}
		body, err := json.Marshal(payload)
		if err != nil {
			errChan <- fmt.Errorf("unable to marshal payload: %w", err)
			return
		}

		// GET request with a payload body matching the Python httpx implementation
		req, err := http.NewRequest(http.MethodGet, url+"/assistant/chat", bytes.NewBuffer(body))
		if err != nil {
			errChan <- fmt.Errorf("unable to create request: %w", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errChan <- fmt.Errorf("unable to execute request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			errChan <- fmt.Errorf("server returned error status: %s", resp.Status)
			return
		}

		fmt.Println("\n===== ASSISTANT =====")
		var currentNode string

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var event AssistantEvent
			if err := json.Unmarshal([]byte(line), &event); err != nil {
				continue // Safely skip corrupted lines
			}

			content := event.Content
			if event.Type == "tool" {
				content = trimContent(content, 200)
			}

			// Format headers and nodes if a transition happens
			var prefix string
			if event.Node != currentNode {
				currentNode = event.Node
				prefix = fmt.Sprintf("\n\n[%s]", currentNode)
			}

			prompt := fmt.Sprintf("%s\nNode: %s\n\nType: %s\n\nTool Name: %s\n\nContent:\n%s\n\n",
				prefix, event.Node, event.Type, event.ToolName, content)

			// "Yield" the fully compiled text block to the reader
			out <- prompt
		}

		errChan <- scanner.Err()
	}()

	return out, errChan
}
