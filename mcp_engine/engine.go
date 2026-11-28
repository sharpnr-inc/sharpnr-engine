package mcp_engine

import (
	"context"
	"log/slog"
	"net/http"

	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// EmployeeDirectory is the slice of the orchestrator the employee tools use.
// Kept narrow so tests can fake it; the orchestrator satisfies it as is.
type EmployeeDirectory interface {
	GetAllEmployeesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]orchestrator_dtos.EmployeeOfOrganization, error)
}

// Deps are the services the tools call. Every field must be set.
type Deps struct {
	Employees EmployeeDirectory
}

// Config configures the HTTP side of the server.
type Config struct {
	// Version is reported to clients during initialization.
	Version string

	// Verifier turns a bearer token into a Principal. Required.
	Verifier auth.TokenVerifier
}

// serverInstructions is sent to clients at initialization; models read it to
// decide when and how to use these tools.
const serverInstructions = `Sharpnr business data for the signed-in employee's organization.
Every tool acts as that employee and returns only what they are allowed to see.
Tools are read-only.`

// NewServer builds the MCP server with every tool registered.
func NewServer(version string, deps Deps) *mcp.Server {
	server := mcp.NewServer(
		&mcp.Implementation{Name: "sharpnr-engine", Title: "Sharpnr", Version: version},
		&mcp.ServerOptions{
			Instructions: serverInstructions,
			Logger:       slog.Default(),
		},
	)

	registerIdentityTools(server)
	registerEmployeeTools(server, deps.Employees)

	return server
}

// NewHandler returns the http.Handler to mount at /mcp: the Streamable HTTP
// transport behind bearer-token auth.
//
// Stateless: every request carries its own token and gets a fresh session, so
// any instance of the engine can serve any request and nothing is kept in
// memory between calls. That rules out server-to-client requests (sampling,
// elicitation), which no tool needs yet.
func NewHandler(cfg Config, deps Deps) http.Handler {
	if cfg.Verifier == nil {
		panic("mcp_engine: Config.Verifier is required")
	}
	if deps.Employees == nil {
		panic("mcp_engine: Deps.Employees is required")
	}

	server := NewServer(cfg.Version, deps)

	transport := mcp.NewStreamableHTTPHandler(
		func(*http.Request) *mcp.Server { return server },
		&mcp.StreamableHTTPOptions{
			Stateless: true,
			Logger:    slog.Default(),
		},
	)

	return auth.RequireBearerToken(cfg.Verifier, nil)(transport)
}

// closedWorld marks a tool as touching only this org's data, never the open
// internet (the MCP openWorldHint annotation, which defaults to true).
var closedWorld = func() *bool { b := false; return &b }()
