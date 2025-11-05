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

					Services: ServicesModule{
						HasAccess:            false,
						ServiceSectionAccess: ServicesSectionAccess{},
					},

					Inventory: InventoryModule{
						HasAccess: true,
						InventorySectionAccess: InventorySectionAccess{
							BarInventory:        true,
							HotelInventory:      true,
							RestaurantInventory: true,
							KitchenInventory:    true,
							GeneralInventory:    true,
						},
					},

					Housekeeping: HouseKeepingModule{
						HasAccess:                 false,
						HouseKeepingSectionAccess: HouseKeepingSectionAccess{},
					},

					Setting: SettingsModule{
						HasAccess: true,
						SettingsSectionAccess: SettingsSectionAccess{
							General:           true,
							Notifications:     true,
							Users:             true,
							Payments:          true,
							RoomConfiguration: true,
						},
					},
				},
			},

			SharpAccounting: SharpAccountingTree{
				HasAccess: true,
				SharpAccountingModule: SharpAccountingModule{
					Transactions: TransactionsModule{
						HasAccess: true,
					},
					Expenses: ExpensesModule{
						HasAccess: true,
					},
					Invoices: InvoicesModule{
						HasAccess: true,
					},
					Ledger: LedgerModule{
						HasAccess: true,
					},

					Customers: CustomersModule{
						HasAccess: true,
					},
					SuppliersAndVendors: SuppliersAndVendorsModule{
						HasAccess: true,
					},
					Reports: ReportsModule{
						HasAccess: true,
					},
				},
			},

			SharpHR: SharpHRTree{
				HasAccess: true,
			},
		},
	}

	data, err := json.Marshal(uiRenderTree)
	if err != nil {
		return nil, err
	}

	return datatypes.JSON(data), nil
}

// case --> Case two
func HotelRendererFactory()

// case --> case three
func RestaurantRendererFactory()

func BarHotelRendererFactory()

func BarRestaurantFactory()

func HotelRestaurantFactory()

func BarHotelRestaurantFactory()
