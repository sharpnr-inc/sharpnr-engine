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
// 	return constants.DefaultBase
// }

// func (m *Module) Initialize(org *models.Organization) error {
// 	fmt.Println("BASE MODULE: Initializing default")

// 	hrDept := models.Department{
// 		Name:           "Human Resources",
// 		OrganizationID: org.ID,
// 	}
// 	if err := m.db.Create(&hrDept).Error; err != nil {
// 		return fmt.Errorf("failed to create HR department: %w", err)
// 	}

// 	return nil

// }
