package models

import "github.com/google/uuid"

// OrgPolicy represents an organization's policy definition.
// Each policy belongs to a single Organization.
type OrgPolicy struct {
	// Primary key
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Foreign key to Organization
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`

	// Example: You may later store the policy document in S3 and keep the link here.
	S3BucketLink string `gorm:"type:text" json:"s3_bucket_link,omitempty"`

	// Relation (optional: if you want to preload organization when querying policy)
	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
