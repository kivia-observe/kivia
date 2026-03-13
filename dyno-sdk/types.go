package dynosdk

import (
	"bytes"
	"net/http"
	"time"
)

type Log struct {

	Id string `json:"id"`
	
	Path string `json:"path"`
	
	Status int `json:"status"`
	
	IPAddress string `json:"ip_address,omitempty"`

	Timestamp time.Time `json:"timestamp"`
	
	Latency int `json:"latency"`
}

type Client struct {
	BaseURL    string
	ApiKey     string
	HttpClient *http.Client
}
type ResponseWriteWrapper struct {
	http.ResponseWriter
	body *bytes.Buffer
	statusCode int
}