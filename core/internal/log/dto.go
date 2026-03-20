package log

import "time"

type createLogRequest struct {
	Id string `json:"id"`

	Path string `json:"path"`

	Status int `json:"status"`

	IPAddress string `json:"ip_address,omitempty"`

	Timestamp time.Time `json:"timestamp"`

	Latency int `json:"latency"`
}

type PaginatedLogResponse struct {
	Logs []LogResponse `json:"logs"`

	Page int `json:"page"`

	Items int `json:"items"`

	TotelItems int `json:"total_items"`
}

type LogResponse struct {
	Id        string    `json:"id"`
	Path      string    `json:"path"`
	Latency   string    `json:"latency"`
	Status    int       `json:"status"`
	IPAddress string    `json:"ip_address,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	ApiKey    string    `json:"api_key"`
}
