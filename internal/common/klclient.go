package cm

import (
	"crypto/tls"
	"net/http"
	"os"
	"strconv"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client http.Client
}

func NewHttpClient() HttpClient {
	return &httpClient{
		client: initClient(),
	}
}

func (i *httpClient) Do(req *http.Request) (*http.Response, error) {
	return i.client.Do(req)
}

func initClient() http.Client {
	timeout, _ := strconv.Atoi(os.Getenv("HTTP_CLIENT_TIMEOUT_IN_SEC"))
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: customTransport,
	}
	return client
}
