package utilitiesApi_service

import (
	"context"
	"strings"

	utilitiesApi_dto "sharpnr-engine/utilities_apis/dto"
)

// normalizeISO upper-cases and trims ISO codes so lookups are case-insensitive,
// matching how the upstream API accepts "in" or "IN".
func normalizeISO(code string) string {
	return strings.ToUpper(strings.TrimSpace(code))
}

func (s *utilitiesApiService) ListCountries(ctx context.Context) ([]utilitiesApi_dto.CountryResponse, error) {
	countries, err := s.repo.ListCountries(ctx)
	if err != nil {
		return nil, err
	}
	return utilitiesApi_dto.NewCountryResponses(countries), nil
}

func (s *utilitiesApiService) GetCountry(ctx context.Context, countryISO2 string) (*utilitiesApi_dto.CountryResponse, error) {
	country, err := s.repo.GetCountryByISO2(ctx, normalizeISO(countryISO2))
	if err != nil {
		return nil, err
	}
	resp := utilitiesApi_dto.NewCountryResponse(*country)
	return &resp, nil
}

func (s *utilitiesApiService) ListStates(ctx context.Context, countryISO2 string) ([]utilitiesApi_dto.StateResponse, error) {
	states, err := s.repo.ListStatesByCountry(ctx, normalizeISO(countryISO2))
	if err != nil {
		return nil, err
	}
	return utilitiesApi_dto.NewStateResponses(states), nil
}

func (s *utilitiesApiService) GetState(ctx context.Context, countryISO2, stateISO2 string) (*utilitiesApi_dto.StateResponse, error) {
	state, err := s.repo.GetStateByISO2(ctx, normalizeISO(countryISO2), normalizeISO(stateISO2))
	if err != nil {
		return nil, err
	}
	resp := utilitiesApi_dto.NewStateResponse(*state)
	return &resp, nil
}

func (s *utilitiesApiService) ListCities(ctx context.Context, countryISO2, stateISO2 string) ([]utilitiesApi_dto.CityResponse, error) {
	cities, err := s.repo.ListCitiesByState(ctx, normalizeISO(countryISO2), normalizeISO(stateISO2))
	if err != nil {
		return nil, err
	}
	return utilitiesApi_dto.NewCityResponses(cities), nil
}
