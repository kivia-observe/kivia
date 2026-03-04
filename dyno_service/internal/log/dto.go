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