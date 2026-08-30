package org_property_services

import (
	"context"
	"errors"
	"testing"

	property_constants "sharpnr-engine/core/org_property/constants"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_property_models "sharpnr-engine/core/org_property/models"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* --------------------------------- fakes ---------------------------------- */

// fakeEngine records the access request it was asked about and answers with a
// canned decision/error.
type fakeEngine struct {
	lastReq  sentinelEngine.AccessRequest
	calls    int
	response error // nil = allow; ErrAccessDenied = deny; anything else = infra error
}

func (f *fakeEngine) Authorize(_ context.Context, req sentinelEngine.AccessRequest) error {
	f.calls++
	f.lastReq = req
	return f.response
}

func (f *fakeEngine) IsAuthorized(ctx context.Context, req sentinelEngine.AccessRequest) (bool, error) {
	err := f.Authorize(ctx, req)
	if errors.Is(err, sentinelEngine.ErrAccessDenied) {
		return false, nil
	}
	return err == nil, err
}

func (f *fakeEngine) InvalidateOrganization(uuid.UUID) {}

// fakePropertyRepo captures created properties.
type fakePropertyRepo struct {
	created []*org_property_models.OrganizationProperty
}

func (f *fakePropertyRepo) CreateProperty(_ context.Context, property *org_property_models.OrganizationProperty) error {
	f.created = append(f.created, property)
	return nil
}

func (f *fakePropertyRepo) GetAllPropertiesByOrgID(context.Context, uuid.UUID) ([]*org_property_models.OrganizationProperty, error) {
	return nil, nil
}

func (f *fakePropertyRepo) GetPropertyByID(context.Context, uuid.UUID) (*org_property_models.OrganizationProperty, error) {
	return nil, nil
}

// GetPropertiesByIDs backs the Sentinel fact resolver. It answers from what the
// fake has created, org-scoped, so a test can assert that a property belonging
// to another tenant yields no facts.
func (f *fakePropertyRepo) GetPropertiesByIDs(
	_ context.Context, orgID uuid.UUID, ids []uuid.UUID,
) ([]*org_property_models.OrganizationProperty, error) {
	want := make(map[uuid.UUID]struct{}, len(ids))
	for _, id := range ids {
		want[id] = struct{}{}
	}
	var out []*org_property_models.OrganizationProperty
	for _, p := range f.created {
		if p.OrganizationID != orgID {
			continue
		}
		if _, ok := want[p.ID]; ok {
			out = append(out, p)
		}
	}
	return out, nil
}

// fakeTxManager runs the function immediately; the fake repo ignores the tx.
type fakeTxManager struct{}

func (fakeTxManager) WithTx(_ context.Context, fn func(tx *gorm.DB) error) error {
	return fn(nil)
}

/* --------------------------------- tests ---------------------------------- */

func newServiceForTest(engine *fakeEngine) (*organizationPropertyService, *fakePropertyRepo) {
	repo := &fakePropertyRepo{}
	var engineIface sentinelEngine.SentinelEngineInterface
	if engine != nil {
		engineIface = engine
	}
	svc := NewOrganizationPropertyService(repo, fakeTxManager{}, nil, engineIface)
	return svc, repo
}

func sampleInput(orgID, userID uuid.UUID) org_property_contracts.CreateOrgPropertyInput {
	return org_property_contracts.CreateOrgPropertyInput{
		OrgID:       orgID,
		UserID:      userID,
		Name:        "Sharpnr Grand",
		Code:        "SG-01",
		AddressLine: "12 MG Road",
		City:        "Mumbai",
		State:       "Maharashtra",
		Country:     "IN",
		Pincode:     "400001",
	}
}

func TestCreatePropertyDefault_Authorized(t *testing.T) {
	orgID, userID := uuid.New(), uuid.New()
	engine := &fakeEngine{response: nil}
	svc, repo := newServiceForTest(engine)

	err := svc.CreatePropertyDefault(context.Background(), sampleInput(orgID, userID))
	require.NoError(t, err)

	// The engine was asked exactly the right question...
	assert.Equal(t, 1, engine.calls)
	assert.Equal(t, orgID, engine.lastReq.OrgID)
	assert.Equal(t, userID, engine.lastReq.UserID)
	assert.Equal(t, property_constants.ActionCanCreateNewProperty, engine.lastReq.Action)
	assert.Equal(t, property_constants.ResourceTypeProperty, engine.lastReq.Resource.Type)
	assert.Empty(t, engine.lastReq.Resource.ID, "create is a type-level check — no instance id")

	// ...and the property was persisted with the mapped fields.
	require.Len(t, repo.created, 1)
	created := repo.created[0]
	assert.Equal(t, orgID, created.OrganizationID)
	assert.Equal(t, "Sharpnr Grand", created.Name)
	assert.Equal(t, "SG-01", created.Code)
	assert.Equal(t, "Mumbai", created.City)
	assert.NotEqual(t, uuid.Nil, created.ID, "property id is app-generated")
}

func TestCreatePropertyDefault_Denied(t *testing.T) {
	engine := &fakeEngine{response: sentinelEngine.ErrAccessDenied}
	svc, repo := newServiceForTest(engine)

	err := svc.CreatePropertyDefault(context.Background(), sampleInput(uuid.New(), uuid.New()))

	require.ErrorIs(t, err, sentinelEngine.ErrAccessDenied)
	assert.Empty(t, repo.created, "a denied request must not touch the repo")
}

func TestCreatePropertyDefault_EngineError(t *testing.T) {
	infraErr := errors.New("policy store unreachable")
	engine := &fakeEngine{response: infraErr}
	svc, repo := newServiceForTest(engine)

	err := svc.CreatePropertyDefault(context.Background(), sampleInput(uuid.New(), uuid.New()))

	require.ErrorIs(t, err, infraErr)
	assert.Empty(t, repo.created, "an engine failure must fail closed")
}

func TestCreatePropertyDefault_NilEngineFailsClosed(t *testing.T) {
	svc, repo := newServiceForTest(nil)

	err := svc.CreatePropertyDefault(context.Background(), sampleInput(uuid.New(), uuid.New()))

	require.Error(t, err, "a missing engine must never mean everything is allowed")
	assert.Empty(t, repo.created)
}
