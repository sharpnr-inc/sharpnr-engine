package utilitiesApi_dto

import (
	"strconv"

	utilitiesApi_models "sharpnr-engine/utilities_apis/models"
)

// These DTOs intentionally mirror, field-for-field, the JSON shapes returned by
// the countrystatecity.in REST API (https://api.countrystatecity.in/v1) so the
// frontend can swap the base URL and keep working unchanged.
//
// Notes on parity with the upstream API:
//   - latitude/longitude are serialized as STRINGS with 8 decimal places
//     (e.g. "19.07283000"), exactly like the upstream API.
//   - nullable fields use pointers and are emitted as JSON `null` (no omitempty)
//     because upstream returns explicit nulls (e.g. "type": null on cities).
//   - timezones/translations on a country are raw JSON-encoded strings, matching
//     upstream (they are stored that way in the DB and passed through verbatim).

// CountryResponse matches GET /v1/countries and GET /v1/countries/{ciso}.
type CountryResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	ISO2             *string  `json:"iso2"`
	ISO3             *string  `json:"iso3"`
	NumericCode      *string  `json:"numeric_code"`
	PhoneCode        *string  `json:"phonecode"`
	Capital          *string  `json:"capital"`
	Currency         *string  `json:"currency"`
	CurrencyName     *string  `json:"currency_name"`
	CurrencySymbol   *string  `json:"currency_symbol"`
	TLD              *string  `json:"tld"`
	Native           *string  `json:"native"`
	Region           *string  `json:"region"`
	RegionID         *int64   `json:"region_id"`
	Subregion        *string  `json:"subregion"`
	SubregionID      *int64   `json:"subregion_id"`
	Nationality      *string  `json:"nationality"`
	Timezones        *string  `json:"timezones"`
	Translations     *string  `json:"translations"`
	Latitude         *string  `json:"latitude"`
	Longitude        *string  `json:"longitude"`
	Emoji            *string  `json:"emoji"`
	EmojiU           *string  `json:"emojiU"`
	Population       *int64   `json:"population"`
	GDP              *int64   `json:"gdp"`
	AreaSqKm         *float64 `json:"area_sq_km"`
	PostalCodeFormat *string  `json:"postal_code_format"`
	PostalCodeRegex  *string  `json:"postal_code_regex"`
	WikiDataID       *string  `json:"wikiDataId"`
}

// StateResponse matches GET /v1/countries/{ciso}/states and
// GET /v1/countries/{ciso}/states/{siso}.
type StateResponse struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ISO2         *string `json:"iso2"`
	CountryID    int64   `json:"country_id"`
	CountryCode  string  `json:"country_code"`
	Latitude     *string `json:"latitude"`
	Longitude    *string `json:"longitude"`
	Timezone     *string `json:"timezone"`
	FipsCode     *string `json:"fips_code"`
	ISO31662     *string `json:"iso3166_2"`
	Type         *string `json:"type"`
	Level        *int    `json:"level"`
	ParentID     *int64  `json:"parent_id"`
	Native       *string `json:"native"`
	Population   *string `json:"population"`
	Translations *string `json:"translations"`
	WikiDataID   *string `json:"wikiDataId"`
}

// CityResponse matches GET /v1/countries/{ciso}/states/{siso}/cities.
type CityResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	StateID     int64   `json:"state_id"`
	StateCode   string  `json:"state_code"`
	CountryID   int64   `json:"country_id"`
	CountryCode string  `json:"country_code"`
	Latitude    string  `json:"latitude"`
	Longitude   string  `json:"longitude"`
	Timezone    *string `json:"timezone"`
	Population  *int64  `json:"population"`
	Type        *string `json:"type"`
	Level       *int    `json:"level"`
	ParentID    *int64  `json:"parent_id"`
	Native      *string `json:"native"`
	WikiDataID  *string `json:"wikiDataId"`
}

// coordStr formats a nullable coordinate to the upstream "%.8f" string form,
// preserving JSON null when the value is absent.
func coordStr(v *float64) *string {
	if v == nil {
		return nil
	}
	s := strconv.FormatFloat(*v, 'f', 8, 64)
	return &s
}

// coordStrVal formats a non-nullable coordinate (cities always have one).
func coordStrVal(v float64) string {
	return strconv.FormatFloat(v, 'f', 8, 64)
}

// NewCountryResponse maps a Country model to its API DTO.
func NewCountryResponse(m utilitiesApi_models.Country) CountryResponse {
	return CountryResponse{
		ID:               m.ID,
		Name:             m.Name,
		ISO2:             m.ISO2,
		ISO3:             m.ISO3,
		NumericCode:      m.NumericCode,
		PhoneCode:        m.PhoneCode,
		Capital:          m.Capital,
		Currency:         m.Currency,
		CurrencyName:     m.CurrencyName,
		CurrencySymbol:   m.CurrencySymbol,
		TLD:              m.TLD,
		Native:           m.Native,
		Region:           m.Region,
		RegionID:         m.RegionID,
		Subregion:        m.Subregion,
		SubregionID:      m.SubregionID,
		Nationality:      m.Nationality,
		Timezones:        m.Timezones,
		Translations:     m.Translations,
		Latitude:         coordStr(m.Latitude),
		Longitude:        coordStr(m.Longitude),
		Emoji:            m.Emoji,
		EmojiU:           m.EmojiU,
		Population:       m.Population,
		GDP:              m.GDP,
		AreaSqKm:         m.AreaSqKm,
		PostalCodeFormat: m.PostalCodeFormat,
		PostalCodeRegex:  m.PostalCodeRegex,
		WikiDataID:       m.WikiDataID,
	}
}

// NewCountryResponses maps a slice of Country models, always returning a
// non-nil slice so the JSON body is `[]` rather than `null` when empty.
func NewCountryResponses(rows []utilitiesApi_models.Country) []CountryResponse {
	out := make([]CountryResponse, 0, len(rows))
	for _, m := range rows {
		out = append(out, NewCountryResponse(m))
	}
	return out
}

// NewStateResponse maps a State model to its API DTO.
func NewStateResponse(m utilitiesApi_models.State) StateResponse {
	return StateResponse{
		ID:           m.ID,
		Name:         m.Name,
		ISO2:         m.ISO2,
		CountryID:    m.CountryID,
		CountryCode:  m.CountryCode,
		Latitude:     coordStr(m.Latitude),
		Longitude:    coordStr(m.Longitude),
		Timezone:     m.Timezone,
		FipsCode:     m.FipsCode,
		ISO31662:     m.ISO31662,
		Type:         m.Type,
		Level:        m.Level,
		ParentID:     m.ParentID,
		Native:       m.Native,
		Population:   m.Population,
		Translations: m.Translations,
		WikiDataID:   m.WikiDataID,
	}
}

// NewStateResponses maps a slice of State models (never nil).
func NewStateResponses(rows []utilitiesApi_models.State) []StateResponse {
	out := make([]StateResponse, 0, len(rows))
	for _, m := range rows {
		out = append(out, NewStateResponse(m))
	}
	return out
}

// NewCityResponse maps a City model to its API DTO.
func NewCityResponse(m utilitiesApi_models.City) CityResponse {
	return CityResponse{
		ID:          m.ID,
		Name:        m.Name,
		StateID:     m.StateID,
		StateCode:   m.StateCode,
		CountryID:   m.CountryID,
		CountryCode: m.CountryCode,
		Latitude:    coordStrVal(m.Latitude),
		Longitude:   coordStrVal(m.Longitude),
		Timezone:    m.Timezone,
		Population:  m.Population,
		Type:        m.Type,
		Level:       m.Level,
		ParentID:    m.ParentID,
		Native:      m.Native,
		WikiDataID:  m.WikiDataID,
	}
}

// NewCityResponses maps a slice of City models (never nil).
func NewCityResponses(rows []utilitiesApi_models.City) []CityResponse {
	out := make([]CityResponse, 0, len(rows))
	for _, m := range rows {
		out = append(out, NewCityResponse(m))
	}
	return out
}
