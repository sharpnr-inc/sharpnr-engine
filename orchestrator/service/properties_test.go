// Package orchestrator_service tests.
//
// This file is the reference template for unit-testing orchestrator services.
// New service tests should follow the same conventions:
//
//  1. Test from the SAME package (white-box). The concrete `orchestratorService`
//     struct is unexported, so testing from inside the package lets us build it
//     directly with only the fields a given method needs — no full constructor,
//     no real database, no real downstream engines.
//
//  2. Mock collaborators with EMBEDDED-INTERFACE STUBS. Each dependency
//     interface is large, but a method under test only touches a few methods.
//     A stub embeds the interface (so it satisfies the whole contract for free)
//     and overrides only what the test exercises via a function field. Any
//     unexpected call lands on the nil embedded interface and panics loudly —
//     which is exactly the signal we want when a method starts depending on
//     something it shouldn't.
//
//  3. Assert with testify's `require`. `require.*` stops the test on the first
//     failed assertion (like t.Fatal), which keeps failure output readable.
//     Prefer it over hand-written `if got != want { t.Errorf(...) }`.
//     Convention: require.Equal(t, EXPECTED, ACTUAL) — expected value first.
//
//  4. Wire downstream-error propagation explicitly. Services wrap errors with
//     %w; assert with require.ErrorIs so the test survives message rewording.
//
// The two unimplemented methods (ListAllBusinessOpsModules and
// ListAllInfraModulesOfProperty) are intentionally NOT tested here.
package orchestrator_service

import (
	"context"
	"errors"
	"testing"

	barservices "sharpnr-engine/business_ops/hospitality/bar/services"
	hotelservices "sharpnr-engine/business_ops/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/business_ops/hospitality/kitchen/services"
	restaurantservices "sharpnr-engine/business_ops/hospitality/restaurant/services"
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_property_services "sharpnr-engine/core/org_property/services"
	organization_constants "sharpnr-engine/core/organization/constants"
	organization_services "sharpnr-engine/core/organization/services"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"
	sentinel_utils "sharpnr-engine/sentinel_engine/utils"

	hospitalitydomain "sharpnr-engine/app/hospitality_domain"

	"github.com/go-openapi/testify/v2/require"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// allowAllSentinelRepo is a fake sentinelengine.Repository that makes every
// principal an "owner": the default owner policy permits any action within the
// principal's own org, so IsAuthorized returns Allowed. Used by service tests
// that exercise logic past the authorization gate.
type allowAllSentinelRepo struct{}

func (allowAllSentinelRepo) ListEnabledPolicies(_ context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	return []*sentinelengine_models.SentinelPolicy{{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Policy:         sentinel_utils.OwnerPolicyText,
		Enabled:        true,
	}}, nil
}

func (allowAllSentinelRepo) UserRoleSlugs(_ context.Context, _, _ uuid.UUID) ([]string, error) {
	return []string{sentinel_utils.OwnerRoleSlug}, nil
}

func (allowAllSentinelRepo) UserGroupClosure(_ context.Context, _, _ uuid.UUID) ([]sentinelengine.GroupMembership, error) {
	return nil, nil
}

// allowAllEngine builds an engine wired to the permissive repo above.
func allowAllEngine() *sentinelengine.Sentinel {
	return sentinelengine.New(allowAllSentinelRepo{})
}

/* ========================================================================== */
/*                                   STUBS                                    */
/* ========================================================================== */
/*
   Pattern: embed the real interface, override only the methods under test.

       type stubFoo struct {
           pkg.FooInterface                      // nil embedded value
           doThing func(...) (...)               // the one method we care about
       }
       func (s *stubFoo) DoThing(...) (...) {    // override
           return s.doThing(...)
       }

   Because the interface is embedded, stubFoo satisfies the FULL interface even
   though we wrote one method. Calling any other method dereferences the nil
   embedded interface and panics — a deliberate tripwire for unexpected calls.
*/

// stubOrgPropertyService fakes the org-property core service.
type stubOrgPropertyService struct {
	org_property_services.OrganizationPropertyServicesInterfaces

	createPropertyDefault  func(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error
	getPropertiesOfOrg     func(ctx context.Context, orgID uuid.UUID) ([]org_property_contracts.PropertiesReader, error)
	getPropertyDetailsByID func(ctx context.Context, orgID, propertyID uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error)
}

func (s *stubOrgPropertyService) CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error {
	return s.createPropertyDefault(ctx, input)
}

func (s *stubOrgPropertyService) GetPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) ([]org_property_contracts.PropertiesReader, error) {
	return s.getPropertiesOfOrg(ctx, orgID)
}

func (s *stubOrgPropertyService) GetPropertyDetailsByID(ctx context.Context, orgID, propertyID uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error) {
	return s.getPropertyDetailsByID(ctx, orgID, propertyID)
}

// stubOrganizationService fakes the organization core service.
type stubOrganizationService struct {
	organization_services.OrganizationServiceInterface

	getStatus    func(ctx context.Context, orgID uuid.UUID) (organization_constants.OrgSetupStatusType, error)
	updateStatus func(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error
}

func (s *stubOrganizationService) GetStatus(ctx context.Context, orgID uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
	return s.getStatus(ctx, orgID)
}

func (s *stubOrganizationService) UpdateStatus(ctx context.Context, orgID uuid.UUID, status organization_constants.OrgSetupStatusType) error {
	return s.updateStatus(ctx, orgID, status)
}

// stubTxManager runs the transactional callback inline, mirroring the
// production contract: whatever the callback returns is returned to the caller,
// and a nil *gorm.DB is fine because the downstream stubs never touch it.
type stubTxManager struct {
	withTx func(ctx context.Context, fn func(tx *gorm.DB) error) error
}

func (s *stubTxManager) WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	if s.withTx != nil {
		return s.withTx(ctx, fn)
	}
	return fn(nil)
}

// --- hospitality sub-service stubs ---------------------------------------- //
// These are the leaf services held inside HospitalityDomain. Each test wires up
// only the ones its scenario should reach.

type stubHotelService struct {
	hotelservices.HotelServiceInterface
	setupDefaultHotel func(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error)
}

func (s *stubHotelService) SetupDefaultHotel(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error) {
	return s.setupDefaultHotel(ctx, orgID, propertyID)
}

type stubBarService struct {
	barservices.BarServicesInterface
	setupDefaultBar func(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error)
}

func (s *stubBarService) SetupDefaultBar(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error) {
	return s.setupDefaultBar(ctx, orgID, propertyID)
}

type stubRestaurantService struct {
	restaurantservices.RestaurantServiceInterface
	setupDefaultRestaurant func(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error)
}

func (s *stubRestaurantService) SetupDefaultRestaurant(ctx context.Context, orgID, propertyID uuid.UUID) (uuid.UUID, error) {
	return s.setupDefaultRestaurant(ctx, orgID, propertyID)
}

type stubKitchenService struct {
	kitchenservices.KitchenServiceInterface
	createDefaultKitchen func(ctx context.Context, orgID, propertyID uuid.UUID, module string) error
}

func (s *stubKitchenService) CreateDefaultKitchen(ctx context.Context, orgID, propertyID uuid.UUID, module string) error {
	return s.createDefaultKitchen(ctx, orgID, propertyID, module)
}

/* ========================================================================== */
/*                                  HELPERS                                   */
/* ========================================================================== */

// hospitalityStubs bundles the optional leaf stubs for a HospitalityDomain.
// Leave a field nil to assert "this service must not be reached" — the domain
// is built with a nil leaf, so any call to it panics.
type hospitalityStubs struct {
	hotel      *stubHotelService
	bar        *stubBarService
	restaurant *stubRestaurantService
	kitchen    *stubKitchenService
}

// newHospitalityDomain builds a real *HospitalityDomain (a concrete struct, not
// an interface) from whichever stubs are supplied. menu and order-engine are
// always nil here because the property orchestration never uses them.
func newHospitalityDomain(h hospitalityStubs) *hospitalitydomain.HospitalityDomain {
	var (
		hotel      hotelservices.HotelServiceInterface
		bar        barservices.BarServicesInterface
		restaurant restaurantservices.RestaurantServiceInterface
		kitchen    kitchenservices.KitchenServiceInterface
	)
	if h.hotel != nil {
		hotel = h.hotel
	}
	if h.bar != nil {
		bar = h.bar
	}
	if h.restaurant != nil {
		restaurant = h.restaurant
	}
	if h.kitchen != nil {
		kitchen = h.kitchen
	}
	return hospitalitydomain.NewHospitalityDomain(hotel, bar, restaurant, kitchen, nil, nil)
}

/* ========================================================================== */
/*                          GetAllPropertiesOfOrg                             */
/* ========================================================================== */

// Happy path: the org-property service returns rows, and the orchestrator maps
// every field onto the response DTO and forwards the orgID downstream.
func TestGetAllPropertiesOfOrg_MapsAndReturnsProperties(t *testing.T) {
	// --- arrange ---
	orgID := uuid.New()
	emplID := uuid.New()
	p1 := org_property_contracts.PropertiesReader{
		ID: uuid.New(), Name: "Grand", City: "Pune", State: "MH",
		Code: "GR1", AddressLine: "MG Road", Country: "IN",
	}
	p2 := org_property_contracts.PropertiesReader{
		ID: uuid.New(), Name: "Plaza", City: "Mumbai", State: "MH",
		Code: "PL2", AddressLine: "Marine Dr", Country: "IN",
	}

	var gotOrgID uuid.UUID
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		orgPropertyService: &stubOrgPropertyService{
			getPropertiesOfOrg: func(_ context.Context, orgID uuid.UUID) ([]org_property_contracts.PropertiesReader, error) {
				gotOrgID = orgID // capture for the forwarding assertion
				return []org_property_contracts.PropertiesReader{p1, p2}, nil
			},
		},
	}

	// --- act ---
	res, err := svc.GetAllPropertiesOfOrg(context.Background(), orgID, emplID)

	// --- assert ---
	require.NoError(t, err)
	require.Equal(t, orgID, gotOrgID, "orgID must be forwarded to the core service")
	require.Len(t, res.Properties, 2)

	require.Equal(t, orchestrator_dtos.PropertiesReaderSlice{
		ID: p1.ID, Name: p1.Name, City: p1.City, State: p1.State,
		Code: p1.Code, AddressLine: p1.AddressLine, Country: p1.Country,
	}, res.Properties[0], "first property must be mapped field-for-field")
}

// Edge case: no rows yields an empty (non-nil-deref) slice and no error.
func TestGetAllPropertiesOfOrg_Empty(t *testing.T) {
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		orgPropertyService: &stubOrgPropertyService{
			getPropertiesOfOrg: func(_ context.Context, _ uuid.UUID) ([]org_property_contracts.PropertiesReader, error) {
				return nil, nil
			},
		},
	}

	res, err := svc.GetAllPropertiesOfOrg(context.Background(), uuid.New(), uuid.New())

	require.NoError(t, err)
	require.Empty(t, res.Properties)
}

// Error path: a downstream error surfaces and the result is nil.
func TestGetAllPropertiesOfOrg_PropagatesError(t *testing.T) {
	wantErr := errors.New("db down")
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		orgPropertyService: &stubOrgPropertyService{
			getPropertiesOfOrg: func(_ context.Context, _ uuid.UUID) ([]org_property_contracts.PropertiesReader, error) {
				return nil, wantErr
			},
		},
	}

	res, err := svc.GetAllPropertiesOfOrg(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, res)
}

/* ========================================================================== */
/*                             GetPropertyDetail                              */
/* ========================================================================== */

// Happy path: detail fields are mapped and both IDs are forwarded.
func TestGetPropertyDetail_MapsDetails(t *testing.T) {
	orgID := uuid.New()
	propertyID := uuid.New()
	detail := &org_property_contracts.PropertyDetailsReader{
		Name: "Grand", Code: "GR1", City: "Pune",
		AddressLine: "MG Road", State: "MH", Country: "IN",
	}

	var gotOrg, gotProp uuid.UUID
	svc := &orchestratorService{
		orgPropertyService: &stubOrgPropertyService{
			getPropertyDetailsByID: func(_ context.Context, orgID, propertyID uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error) {
				gotOrg, gotProp = orgID, propertyID
				return detail, nil
			},
		},
	}

	res, err := svc.GetPropertyDetail(context.Background(), orgID, propertyID)

	require.NoError(t, err)
	require.Equal(t, orgID, gotOrg)
	require.Equal(t, propertyID, gotProp)
	require.Equal(t, &orchestrator_dtos.PropertyDetails_res_dto{
		Name: detail.Name, Code: detail.Code, City: detail.City,
		AddressLine: detail.AddressLine, State: detail.State, Country: detail.Country,
	}, res)
}

// Error path: downstream error surfaces, result is nil.
func TestGetPropertyDetail_PropagatesError(t *testing.T) {
	wantErr := errors.New("not found")
	svc := &orchestratorService{
		orgPropertyService: &stubOrgPropertyService{
			getPropertyDetailsByID: func(_ context.Context, _, _ uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error) {
				return nil, wantErr
			},
		},
	}

	res, err := svc.GetPropertyDetail(context.Background(), uuid.New(), uuid.New())

	require.ErrorIs(t, err, wantErr)
	require.Nil(t, res)
}

/* ========================================================================== */
/*                               CreateProperty                               */
/* ========================================================================== */
/*
   CreateProperty runs inside a transaction and has a small state machine:
     1. map the request -> CreateOrgPropertyInput and persist it
     2. read the org's setup status
     3. only when status == NOT_CONFIGURED, promote it to CONFIGURED
   The tests below pin each branch and each failure point.
*/

// NOT_CONFIGURED org: request is mapped correctly and status is promoted.
func TestCreateProperty_NotConfigured_PromotesToConfigured(t *testing.T) {
	orgID := uuid.New()
	emplID := uuid.New()
	req := &orchestrator_dtos.CreatePropertyReq_dto{
		Name: "Grand", Code: "GR1", AddressLine: "MG Road",
		City: "Pune", State: "MH", Country: "IN", Pincode: "411001",
	}

	var gotInput org_property_contracts.CreateOrgPropertyInput
	var updatedStatus organization_constants.OrgSetupStatusType
	updateCalled := false

	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		txManager:      &stubTxManager{}, // runs the callback inline
		orgPropertyService: &stubOrgPropertyService{
			createPropertyDefault: func(_ context.Context, input org_property_contracts.CreateOrgPropertyInput) error {
				gotInput = input
				return nil
			},
		},
		organizationService: &stubOrganizationService{
			getStatus: func(_ context.Context, _ uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
				return organization_constants.OrgConfigStatusNotConfigured, nil
			},
			updateStatus: func(_ context.Context, _ uuid.UUID, status organization_constants.OrgSetupStatusType) error {
				updateCalled = true
				updatedStatus = status
				return nil
			},
		},
	}

	err := svc.CreateProperty(context.Background(), orgID, emplID, req)

	require.NoError(t, err)
	// Request DTO -> contract input mapping (note OrgID is injected by the service).
	require.Equal(t, org_property_contracts.CreateOrgPropertyInput{
		OrgID: orgID, Name: req.Name, Code: req.Code, AddressLine: req.AddressLine,
		City: req.City, State: req.State, Country: req.Country, Pincode: req.Pincode,
	}, gotInput)
	require.True(t, updateCalled, "status must be updated when org is NOT_CONFIGURED")
	require.Equal(t, organization_constants.OrgConfigStatusConfigured, updatedStatus)
}

// Already CONFIGURED org: status must be left untouched.
func TestCreateProperty_AlreadyConfigured_DoesNotUpdateStatus(t *testing.T) {
	updateCalled := false
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		txManager:      &stubTxManager{},
		orgPropertyService: &stubOrgPropertyService{
			createPropertyDefault: func(_ context.Context, _ org_property_contracts.CreateOrgPropertyInput) error { return nil },
		},
		organizationService: &stubOrganizationService{
			getStatus: func(_ context.Context, _ uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
				return organization_constants.OrgConfigStatusConfigured, nil
			},
			updateStatus: func(_ context.Context, _ uuid.UUID, _ organization_constants.OrgSetupStatusType) error {
				updateCalled = true
				return nil
			},
		},
	}

	err := svc.CreateProperty(context.Background(), uuid.New(), uuid.New(), &orchestrator_dtos.CreatePropertyReq_dto{})

	require.NoError(t, err)
	require.False(t, updateCalled, "status must NOT be updated when already CONFIGURED")
}

// Persist fails: the error short-circuits before the status is ever read.
func TestCreateProperty_CreateFails_DoesNotCheckStatus(t *testing.T) {
	wantErr := errors.New("insert failed")
	statusChecked := false
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		txManager:      &stubTxManager{},
		orgPropertyService: &stubOrgPropertyService{
			createPropertyDefault: func(_ context.Context, _ org_property_contracts.CreateOrgPropertyInput) error { return wantErr },
		},
		organizationService: &stubOrganizationService{
			getStatus: func(_ context.Context, _ uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
				statusChecked = true
				return "", nil
			},
		},
	}

	err := svc.CreateProperty(context.Background(), uuid.New(), uuid.New(), &orchestrator_dtos.CreatePropertyReq_dto{})

	require.ErrorIs(t, err, wantErr)
	require.False(t, statusChecked, "status must not be read when create fails")
}

// Reading status fails: error propagates.
func TestCreateProperty_GetStatusFails_Propagates(t *testing.T) {
	wantErr := errors.New("status lookup failed")
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		txManager:      &stubTxManager{},
		orgPropertyService: &stubOrgPropertyService{
			createPropertyDefault: func(_ context.Context, _ org_property_contracts.CreateOrgPropertyInput) error { return nil },
		},
		organizationService: &stubOrganizationService{
			getStatus: func(_ context.Context, _ uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
				return "", wantErr
			},
		},
	}

	err := svc.CreateProperty(context.Background(), uuid.New(), uuid.New(), &orchestrator_dtos.CreatePropertyReq_dto{})

	require.ErrorIs(t, err, wantErr)
}

// Updating status fails: error propagates.
func TestCreateProperty_UpdateStatusFails_Propagates(t *testing.T) {
	wantErr := errors.New("update failed")
	svc := &orchestratorService{
		sentinelEngine: allowAllEngine(),
		txManager:      &stubTxManager{},
		orgPropertyService: &stubOrgPropertyService{
			createPropertyDefault: func(_ context.Context, _ org_property_contracts.CreateOrgPropertyInput) error { return nil },
		},
		organizationService: &stubOrganizationService{
			getStatus: func(_ context.Context, _ uuid.UUID) (organization_constants.OrgSetupStatusType, error) {
				return organization_constants.OrgConfigStatusNotConfigured, nil
			},
			updateStatus: func(_ context.Context, _ uuid.UUID, _ organization_constants.OrgSetupStatusType) error {
				return wantErr
			},
		},
	}

	err := svc.CreateProperty(context.Background(), uuid.New(), uuid.New(), &orchestrator_dtos.CreatePropertyReq_dto{})

	require.ErrorIs(t, err, wantErr)
}

/* ========================================================================== */
/*                  AddBusinessOperationModulesInProperty                     */
/* ========================================================================== */
/*
   For each hospitality module in the request, the orchestrator provisions
   `Quantity` default instances by dispatching to the matching leaf service.
   Healthcare modules are intentionally ignored until that domain is wired up.
*/

// Each supported module dispatches to its service, once per unit of quantity,
// with the org/property IDs forwarded.
func TestAddBusinessOpModules_SetsUpEachModuleType(t *testing.T) {
	orgID := uuid.New()
	emplID := uuid.New()
	propertyID := uuid.New()

	var hotelCalls, barCalls, restCalls int
	var hotelOrg, hotelProp uuid.UUID

	hd := newHospitalityDomain(hospitalityStubs{
		hotel: &stubHotelService{setupDefaultHotel: func(_ context.Context, org, prop uuid.UUID) (uuid.UUID, error) {
			hotelCalls++
			hotelOrg, hotelProp = org, prop
			return uuid.New(), nil
		}},
		bar: &stubBarService{setupDefaultBar: func(_ context.Context, _, _ uuid.UUID) (uuid.UUID, error) {
			barCalls++
			return uuid.New(), nil
		}},
		restaurant: &stubRestaurantService{setupDefaultRestaurant: func(_ context.Context, _, _ uuid.UUID) (uuid.UUID, error) {
			restCalls++
			return uuid.New(), nil
		}},
	})
	svc := &orchestratorService{hospitalityDomainSvc: hd}

	req := orchestrator_dtos.AddBusinessOpModulesRequestDto{
		HospitalityModules: []orchestrator_dtos.HospitalityModuleQuantity{
			{Module: hospitality_utilities.Hotel, Quantity: 1},
			{Module: hospitality_utilities.Bar, Quantity: 2},
			{Module: hospitality_utilities.Restaurant, Quantity: 3},
		},
	}

	err := svc.AddBusinessOperationModulesInProperty(context.Background(), orgID, emplID, propertyID, req)

	require.NoError(t, err)
	require.Equal(t, 1, hotelCalls)
	require.Equal(t, 2, barCalls)
	require.Equal(t, 3, restCalls)
	require.Equal(t, orgID, hotelOrg, "orgID must be forwarded to the leaf service")
	require.Equal(t, propertyID, hotelProp, "propertyID must be forwarded to the leaf service")
}

// Healthcare modules are ignored, and an empty request is a no-op.
// No hospitality stubs are wired, so any dispatch would panic on a nil leaf.
func TestAddBusinessOpModules_HealthcareIgnoredAndEmptyIsNoop(t *testing.T) {
	svc := &orchestratorService{hospitalityDomainSvc: newHospitalityDomain(hospitalityStubs{})}

	req := orchestrator_dtos.AddBusinessOpModulesRequestDto{
		HealthcareModules: []orchestrator_dtos.HealthCareModuleQuantity{{Quantity: 5}},
	}

	err := svc.AddBusinessOperationModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.NoError(t, err)
}

// A leaf-service error is wrapped (%w) and propagated — assert with ErrorIs.
func TestAddBusinessOpModules_SetupErrorIsWrapped(t *testing.T) {
	wantErr := errors.New("hotel boom")
	hd := newHospitalityDomain(hospitalityStubs{
		hotel: &stubHotelService{setupDefaultHotel: func(_ context.Context, _, _ uuid.UUID) (uuid.UUID, error) {
			return uuid.Nil, wantErr
		}},
	})
	svc := &orchestratorService{hospitalityDomainSvc: hd}

	req := orchestrator_dtos.AddBusinessOpModulesRequestDto{
		HospitalityModules: []orchestrator_dtos.HospitalityModuleQuantity{
			{Module: hospitality_utilities.Hotel, Quantity: 1},
		},
	}

	err := svc.AddBusinessOperationModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.ErrorIs(t, err, wantErr)
}

// Declared-but-unsupported modules (no setup service yet) must error.
// Table-driven: one sub-test per module via t.Run.
func TestAddBusinessOpModules_DeclaredButUnsupportedModule(t *testing.T) {
	svc := &orchestratorService{hospitalityDomainSvc: newHospitalityDomain(hospitalityStubs{})}

	unsupported := []hospitality_utilities.HospitalityModulesType{
		hospitality_utilities.NightClub,
		hospitality_utilities.Resort,
		hospitality_utilities.CloudKitchen_Takeaways,
	}
	for _, module := range unsupported {
		t.Run(string(module), func(t *testing.T) {
			req := orchestrator_dtos.AddBusinessOpModulesRequestDto{
				HospitalityModules: []orchestrator_dtos.HospitalityModuleQuantity{
					{Module: module, Quantity: 1},
				},
			}
			err := svc.AddBusinessOperationModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)
			require.Error(t, err)
		})
	}
}

// An unrecognised module value hits the default branch and errors.
func TestAddBusinessOpModules_UnknownModule(t *testing.T) {
	svc := &orchestratorService{hospitalityDomainSvc: newHospitalityDomain(hospitalityStubs{})}

	req := orchestrator_dtos.AddBusinessOpModulesRequestDto{
		HospitalityModules: []orchestrator_dtos.HospitalityModuleQuantity{
			{Module: hospitality_utilities.HospitalityModulesType("WATER_PARK"), Quantity: 1},
		},
	}

	err := svc.AddBusinessOperationModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.Error(t, err)
}

/* ========================================================================== */
/*                        AddInfraModulesInProperty                           */
/* ========================================================================== */
/*
   Current scope: only Hospitality KITCHEN infra is provisioned. Other
   hospitality infra types are skipped (not errored), and general/healthcare
   infra are ignored entirely.
*/

// KITCHEN infra: one CreateDefaultKitchen per unit of quantity, with IDs and
// the module name forwarded as a string.
func TestAddInfraModules_CreatesKitchenPerQuantity(t *testing.T) {
	orgID := uuid.New()
	propertyID := uuid.New()

	var calls int
	var gotOrg, gotProp uuid.UUID
	var gotModule string

	hd := newHospitalityDomain(hospitalityStubs{
		kitchen: &stubKitchenService{createDefaultKitchen: func(_ context.Context, org, prop uuid.UUID, module string) error {
			calls++
			gotOrg, gotProp, gotModule = org, prop, module
			return nil
		}},
	})
	svc := &orchestratorService{hospitalityDomainSvc: hd}

	req := orchestrator_dtos.AddInfraModulesRequestDto{
		HospitalityInfra: []orchestrator_dtos.HospitalityInfraModuleQuantity{
			{Module: hospitality_utilities.Kitchen, Quantity: 3},
		},
	}

	err := svc.AddInfraModulesInProperty(context.Background(), orgID, uuid.New(), propertyID, req)

	require.NoError(t, err)
	require.Equal(t, 3, calls)
	require.Equal(t, orgID, gotOrg)
	require.Equal(t, propertyID, gotProp)
	require.Equal(t, string(hospitality_utilities.Kitchen), gotModule)
}

// Non-kitchen hospitality infra is skipped silently (no call, no error).
// The kitchen stub records calls so we can assert it was never reached.
func TestAddInfraModules_NonKitchenSkippedNoError(t *testing.T) {
	called := false
	hd := newHospitalityDomain(hospitalityStubs{
		kitchen: &stubKitchenService{createDefaultKitchen: func(_ context.Context, _, _ uuid.UUID, _ string) error {
			called = true
			return nil
		}},
	})
	svc := &orchestratorService{hospitalityDomainSvc: hd}

	req := orchestrator_dtos.AddInfraModulesRequestDto{
		HospitalityInfra: []orchestrator_dtos.HospitalityInfraModuleQuantity{
			{Module: hospitality_utilities.BarCounter, Quantity: 2},
			{Module: hospitality_utilities.DiningHall, Quantity: 1},
		},
	}

	err := svc.AddInfraModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.NoError(t, err)
	require.False(t, called, "non-kitchen infra must not reach the kitchen service")
}

// A kitchen-creation error is wrapped (%w) and propagated.
func TestAddInfraModules_KitchenErrorIsWrapped(t *testing.T) {
	wantErr := errors.New("kitchen boom")
	hd := newHospitalityDomain(hospitalityStubs{
		kitchen: &stubKitchenService{createDefaultKitchen: func(_ context.Context, _, _ uuid.UUID, _ string) error {
			return wantErr
		}},
	})
	svc := &orchestratorService{hospitalityDomainSvc: hd}

	req := orchestrator_dtos.AddInfraModulesRequestDto{
		HospitalityInfra: []orchestrator_dtos.HospitalityInfraModuleQuantity{
			{Module: hospitality_utilities.Kitchen, Quantity: 1},
		},
	}

	err := svc.AddInfraModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.ErrorIs(t, err, wantErr)
}

// General and healthcare infra are ignored entirely.
// No kitchen stub is wired, so any dispatch would panic on a nil leaf.
func TestAddInfraModules_GeneralAndHealthcareIgnored(t *testing.T) {
	svc := &orchestratorService{hospitalityDomainSvc: newHospitalityDomain(hospitalityStubs{})}

	req := orchestrator_dtos.AddInfraModulesRequestDto{
		GeneralModules:  []orchestrator_dtos.GeneralInfraModuleQuantity{{Quantity: 2}},
		HealthCareInfra: []orchestrator_dtos.HealthCareInfraModuleQuantity{{Quantity: 1}},
	}

	err := svc.AddInfraModulesInProperty(context.Background(), uuid.New(), uuid.New(), uuid.New(), req)

	require.NoError(t, err)
}
