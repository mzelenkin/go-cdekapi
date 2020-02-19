package cdekapi

import "context"

type ApiService interface {
	GetOffices(ctx context.Context, filter OfficesFilter) (*Offices, error)
}
