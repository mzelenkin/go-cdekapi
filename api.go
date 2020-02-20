package cdekapi

import "context"

type ApiService interface {
	GetOffices(ctx context.Context, filter OfficesFilter) ([]Office, error)
	GetCities(ctx context.Context, filter CitiesFilter) ([]City, error)
	Calculator(ctx context.Context, request CalculatorRequest) (CalculatorResult, error)
}
