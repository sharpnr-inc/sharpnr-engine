package sector_register

type SharpnrSectorType string

const (
	ShapnrHospitalitySector SharpnrSectorType = "HOSPITALITY"
)

type HospitalitySectorModulesType string

const (
	HospitalityBarModule        HospitalitySectorModulesType = "BAR"
	HospitalityRestaurantModule HospitalitySectorModulesType = "RESTAURANT"
	HospitalityHotelModule      HospitalitySectorModulesType = "HOTEL"
)
