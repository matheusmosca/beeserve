package beeserve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BeeServeClient struct {
	appName         string
	host            string
	port            int
	url             string
	metricsEndpoint string
}

func NewClient(appName string, host string, port int) BeeServeClient {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return BeeServeClient{
		appName:         appName,
		host:            host,
		port:            port,
		url:             url,
		metricsEndpoint: fmt.Sprintf("%s/v1/metrics", url),
	}
}

func (c BeeServeClient) PushMetrics(metricsData Metrics) {
	b, err := json.Marshal(metricsData)
	if err != nil {
		logrus.WithError(err).Error("could not marshal metrics data")
		return
	}
	reader := bytes.NewReader(b)

	req, err := http.NewRequest(http.MethodPost, c.metricsEndpoint, reader)
	if err != nil {
		logrus.WithError(err).Error("could not create the request to metrics")
		return
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logrus.WithError(err).Error("could not send data to /metrics")
	}
}
