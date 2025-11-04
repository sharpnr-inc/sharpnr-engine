package sharpnruirenderer

import (
	"context"
	"encoding/json"
	"sharpnr-engine/core/constants"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func BarRendererFactory(ctx context.Context, orgId uuid.UUID, orgType constants.OrgType) (datatypes.JSON, error) {
	uiRenderTree := UIRendererTree{
		OrganizationID:           orgId,
		UIRendererTreeID:         uuid.New(),
		OrganizationBusinessType: string(orgType),
		MainMenuRenderTree: MainMenuRenderTree{
			SharpHospitality: SharpHospitalityTree{
				HasAccess: true,
				SharpHospitalityModule: SharpHospitalityModule{
					Hotel: HotelModule{
						HasAccess: false,
					},

					Restaurant: RestaurantModule{
						HasAccess: false,
					},

					Bar: BarModule{
						HasAccess: true,
						BarSectionAccess: BarSectionAccess{
							Menu:      true,
							Orders:    true,
							Inventory: true,
						},
					},

					Menu: MenuModule{
						HasAccess: true,
						MenuSectionAccess: MenuSectionAccess{
							BarMenu:        true,
							HotelMenu:      false,
							RestaurantMenu: false,
						},
					},

					Kitchen: KitchenModule{
						HasAccess: true,
						KitchenSectionAccess: KitchenSectionAccess{
							Orders:    true,
							Menu:      true,
							Inventory: true,
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(uiRenderTree)
	if err != nil {
		return nil, err
	}

	return datatypes.JSON(data), nil
}
