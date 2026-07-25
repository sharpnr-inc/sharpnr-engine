package utlitiesApi_repo

import (
	"context"

	utilitiesApi_models "sharpnr-engine/utilities_apis/models"
)

// ListCountries returns every country, ordered by name (matches the upstream
// alphabetical ordering used by countrystatecity.in).
func (r *utlitiesApiRepo) ListCountries(ctx context.Context) ([]utilitiesApi_models.Country, error) {
	var countries []utilitiesApi_models.Country
	err := r.db.WithContext(ctx).
		Order("name ASC").
		Find(&countries).Error
	return countries, err
}

// GetCountryByISO2 looks up a single country by its ISO-3166-1 alpha-2 code.
func (r *utlitiesApiRepo) GetCountryByISO2(ctx context.Context, iso2 string) (*utilitiesApi_models.Country, error) {
	var country utilitiesApi_models.Country
	err := r.db.WithContext(ctx).
		Where("iso2 = ?", iso2).
		First(&country).Error
	if err != nil {
		return nil, err
	}
	return &country, nil
}

// ListStatesByCountry returns all states of a country (by country alpha-2 code),
// using the denormalized country_code column on geo.states.
func (r *utlitiesApiRepo) ListStatesByCountry(ctx context.Context, countryISO2 string) ([]utilitiesApi_models.State, error) {
	var states []utilitiesApi_models.State
	err := r.db.WithContext(ctx).
		Where("country_code = ?", countryISO2).
		Order("name ASC").
		Find(&states).Error
	return states, err
}

// GetStateByISO2 looks up a single state by country alpha-2 + state ISO code.
func (r *utlitiesApiRepo) GetStateByISO2(ctx context.Context, countryISO2, stateISO2 string) (*utilitiesApi_models.State, error) {
	var state utilitiesApi_models.State
	err := r.db.WithContext(ctx).
		Where("country_code = ? AND iso2 = ?", countryISO2, stateISO2).
		First(&state).Error
	if err != nil {
		return nil, err
	}
	return &state, nil
}

// ListCitiesByState returns all cities of a state, keyed on the denormalized
// country_code + state_code columns on geo.cities.
func (r *utlitiesApiRepo) ListCitiesByState(ctx context.Context, countryISO2, stateISO2 string) ([]utilitiesApi_models.City, error) {
	var cities []utilitiesApi_models.City
	err := r.db.WithContext(ctx).
		Where("country_code = ? AND state_code = ?", countryISO2, stateISO2).
		Order("name ASC").
		Find(&cities).Error
	return cities, err
}
