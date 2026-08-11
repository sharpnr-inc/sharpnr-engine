package orchestrator_service

import (
	"context"

	restaurant_contracts "sharpnr-engine/property_business_ops/hospitality/restaurant/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	property_orchestrator "sharpnr-engine/orchestrator/service/property"

	"github.com/google/uuid"
)

// Property Operations in Orchestraor services interface.
//
// Implemented by property_orchestrator.PropertyOrchestrator, which is embedded
// in orchestratorService. Per-entity methods (hotel, restaurant, bar, kitchen,
// menu) belong in their own sub-interface here as they land — see
// orchestrator/service/property/doc.go for the layout.
type OrganizationPropertyOrchestrationInterface interface {

	// General Property service methods
	GetAllPropertiesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error)
	GetPropertyDetail(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error)
	CreateProperty(ctx context.Context, orgID, emplID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error

	// Business Modules
	AddBusinessOperationModulesInProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID, req orchestrator_dtos.AddBusinessOpModulesRequestDto) error
	ListAllBusinessOpsModules(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (orchestrator_dtos.PropertyBusinessOperationModulesResponseDto, error)

	// Infra Modules
	AddInfraModulesInProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID, request orchestrator_dtos.AddInfraModulesRequestDto) error
	ListAllInfraModulesOfProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID) (orchestrator_dtos.PropertyInfraModulesResponseDto, error)

	// Hospitality
	GetRestaurantDetailsByRestaurantID(ctx context.Context, orgID, employeeID, restaurantID uuid.UUID) (restaurant_contracts.RestaurantGeneralInfo_resp_dto, error)
	GetRestaurantFloorLayouts(ctx context.Context, orgID, employeeID, restaurantID uuid.UUID) (restaurant_contracts.RestaurantLayout_Response_dto, error)
}

// Fail at compile time in the property package's own terms, rather than at the
// facade, if a method is renamed or dropped over there.
var _ OrganizationPropertyOrchestrationInterface = (*property_orchestrator.PropertyOrchestrator)(nil)
