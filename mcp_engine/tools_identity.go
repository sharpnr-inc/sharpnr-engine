package mcp_engine

import (
	"context"

	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type whoamiInput struct{}

type whoamiOutput struct {
	EmployeeID string `json:"employee_id" jsonschema:"id of the employee the agent acts for"`
	OrgID      string `json:"org_id" jsonschema:"id of the organization the agent acts in"`
	Channel    string `json:"channel" jsonschema:"how the request reached Sharpnr; always mcp here"`
}

// registerIdentityTools adds sharpnr_whoami: the first call to make when wiring
// up a client, since it proves the token resolves to the right org and person
// without touching any business data.
func registerIdentityTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "sharpnr_whoami",
		Title:       "Who am I",
		Description: "Returns the employee and organization this session acts for. Call it to confirm which org the other tools will read from.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: closedWorld},
	}, whoami)
}

func whoami(ctx context.Context, req *mcp.CallToolRequest, _ whoamiInput) (*mcp.CallToolResult, whoamiOutput, error) {
	_, p, err := authorize(ctx, req)
	if err != nil {
		return nil, whoamiOutput{}, err
	}

	return nil, whoamiOutput{
		EmployeeID: p.EmployeeID.String(),
		OrgID:      p.OrgID.String(),
		Channel:    sentinelContexts.ChannelMCP,
	}, nil
}
