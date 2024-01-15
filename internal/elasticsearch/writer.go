package elasticsearch

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type ESWriter struct {
	URL string // Elasticsearch endpoint URL
}

func NewESWriter(URL string) (*ESWriter, error) {
	if URL == "" {
		return nil, errors.New("url cannot be empty")
	}
	return &ESWriter{URL: URL}, nil
}

// Write satisfies the io.Writer interface and sends data to Elasticsearch.
func (w ESWriter) Write(p []byte) (n int, err error) {

	u := fmt.Sprintf("%s/%s", w.URL, "logs/_doc")
	// Create a new HTTP POST request with the log data.
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the request using the default client.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check if the request was successful.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to send log, status code: %d", resp.StatusCode)
	}

	return len(p), nil
}
