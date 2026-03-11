package dynosdk

import (
	"bytes"
	"net/http"
)

type ResponseWriteWrapper struct {
	http.ResponseWriter
	body *bytes.Buffer
	statusCode int
}