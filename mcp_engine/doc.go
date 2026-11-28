// Package mcp_engine is the app's Model Context Protocol server: it lets AI
// agents (Claude, IDEs, other MCP clients) work with an organization's data on
// behalf of a signed-in employee.
//
// The rules every tool in this package follows:
//
//  1. An agent is the employee, never more. Every token resolves to an
//     (org, employee) pair and tools call the same orchestrator and domain
//     services the web app calls, so Sentinel decides exactly as it would for
//     that person. No tool queries a repository directly.
//  2. Requests are stamped with channel "mcp" (sentinelContexts.ChannelMCP),
//     so policies can tell an agent from a person in the app.
//  3. Tools are task-level ("list the employees of my org"), not one per HTTP
//     route, and read-only until writes get a prepare/commit flow.
//
// The package takes net/http types and knows nothing about gin; mounting it is
// mcp_engine/routes' job, the same split as ws_engine.
//
// Auth is temporary: tokens are the engine's own mobile access tokens
// (see AccessTokenVerifier). Once sharp-account issues tokens with an org
// claim and a per-API audience, swap the verifier; nothing else changes.
package mcp_engine
