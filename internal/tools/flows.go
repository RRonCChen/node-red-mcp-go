package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RRonCChen/node-red-mcp-go/internal/client"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	flows_path       = "/flows"
	flows_state_path = "/flows/state"
	flows_by_id_path = "/flows/%s"
)

func NewGetFlowTool() mcp.Tool {
	return mcp.NewTool("get_flows", mcp.WithDescription("Get the active flow configuration."))
}

// Implemment get_flow handler func.
// Get the active flow configuration by calling nodered '/flows' api.
// https://nodered.org/docs/api/admin/methods/get/flowsㄥ.
func GetFlowsHandlerFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return CallNodeRed(http.MethodGet, flows_path, nil)
}

func NewGetFlowsStateTool() mcp.Tool {
	return mcp.NewTool("get_flows_state", mcp.WithDescription("Get the active flow’s runtime state"))
}

// Implemment get_flows_state handler func.
// Get the active flow’s runtime state by calling nodered '/flows/states' api.
// https://nodered.org/docs/api/admin/methods/get/flows/state/.
func GetFlowsStateHandlerFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return CallNodeRed(http.MethodGet, flows_state_path, nil)
}

func NewGetFlowByIdTool() mcp.Tool {
	return mcp.NewTool("get_flow_by_id",
		mcp.WithDescription("Get an individual flow configuration by flow id"),
		mcp.WithString("flowId",
			mcp.Required(),
			mcp.Description("flow id")))
}

// Implemment get_flow_by_id handler func.
// Get an individual flow configuration by flow id calling nodered '/flows/states/:id' api.
// https://nodered.org/docs/api/admin/methods/get/flows/state/.
func GetFlowsByIdHandlerFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	flowId, err := request.RequireString("flowId")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	path := fmt.Sprintf(flows_by_id_path, flowId)

	return CallNodeRed(http.MethodGet, path, nil)
}

func CallNodeRed(method, path string, body any) (*mcp.CallToolResult, error) {
	noderedCli := client.NewNodeRedClient()

	resp, err := noderedCli.Call(method, path, body)
	if err != nil {
		return mcp.NewToolResultError("call nodered api error occurred"), nil
	}

	return mcp.NewToolResultText(string(resp.Body())), nil
}
