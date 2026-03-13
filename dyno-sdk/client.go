package dynosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func NewClient(apiKey string) *Client {
	client := &Client{
		BaseURL:    "http://localhost:80/api",
		ApiKey:     apiKey,
		HttpClient: &http.Client{},
	}

	return client
}

func (c *Client) NewLog(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// start time for the request
		start := time.Now()
		
		rw := &ResponseWriteWrapper{
			ResponseWriter: w,
			body:           &bytes.Buffer{},
			statusCode:     200,
		}

		next.ServeHTTP(rw, r)

		latency := time.Since(start)

		go c.sendLog(r, rw, latency)
	}
}

func (c *Client) sendLog(r *http.Request, rw *ResponseWriteWrapper, latency time.Duration) {
    url := fmt.Sprintf("%s/logs/create", c.BaseURL)
    logEntry := &Log{
        Path:      r.URL.String(),
        Status:    rw.statusCode,
        IPAddress: r.RemoteAddr,
        Timestamp: time.Now(),
        Latency:   int(latency.Milliseconds()),
    }

    bodyBytes, err := json.Marshal(logEntry)
    if err != nil {
        log.Printf("dynosdk: failed to marshal log: %v", err)
        return
    }

    req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
    if err != nil {
        log.Printf("dynosdk: failed to create request: %v", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-dyno-api-key", c.ApiKey)

    res, err := c.HttpClient.Do(req)
    if err != nil {
        log.Printf("dynosdk: failed to send log: %v", err)
        return
    }
    defer res.Body.Close()

    if res.StatusCode != 201 {
        body, _ := io.ReadAll(res.Body)
        log.Printf("dynosdk: log rejected (%d): %s", res.StatusCode, string(body))
    }
}