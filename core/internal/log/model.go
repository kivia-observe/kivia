package log

import "time"

type Log struct {
	
	Id string `json:"id"`
	
	Path string `json:"path"`
	
	Status int `json:"status"`
	
	IPAddress string `json:"ip_address,omitempty"`

	Timestamp time.Time `json:"timestamp"`
	
	Latency string `json:"latency"`
	
	ApiKey string `json:"api_key"`
	
}