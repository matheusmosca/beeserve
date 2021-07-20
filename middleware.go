package beeserve

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (r *ResponseWriter) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func WithMetrics(h http.Handler, client BeeServeClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var metrics Metrics
		metrics.StartTime = time.Now()
		rw := &ResponseWriter{
			ResponseWriter: w,
			Status:         200,
		}

		json.NewDecoder(r.Body).Decode(&metrics.Body)
		h.ServeHTTP(rw, r)

		metrics.EndTime = time.Now()
		metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)

		json.NewDecoder(r.Body).Decode(&metrics.Body)
		metrics.Endpoint = r.RequestURI
		metrics.Headers = r.Header
		metrics.StatusCode = rw.Status
		metrics.AppName = client.appName
		client.PushMetrics(metrics)
	})
}
