package apirequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Shutdown(url string) error {
	fmt.Println("\n===== SHUTDOWN =====")

	// 2. Perform POST request
	resp, err := http.Post(url+"/shutdown", "application/json", bytes.NewBuffer(nil))
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

// ListThreads streams lines from the server via a channel, yielding values like a Python generator.
// func Shutdown(url string) (<-chan string, <-chan error) {
// 	out := make(chan string)
// 	errChan := make(chan error, 1) // Buffered to prevent leaking goroutine on exit

// 	go func() {
// 		// Ensure channels close when the stream finishes or encounters an error
// 		defer close(out)
// 		defer close(errChan)

// 		req, err := http.NewRequest(http.MethodPost, url+"/shutdown", nil)
// 		if err != nil {
// 			errChan <- fmt.Errorf("Unable to create http request: %w", err)
// 			return
// 		}
// 		// Expecting JSON payload back from the server
// 		req.Header.Set("Accept", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			errChan <- fmt.Errorf("Unable to do request: %w", err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 			errChan <- fmt.Errorf("Server returned %s", resp.Status)
// 			return
// 		}

// 		scanner := bufio.NewScanner(resp.Body)
// 		buff := make([]byte, 0, 64*1024)
// 		scanner.Buffer(buff, 10*1024*1024)

// 		for scanner.Scan() {
// 			line := scanner.Text()
// 			if line == "" {
// 				continue
// 			}
// 			// "Yield" the line to the consumer
// 			out <- line
// 		}

// 		if err := scanner.Err(); err != nil {
// 			errChan <- fmt.Errorf("scanner error: %w", err)
// 		}
// 	}()

// 	return out, errChan
// }
