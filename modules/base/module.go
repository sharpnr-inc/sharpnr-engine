package base

// import (
// 	"fmt"

// 	"sharpnr-engine/core/constants"
// 	"sharpnr-engine/core/interfaces"
// 	"sharpnr-engine/core/models"
// 	"sharpnr-engine/modules"
// )

// type InitialiseBaseModule struct {
// 	orgId InitBaseModule
// }

// var _ modules.IModule = (*Module)(nil)

// type Module struct {
// 	opRepo   interfaces.OperationsRepository
// 	termRepo interfaces.TerminalRepository
// }

// func NewModule(

// 	opRepo interfaces.OperationsRepository,
// 	termRepo interfaces.TerminalRepository,
// ) *Module {
// 	return &Module{

// 		opRepo:   opRepo,
// 		termRepo: termRepo,
// 	}
// }

// func (m *Module) Type() constants.OrgType {
// 	return constants.OrgTypeBase
// }

// func (m *Module) Initialize(org *models.Organization) error {
// 	fmt.Println("BASE MODULE: Initializing default")

// 	// hrDept := models.Department{
// 	// 	Name: "Human Resources",
// 	// 	OrganizationID: org.SharpnrOrgUUID,
// 	// 	ModuleType: constants.OrgTypeBase,
// 	// }

// 	return nil

// }
