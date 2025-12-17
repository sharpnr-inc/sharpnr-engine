package uisetup

import "context"

type UIRendererSetupStep interface {
	Name() string
	Execute(ctx context.Context, c SetupContext) error
}
