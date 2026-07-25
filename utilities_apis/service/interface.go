package utilitiesApi_service

import (
	"context"

	utilitiesApi_dto "sharpnr-engine/utilities_apis/dto"
)

type UtilitiesAPIServiceInterface interface {
	// Geo lookups returning DTOs shaped exactly like the countrystatecity.in API.
	ListCountries(ctx context.Context) ([]utilitiesApi_dto.CountryResponse, error)
	GetCountry(ctx context.Context, countryISO2 string) (*utilitiesApi_dto.CountryResponse, error)
	ListStates(ctx context.Context, countryISO2 string) ([]utilitiesApi_dto.StateResponse, error)
	GetState(ctx context.Context, countryISO2, stateISO2 string) (*utilitiesApi_dto.StateResponse, error)
	ListCities(ctx context.Context, countryISO2, stateISO2 string) ([]utilitiesApi_dto.CityResponse, error)
}
