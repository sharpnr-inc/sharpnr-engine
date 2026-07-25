package utlitiesApi_repo

import (
	"context"

	utilitiesApi_models "sharpnr-engine/utilities_apis/models"
)

type UtilitesAPIRepositoryInterface interface {
	// Geo reference data (countries / states / cities), sourced from the
	// self-hosted countrystatecity dataset in the `geo` schema.
	ListCountries(ctx context.Context) ([]utilitiesApi_models.Country, error)
	GetCountryByISO2(ctx context.Context, iso2 string) (*utilitiesApi_models.Country, error)
	ListStatesByCountry(ctx context.Context, countryISO2 string) ([]utilitiesApi_models.State, error)
	GetStateByISO2(ctx context.Context, countryISO2, stateISO2 string) (*utilitiesApi_models.State, error)
	ListCitiesByState(ctx context.Context, countryISO2, stateISO2 string) ([]utilitiesApi_models.City, error)
}
