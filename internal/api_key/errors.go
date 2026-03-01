package apikey

import "errors"

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrDuplicateKey = errors.New("Duplicate API key name for the project")
	ErrApiKeyNotFound = errors.New("API key not found")
	ErrInvalidApiKey = errors.New("Invalid API key")
)