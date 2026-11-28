package mcp_engine

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	coreEmployeeConstants "sharpnr-engine/core/employees/constants"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// These tests drive the real handler over HTTP with the SDK's own client, so
// they cover the transport, the bearer-token middleware and the tools together.
// Only the token verifier and the orchestrator are faked.

const goodToken = "good-token"

var testPrincipal = Principal{
	EmployeeID: uuid.New(),
	OrgID:      uuid.New(),
	IPAddress:  "203.0.113.7",
}

// fakeVerifier accepts goodToken and maps it to testPrincipal.
func fakeVerifier(_ context.Context, token string, _ *http.Request) (*auth.TokenInfo, error) {
	if token != goodToken {
		return nil, auth.ErrInvalidToken
	}
	return newTokenInfo(testPrincipal, time.Now().Add(time.Hour)), nil
}

// fakeDirectory records what the tool passed to the orchestrator.
type fakeDirectory struct {
	employees []orchestrator_dtos.EmployeeOfOrganization
	err       error

	gotOrgID, gotEmplID uuid.UUID
	gotAuthz            sentinelContexts.AuthzContext
}

func (f *fakeDirectory) GetAllEmployeesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]orchestrator_dtos.EmployeeOfOrganization, error) {
	f.gotOrgID, f.gotEmplID = orgID, emplID
	f.gotAuthz, _ = sentinelContexts.Extract(ctx)
	return f.employees, f.err
}

func startServer(t *testing.T, dir *fakeDirectory) string {
	t.Helper()
	srv := httptest.NewServer(NewHandler(Config{Version: "test", Verifier: fakeVerifier}, Deps{Employees: dir}))
	t.Cleanup(srv.Close)
	return srv.URL
}

// bearerTransport adds the Authorization header the way an MCP client
// configured with a static token does.
type bearerTransport struct{ token string }

func (b bearerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	r.Header.Set("Authorization", "Bearer "+b.token)
	return http.DefaultTransport.RoundTrip(r)
}

func connect(t *testing.T, url, token string) (*mcp.ClientSession, error) {
	t.Helper()
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "v0"}, nil)
	session, err := client.Connect(t.Context(), &mcp.StreamableClientTransport{
		Endpoint:             url,
		HTTPClient:           &http.Client{Transport: bearerTransport{token: token}},
		MaxRetries:           -1,
		DisableStandaloneSSE: true,
	}, nil)
	if err == nil {
		t.Cleanup(func() { _ = session.Close() })
	}
	return session, err
}

// decode turns a tool's structured output back into its Go type.
func decode[T any](t *testing.T, res *mcp.CallToolResult) T {
	t.Helper()
	require.False(t, res.IsError, "tool returned an error: %v", res.Content)
	raw, err := json.Marshal(res.StructuredContent)
	require.NoError(t, err)
	var out T
	require.NoError(t, json.Unmarshal(raw, &out))
	return out
}

/* ---------------------------------- <***> --------------------------------- */

func TestRejectsRequestsWithoutAValidToken(t *testing.T) {
	url := startServer(t, &fakeDirectory{})

	t.Run("no Authorization header is a 401", func(t *testing.T) {
		resp, err := http.Post(url, "application/json", strings.NewReader(`{}`))
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("a client with a bad token cannot connect", func(t *testing.T) {
		_, err := connect(t, url, "forged")
		require.Error(t, err)
	})
}

func TestListsOnlyReadOnlyTools(t *testing.T) {
	session, err := connect(t, startServer(t, &fakeDirectory{}), goodToken)
	require.NoError(t, err)

	res, err := session.ListTools(t.Context(), nil)
	require.NoError(t, err)

	names := make([]string, 0, len(res.Tools))
	for _, tool := range res.Tools {
		names = append(names, tool.Name)
		require.NotNil(t, tool.Annotations, tool.Name)
		assert.True(t, tool.Annotations.ReadOnlyHint, "%s must be read-only", tool.Name)
	}
	assert.ElementsMatch(t, []string{"sharpnr_whoami", "org_list_employees"}, names)
}

func TestWhoamiReturnsTheTokensPrincipal(t *testing.T) {
	session, err := connect(t, startServer(t, &fakeDirectory{}), goodToken)
	require.NoError(t, err)

	res, err := session.CallTool(t.Context(), &mcp.CallToolParams{Name: "sharpnr_whoami"})
	require.NoError(t, err)

	out := decode[whoamiOutput](t, res)
	assert.Equal(t, testPrincipal.EmployeeID.String(), out.EmployeeID)
	assert.Equal(t, testPrincipal.OrgID.String(), out.OrgID)
	assert.Equal(t, sentinelContexts.ChannelMCP, out.Channel)
}

func TestListEmployeesActsAsTheCaller(t *testing.T) {
	dir := &fakeDirectory{employees: []orchestrator_dtos.EmployeeOfOrganization{
		{ID: uuid.New(), Name: "Asha Rao", Role: orchestrator_dtos.EmployeeRole{Role: "Manager"}, Status: coreEmployeeConstants.EmployeeStatusActive},
		{ID: uuid.New(), Name: "Ravi Kumar", Nickname: "Ravi", Status: coreEmployeeConstants.EmployeeStatusActive},
	}}
	session, err := connect(t, startServer(t, dir), goodToken)
	require.NoError(t, err)

	res, err := session.CallTool(t.Context(), &mcp.CallToolParams{Name: "org_list_employees"})
	require.NoError(t, err)

	out := decode[listEmployeesOutput](t, res)
	assert.Equal(t, 2, out.Count)
	assert.Equal(t, "Asha Rao", out.Employees[0].Name)
	assert.Equal(t, "Manager", out.Employees[0].Role)

	// The orchestrator was asked as the token's org and employee, with the
	// Sentinel facts stamped as an MCP request.
	assert.Equal(t, testPrincipal.OrgID, dir.gotOrgID)
	assert.Equal(t, testPrincipal.EmployeeID, dir.gotEmplID)
	assert.Equal(t, sentinelContexts.ChannelMCP, dir.gotAuthz.Request.Channel)
	assert.Equal(t, testPrincipal.IPAddress, dir.gotAuthz.Request.IPAddress)
}

func TestServiceErrorsComeBackAsToolErrors(t *testing.T) {
	dir := &fakeDirectory{err: errors.New("access denied")}
	session, err := connect(t, startServer(t, dir), goodToken)
	require.NoError(t, err)

	res, err := session.CallTool(t.Context(), &mcp.CallToolParams{Name: "org_list_employees"})
	require.NoError(t, err, "a failing tool is a result, not a protocol error")
	assert.True(t, res.IsError)
}

func TestPrincipalFromRejectsRequestsWithoutTokenInfo(t *testing.T) {
	_, err := principalFrom(&mcp.CallToolRequest{})
	assert.ErrorIs(t, err, ErrUnauthenticated)
}
