package constants

type BaseModule string

const (
	SHARPNR_ACCOUNTING     BaseModule = "SHARPNR_ACCOUNTING"
	SHARPNR_HUMAN_RESOUCES BaseModule = "SHARPNR_HR"
	SHARPNR_MARKETING      BaseModule = "SHARPNR_MARKETING"
)

type HospitalityModule string

const (
	Hotels      HospitalityModule = "HOSPITALITY_HOTELS"
	Bars        HospitalityModule = "HOSPITALITY_BARS"
	Restaurants HospitalityModule = "HOSPITALITY_RESTAURANTS"
	Cafes       HospitalityModule = "HOSPITALITY_CAFES"
	Clubs       HospitalityModule = "HOSPITALITY_CLUBS"
	EventVenues HospitalityModule = "HOSPITALITY_EVENT_VENUES"
	Resorts     HospitalityModule = "HOSPITALITY_RESORTS"
)
