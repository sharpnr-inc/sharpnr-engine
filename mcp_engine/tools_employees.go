package mcp_engine

import (
	"context"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type listEmployeesInput struct{}

type employeeView struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status"`
}

type listEmployeesOutput struct {
	Employees []employeeView `json:"employees" jsonschema:"every employee of the organization"`
	Count     int            `json:"count" jsonschema:"number of employees returned"`
}

// registerEmployeeTools adds org_list_employees, a read over the orchestrator's
// employee directory: the same call the web app's employee list makes.
func registerEmployeeTools(server *mcp.Server, directory EmployeeDirectory) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "org_list_employees",
		Title:       "List employees",
		Description: "Lists the employees of the organization with their role and status. Use it to look up who works here or to resolve a person's id.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: closedWorld},
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ listEmployeesInput) (*mcp.CallToolResult, listEmployeesOutput, error) {
		ctx, p, err := authorize(ctx, req)
		if err != nil {
			return nil, listEmployeesOutput{}, err
		}

		employees, err := directory.GetAllEmployeesOfOrg(ctx, p.OrgID, p.EmployeeID)
		if err != nil {
			slog.Error("mcp tool failed", "tool", "org_list_employees", "org_id", p.OrgID, "error", err.Error())
			return nil, listEmployeesOutput{}, err
		}

		out := listEmployeesOutput{Employees: make([]employeeView, 0, len(employees))}
		for _, e := range employees {
			out.Employees = append(out.Employees, employeeView{
				ID:       e.ID.String(),
				Name:     e.Name,
				Nickname: e.Nickname,
				Role:     e.Role.Role,
				Status:   string(e.Status),
			})
		}
		out.Count = len(out.Employees)

		return nil, out, nil
	})
}
