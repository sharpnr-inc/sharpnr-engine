package models

type OrgDomainType string

const (
	DomainHospitality OrgDomainType = "HOSPITALITY"
	DomainMedical     OrgDomainType = "MEDICAL"
	DomainRetail      OrgDomainType = "RETAIL"
)

type SettingScope string

const (
	ScopeOrganization SettingScope = "ORGANIZATION"
	ScopeUnit         SettingScope = "UNIT" // hotel / hospital / store
)
