package dynosdk

import (
	"bytes"
	"net/http"
)
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