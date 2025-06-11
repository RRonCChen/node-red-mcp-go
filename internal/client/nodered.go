package client

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/resty.v1"
)

const (
	content_type     = "Content-Type"
	application_json = "application/json"
)

// Environment variable name for NodeRed URL
const EnvNodeRedHost = "NODERED_HOST"

// Default NodeRed URL when environment variable is not set
const DefaultNodeRedHost = "localhost:1880"

type NodeRedClient struct {
	client *resty.Client
}

func NewNodeRedClient() *NodeRedClient {
	return &NodeRedClient{client: resty.New()}
}

func (n *NodeRedClient) Call(method string, path string, body any) (resp *resty.Response, err error) {
	nodeRedURL := os.Getenv(EnvNodeRedHost)
	if nodeRedURL == "" {
		nodeRedURL = DefaultNodeRedHost
	}

	url := fmt.Sprintf("http://%s/%s", nodeRedURL, path)

	switch method {
	case http.MethodGet:
		return n.client.R().Get(url)
	case http.MethodPost:
		return n.client.R().SetHeader(content_type, application_json).SetBody(&body).Post(url)
	case http.MethodPut:
		return n.client.R().SetHeader(content_type, application_json).SetBody(&body).Put(url)
	case http.MethodDelete:
		return n.client.R().Delete(url)
	default:
		return nil, fmt.Errorf("method not support")
	}
}
