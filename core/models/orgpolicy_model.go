package models

import "github.com/google/uuid"

// This is Org Policy Model
type OrgPolicy struct {
	ID             uuid.UUID // Make this Primary
	OrganizationID uuid.UUID // make this Forein key
	S3_bucket_link string    // Not needed for now
}
