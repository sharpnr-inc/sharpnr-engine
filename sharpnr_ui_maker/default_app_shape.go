package sharpnruimaker

type ShapeMetaData struct {
	OrgID        string
	SharpnrId    string
	BusinessType string
}

type AccountingModuleShape struct {
	ToShow bool
}

type HRModuleShape struct {
	ToShow bool
}

type MarketingModuleShape struct {
	ToShow bool
}

type BaseModulesShape struct {
	Accounting AccountingModuleShape
	HR         HRModuleShape
	Marketing  MarketingModuleShape
}

type HospilatityModuleShape struct {
	Hotel      bool
	Bar        bool
	Kitchen    bool
	Restaurant bool
	PubAndClub bool
}

type MainModulesShape struct {
	Hospitality HospilatityModuleShape
}

type SharpnrAppUIShape struct {
	MetaData        ShapeMetaData `json:"METADATA"`
	BaseModuleShape BaseModulesShape
	MainModuleShape MainModulesShape
}
