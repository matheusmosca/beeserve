package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusmosca/beeserve"
	"github.com/sirupsen/logrus"
)

const (
	port = 8000
	host = "localhost"
)

func main() {
	address := fmt.Sprintf("%s:%d", host, port)

	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/metrics", helloHandler).Methods(http.MethodPost)

	http.ListenAndServe(address, v1)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var metrics beeserve.Metrics
	json.NewDecoder(r.Body).Decode(&metrics)
	loggger := logrus.New()
	log := logrus.NewEntry(loggger).WithFields(logrus.Fields{
		"endpoint":    metrics.Endpoint,
		"app_name":    metrics.AppName,
		"status_code": metrics.StatusCode,
		"body":        metrics.Body,
		"duration":    metrics.Duration,
		"start_time":  metrics.StartTime,
		"end_time":    metrics.EndTime,
	})
	log.Infoln("")
	w.WriteHeader(http.StatusOK)
	return
}
