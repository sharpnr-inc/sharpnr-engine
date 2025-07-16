package base

import (
	"fmt"

	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"github.com/sharpnr-inc/sharpnr-engine/modules"
)

type InitialiseBaseModule struct {
	orgId InitBaseModule
}

var _ modules.IModule = (*Module)(nil)

type Module struct {
	depRepo  interfaces.DepartmentRepository
	opRepo   interfaces.OperationsRepository
	termRepo interfaces.TerminalRepository
}

func NewModule(
	depRepo interfaces.DepartmentRepository,
	opRepo interfaces.OperationsRepository,
	termRepo interfaces.TerminalRepository,
) *Module {
	return &Module{
		depRepo:  depRepo,
		opRepo:   opRepo,
		termRepo: termRepo,
	}
}

func (m *Module) Type() constants.OrgType {
	return constants.OrgTypeBase
}

func (m *Module) Intialize(org *models.Organization) error {
	fmt.Println("BASE MODULE: Initializing default")
}
