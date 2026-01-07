package sere

import "context"

type Workflow struct {
	Ctx  context.Context
	Name EventName

	// shared mutable state
	Data map[string]any
}

type WorkflowHandler func(*Workflow) error
