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

	BaseURL string
	ApiKey string
	HttpClient *http.Client
	Request *http.Request
	Response *http.Response
}

type Log struct {

	Id string `json:"id"`
	
	Path string `json:"path"`
	
	Status int `json:"status"`
	
	IPAddress string `json:"ip_address,omitempty"`

	Timestamp time.Time `json:"timestamp"`
	
	Latency int `json:"latency"`
}

func NewClient(apiKey string, request *http.Request, response *http.Response) *Client {
	client := &Client{
		BaseURL: "http://localhost:8080/api",
		ApiKey: apiKey,
		HttpClient: &http.Client{},
		Request: request,
		Response: response,
	}

	return client
}


func (c *Client) NewLog(latency time.Duration) error {

	fmt.Println("latency", latency.Milliseconds())

	url := fmt.Sprintf("%s/logs/create", c.BaseURL)

	log := &Log{
		Path: c.Request.URL.String(),
		Status: c.Response.StatusCode,
		IPAddress: getPublicIP(c.HttpClient),
		Timestamp: time.Now(),
		Latency: int(latency.Milliseconds()),
	}

	bodyBytes, err := json.Marshal(log)

	if err != nil {
		return err
	}

	bytesReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(http.MethodPost, url, bytesReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-dyno-api-key", c.ApiKey)

	if err != nil {
		return err
	}

	res, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		body, _ := io.ReadAll(res.Body)
		logger.Println(string(body))
	}

	logger.Println("Logged successfully", log, res)

	return nil
}