// models/employee_org_profile.go
package employee_models

type EmployeeOrganizationProfile struct {
	ID         string `gorm:"type:uuid;primaryKey"`
	EmployeeID string `gorm:"type:uuid;not null;uniqueIndex"`

	OrganizationID string `gorm:"type:uuid;not null;index"`

	Department string `gorm:"type:varchar(100)"`
	Position   string `gorm:"type:varchar(100)"`
	ManagerID  string `gorm:"type:uuid;index"`

	CreatedAt int64
	UpdatedAt int64
}

func (EmployeeOrganizationProfile) TableName() string {
	return "employee.employee_organization_profile"
}
