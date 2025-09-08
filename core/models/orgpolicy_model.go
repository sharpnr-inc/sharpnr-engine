package models

import "github.com/google/uuid"

// This is Org Policy Model
type OrgPolicy struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	S3_bucket_link string
}
