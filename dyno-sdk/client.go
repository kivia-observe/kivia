package dynosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	logger "log"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	ApiKey     string
	HttpClient *http.Client
}

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

		url := fmt.Sprintf("%s/logs/create", c.BaseURL)

		log := &Log{
			Path:      r.URL.String(),
			Status:    rw.statusCode,
			IPAddress: r.RemoteAddr,
			Timestamp: time.Now(),
			Latency: int(latency.Milliseconds()),
		}

		bodyBytes, err := json.Marshal(log)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		bytesReader := bytes.NewReader(bodyBytes)

		req, err := http.NewRequest(http.MethodPost, url, bytesReader)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-dyno-api-key", c.ApiKey)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := c.HttpClient.Do(req)

		if err != nil {
			http.Error(rw, err.Error(), res.StatusCode)
			return
		}

		if res.StatusCode != 201 {
			body, _ := io.ReadAll(res.Body)
			logger.Println(string(body))
		}

		logger.Println("Logged successfully", rw.statusCode)
	}
}
