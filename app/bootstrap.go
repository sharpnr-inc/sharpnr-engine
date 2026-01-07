package app

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"
	hr_employee_services "sharpnr-engine/modules/hr/employees/services"
	"sharpnr-engine/sere"

	"gorm.io/gorm"
)

type AppModules struct {
	hrService hr_employee_services.HrEmployeeServiceInterface
}

func (m *AppModules) HrService() hr_employee_services.HrEmployeeServiceInterface {
	return m.hrService
}

func BootstrapModules(
	db *gorm.DB,
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
) *AppModules {

	hrEmployeeRepo := hr_employee_repo.NewHrEmployeeRepository(db)
	hrEmployeeService := hr_employee_services.NewHrEmployeeService(hrEmployeeRepo)

	sereEngine.UseTxExecutor(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)
			return fn(txCtx)
		})
	})

	RegisterAllWorkFlow()
	// RegisterSereEventHandler(sereEngine,)

	return &AppModules{
		hrService: hrEmployeeService,
	}

}
