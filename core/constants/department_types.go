package constants

type BaseDepartmentType string

const (
	Accounting     BaseDepartmentType = "ACCOUNTING"
	HumanResources BaseDepartmentType = "HUMAN RESOURCES"
	Sales          BaseDepartmentType = "SALES"
	Marketing      BaseDepartmentType = "MARKETING"
)

func (b BaseDepartmentType) IsValid() bool {
	switch b {
	case Accounting, HumanResources, Sales, Marketing:
		return true
	}
	return false
}

func (b BaseDepartmentType) String() string {
	return string(b)
}
