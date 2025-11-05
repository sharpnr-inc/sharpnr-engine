package sharpnruirenderer

type SharpHospitalityTree struct {
	HasAccess              bool                   `json:"has_access"`
	SharpHospitalityModule SharpHospitalityModule `json:"sub_module_section"`
}

type SharpHospitalityModule struct {
	Hotel        HotelModule        `json:"hotel"`
	Bar          BarModule          `json:"bar"`
	Restaurant   RestaurantModule   `json:"restaurant"`
	Services     ServicesModule     `json:"services"`
	Menu         MenuModule         `json:"menu"`
	Kitchen      KitchenModule      `json:"kitchen"`
	Inventory    InventoryModule    `json:"inventory"`
	Housekeeping HouseKeepingModule `json:"housekeeping"`
	Setting      SettingsModule     `json:"settings"`
	OrgConfig    OrgConfigModule    `json:"orgConfig"`
}

//-----------------Hotel----------------------------------------------

type HotelModule struct {
	HasAccess          bool               `json:"has_access"`
	HotelSectionAccess HotelSectionAccess `json:"sub_section_access"`
}

type HotelSectionAccess struct {
	Overview       bool `json:"overview"`
	Booking        bool `json:"booking"`
	Request        bool `json:"request"`
	RoomManagement bool `json:"room_management"`
}

//-----------------Bar----------------------------------------------

type BarModule struct {
	HasAccess        bool             `json:"has_access"`
	BarSectionAccess BarSectionAccess `json:"sub_section_access"`
}

type BarSectionAccess struct {
	Menu      bool `json:"menu"`
	Orders    bool `json:"orders"`
	Inventory bool `json:"inventory"`
}

//-----------------Restaurant----------------------------------------------

type RestaurantModule struct {
	HasAccess               bool                    `json:"has_access"`
	RestaurantSectionAccess RestaurantSectionAccess `json:"sub_section_access"`
}

type RestaurantSectionAccess struct {
	TableManagement       bool `json:"table_management"`
	Orders                bool `json:"orders"`
	Pickup                bool `json:"pickup"`
	ReservationAndBooking bool `json:"reservation_and_booking"`
}

//-----------------Services----------------------------------------------

type ServicesModule struct {
	HasAccess            bool                  `json:"has_access"`
	ServiceSectionAccess ServicesSectionAccess `json:"sub_section_access"`
}

type ServicesSectionAccess struct {
	RunningService bool `json:"running_service"`
	OrderStatus    bool `json:"order_status"`
	Assigned       bool `json:"assigned"`
	Upcoming       bool `json:"upcoming"`
	History        bool `json:"history"`
	Reviews        bool `json:"reviews"`
	Menu           bool `json:"menu"`
}

//-----------------Menu----------------------------------------------

type MenuModule struct {
	HasAccess         bool              `json:"has_access"`
	MenuSectionAccess MenuSectionAccess `json:"sub_section_access"`
}

type MenuSectionAccess struct {
	BarMenu        bool `json:"bar_menu"`
	HotelMenu      bool `json:"hotel_menu"`
	RestaurantMenu bool `json:"restaurant_menu"`
}

// ----------------------------Kitchen-----------------------------------
type KitchenModule struct {
	HasAccess            bool                 `json:"has_access"`
	KitchenSectionAccess KitchenSectionAccess `json:"sub_section_access"`
}

type KitchenSectionAccess struct {
	Orders    bool `json:"orders"`
	Menu      bool `json:"menu"`
	Inventory bool `json:"inventory"`
	Staff     bool `json:"staff"`
}

// ----------------------------Inventory-----------------------------------
type InventoryModule struct {
	HasAccess              bool                   `json:"has_access"`
	InventorySectionAccess InventorySectionAccess `json:"sub_section_access"`
}

type InventorySectionAccess struct {
	BarInventory        bool `json:"bar_inventory"`
	HotelInventory      bool `json:"hotel_inventory"`
	RestaurantInventory bool `json:"restaurant_inventory"`
	KitchenInventory    bool `json:"kitchen_inventory"`
	GeneralInventory    bool `json:"general_inventory"`
}

// ----------------------------HouseKeeping-----------------------------------
type HouseKeepingModule struct {
	HasAccess                 bool                      `json:"has_access"`
	HouseKeepingSectionAccess HouseKeepingSectionAccess `json:"sub_section_access"`
}

type HouseKeepingSectionAccess struct {
	Requests bool `json:"requests"`
	Laundry  bool `json:"laundry"`
}

// ----------------------------Setting-----------------------------------
type SettingsModule struct {
	HasAccess             bool                  `json:"has_access"`
	SettingsSectionAccess SettingsSectionAccess `json:"sub_section_access"`
}

type SettingsSectionAccess struct {
	General                 bool `json:"general"`
	Notifications           bool `json:"notifications"`
	Users                   bool `json:"users"`
	Payments                bool `json:"payments"`
	RoomConfiguration       bool `json:"room_configuration"`
	RestaurantConfiguration bool `json:"restaurant_configuration"`
}

// ----------------------------OrgConfig-----------------------------------

type OrgConfigModule struct {
	HasAccess              bool                   `json:"has_access"`
	OrgConfigSectionAccess OrgConfigSectionAccess `json:"sub_section_access"`
}

type OrgConfigSectionAccess struct {
	Overview             bool `json:"overview"`
	RestaurantConfigs    bool `json:"restaurant_configs"`
	BarConfigs           bool `json:"bar_configs"`
	HotelConfigs         bool `json:"hotel_configs"`
	Owners               bool `json:"owners"`
	RegistrationLicenses bool `json:"registration_licenses"`
	Integrations         bool `json:"integrations"`
	GSTTaxes             bool `json:"gst_taxes"`
	Branding             bool `json:"branding"`
	CustomerSupport      bool `json:"customer_support"`
}
