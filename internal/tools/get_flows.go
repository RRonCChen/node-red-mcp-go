package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	get_flows_url_path = "/flows"
)

type GetFlowsTool struct {
	nodeRedHost string
	client      resty.Client
}

func NewGetFlowsTool(nodeRedHost string) *GetFlowsTool {
	return &GetFlowsTool{
		client:      *resty.New(),
		nodeRedHost: nodeRedHost,
	}
}

func (f *GetFlowsTool) GetTool() mcp.Tool {
	return mcp.NewTool("get_flows", mcp.WithDescription("Get the active flow configuration."))
}

func (f *GetFlowsTool) ToolHandlerFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	response, err := f.client.R().Get(fmt.Sprintf("http://%s/%s", f.nodeRedHost, get_flows_url_path))
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("get flows api return status code : %d, response : %s", response.StatusCode(), string(response.Body()))
	}

	return mcp.NewToolResultText(string(response.Body())), nil
}
