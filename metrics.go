package beeserve

import "time"

type Metrics struct {
	Endpoint   string        `json:"endpoint"`
	AppName    string        `json:"app_name"`
	Headers    interface{}   `json:"headers"`
	StatusCode int           `json:"status_code"`
	Body       interface{}   `json:"body"`
	StartTime  time.Time     `json:"start_time"`
	EndTime    time.Time     `json:"end_time"`
	Duration   time.Duration `json:"duration"`
}
