package utils

import "errors"

var (

	// api key related errors
	ErrUnauthorized = errors.New("Unauthorized")
	ErrDuplicateKey = errors.New("Duplicate API key name for the project")
	ErrApiKeyNotFound = errors.New("API key not found")
	ErrInvalidApiKey = errors.New("Invalid API key")

	// project related errors
	ErrProjectNotFound = errors.New("Project not found")
)