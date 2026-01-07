// sere/workflow_exec.go

package sere

import "context"

func (s *SERE) SubscribeWorkflow(
	name EventName,
	handler WorkflowHandler,
) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.workflowSubscribers[string(name)] =
		append(s.workflowSubscribers[string(name)], handler)
}

// Execute runs a workflow synchronously
func (s *SERE) Execute(
	ctx context.Context,
	name EventName,
	data map[string]any,
) (*Workflow, error) {

	wf := &Workflow{
		Ctx:  ctx,
		Name: name,
		Data: data,
	}

	s.mu.RLock()
	handlers := s.workflowSubscribers[string(name)]
	s.mu.RUnlock()

	for _, h := range handlers {
		if err := h(wf); err != nil {
			return nil, err
		}
	}

	return wf, nil
}
