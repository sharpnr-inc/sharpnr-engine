package sere

import "context"

// ExecuteTx runs a workflow inside a transaction
func (s *SERE) ExecuteTx(
	ctx context.Context,
	name EventName,
	data map[string]any,
) (*Workflow, error) {

	if s.txExecutor == nil {
		// fallback: behave like Execute
		return s.Execute(ctx, name, data)
	}

	var wf *Workflow

	err := s.txExecutor(ctx, func(txCtx context.Context) error {

		w := &Workflow{
			Ctx:  txCtx, // 🔥 transactional context
			Name: name,
			Data: data,
		}

		s.mu.RLock()
		handlers := s.workflowSubscribers[string(name)]
		s.mu.RUnlock()

		for _, h := range handlers {
			if err := h(w); err != nil {
				return err // rollback
			}
		}

		wf = w
		return nil
	})

	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (s *SERE) UseTxExecutor(exec TxExecutor) {
	s.txExecutor = exec
}
