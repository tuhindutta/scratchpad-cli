package apirequests

import (
	"bufio"
	"fmt"
	"net/http"
)

// ListThreads streams lines from the server via a channel, yielding values like a Python generator.
func ListThreads(url string) (<-chan string, <-chan error) {
	out := make(chan string)
	errChan := make(chan error, 1) // Buffered to prevent leaking goroutine on exit

	go func() {
		// Ensure channels close when the stream finishes or encounters an error
		defer close(out)
		defer close(errChan)

		req, err := http.NewRequest(http.MethodGet, url+"/assistant/chats/list", nil)
		if err != nil {
			errChan <- fmt.Errorf("Unable to create http request: %w", err)
			return
		}
		// Expecting JSON payload back from the server
		req.Header.Set("Accept", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- fmt.Errorf("Unable to do request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			errChan <- fmt.Errorf("Server returned %s", resp.Status)
			return
		}

		scanner := bufio.NewScanner(resp.Body)
		buff := make([]byte, 0, 64*1024)
		scanner.Buffer(buff, 10*1024*1024)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			// "Yield" the line to the consumer
			out <- line
		}

		if err := scanner.Err(); err != nil {
			errChan <- fmt.Errorf("scanner error: %w", err)
		}
	}()

	return out, errChan
}
