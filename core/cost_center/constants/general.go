package orgCostCenterConstants

// CostCenterStatus is the lifecycle of a cost center. ARCHIVED is distinct from
// a soft delete: an archived cost center takes no new departments or postings
// but stays readable, so historical reports that reference it still resolve.
type CostCenterStatus string

const (
	CostCenterActive   CostCenterStatus = "ACTIVE"
	CostCenterInactive CostCenterStatus = "INACTIVE"
	CostCenterArchived CostCenterStatus = "ARCHIVED"
)

func (s CostCenterStatus) Valid() bool {
	switch s {
	case CostCenterActive, CostCenterInactive, CostCenterArchived:
		return true
	default:
		return false
	}
}

/* ---------------------------------- <***> --------------------------------- */

// CostCenterType describes what kind of site the cost center is. Descriptive
// only — it drives labelling and filtering in the UI, never behaviour.
type CostCenterType string

const (
	CostCenterTypeHeadOffice CostCenterType = "HEAD_OFFICE"
	CostCenterTypeBranch     CostCenterType = "BRANCH"
	CostCenterTypeOffice     CostCenterType = "OFFICE"
	CostCenterTypeWarehouse  CostCenterType = "WAREHOUSE"
	CostCenterTypePlant      CostCenterType = "PLANT"
	CostCenterTypeSite       CostCenterType = "SITE"
)

func (t CostCenterType) Valid() bool {
	switch t {
	case CostCenterTypeHeadOffice, CostCenterTypeBranch, CostCenterTypeOffice,
		CostCenterTypeWarehouse, CostCenterTypePlant, CostCenterTypeSite:
		return true
	default:
		return false
	}
}
