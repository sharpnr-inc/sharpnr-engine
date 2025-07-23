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
	opRepo   interfaces.OperationsRepository
	termRepo interfaces.TerminalRepository
}

func NewModule(

	opRepo interfaces.OperationsRepository,
	termRepo interfaces.TerminalRepository,
) *Module {
	return &Module{

		opRepo:   opRepo,
		termRepo: termRepo,
	}
}

func (m *Module) Type() constants.OrgType {
	return constants.OrgTypeBase
}

func (m *Module) Initialize(org *models.Organization) error {
	fmt.Println("BASE MODULE: Initializing default")

	// hrDept := models.Department{
	// 	Name: "Human Resources",
	// 	OrganizationID: org.SharpnrOrgUUID,
	// 	ModuleType: constants.OrgTypeBase,
	// }

	return nil

}
