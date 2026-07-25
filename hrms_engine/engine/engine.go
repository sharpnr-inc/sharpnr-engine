package hrmsengine

import (
	hrmsEmployeeServices "sharpnr-engine/hrms_engine/employees/services"
)

// HRMSEngine aggregates every HRMS module service (employees, attendance,
// payroll, recruitment, ...). Bootstrapped once in app and exposed to the
// route layer through ModulesProvider.HRMS().
type HRMSEngine struct {
	employeeService hrmsEmployeeServices.HRMSEmployeeServiceInterface

	// attendanceService  attendance_services.AttendanceServiceInterface
	// payrollService     payroll_services.PayrollServiceInterface
	// recruitmentService recruitment_services.RecruitmentServiceInterface
}

func NewHRMSEngine(
	employeeService hrmsEmployeeServices.HRMSEmployeeServiceInterface,
) *HRMSEngine {
	return &HRMSEngine{
		employeeService: employeeService,
	}
}

/* ---- getters (ONLY what you want to expose) ---- */

func (e *HRMSEngine) EmployeeService() hrmsEmployeeServices.HRMSEmployeeServiceInterface {
	return e.employeeService
}
