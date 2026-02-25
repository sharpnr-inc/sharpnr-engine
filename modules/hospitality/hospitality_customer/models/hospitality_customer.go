package hospitalityCustomersModels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

//
// ============================================================
// 1️⃣ Hospitality Domain Customer (Global Shared Customer)
// ============================================================
//

type HospitalityDomainCustomer struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name   string `gorm:"type:varchar(150);not null"`
	Email  string `gorm:"type:varchar(150);index"`
	Mobile string `gorm:"type:varchar(20);not null;index"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HospitalityDomainCustomer) TableName() string {
	return "hospitality.domain_customers"
}

//
// ============================================================
// 2️⃣ Customer ↔ Organization (Tenant Boundary)
// ============================================================
// Same customer can belong to multiple organizations,
// but cannot be duplicated inside the same organization.
//

type HospitalityDomainCustomerOrganization struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HospitalityDomainCustomersID uuid.UUID `gorm:"type:uuid;not null;index:idx_customer_org,unique"`
	OrganizationID               uuid.UUID `gorm:"type:uuid;not null;index:idx_customer_org,unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HospitalityDomainCustomerOrganization) TableName() string {
	return "hospitality.domain_customer_organizations"
}

//
// ============================================================
// 3️⃣ Customer Profile (Flexible JSON Per Customer)
// ============================================================
// Used for restaurant/hotel-specific preferences.
//

type HospitalityDomainCustomerProfile struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CustomerID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex"`
	Data       datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HospitalityDomainCustomerProfile) TableName() string {
	return "hospitality.domain_customer_profiles"
}

//
// ============================================================
// 4️⃣ Customer Identification (Extensible ID System)
// ============================================================
// Supports: aadhar | passport | pan | gst | driving_license | emirates_id
//

type HospitalityDomainCustomerIdentification struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HospitalityDomainCustomerID uuid.UUID `gorm:"type:uuid;not null;index:idx_customer_id_type,unique"`

	IdentificationType string `gorm:"type:varchar(50);not null;index:idx_customer_id_type,unique"`

	IdentificationNumber string `gorm:"type:varchar(150);not null"`

	CountryCode string `gorm:"type:varchar(10)"`

	IsVerified bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HospitalityDomainCustomerIdentification) TableName() string {
	return "hospitality.domain_customer_identifications"
}
