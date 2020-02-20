package cdekapi

import (
	"context"
	"errors"
	"fmt"
)

func (c Client) GetCities(ctx context.Context, filter CitiesFilter) ([]City, error) {
	resp, err := c.getRequest(ctx).
		SetResult(&[]City{}).
		SetQueryParams(filter.BuildQueryParams()).
		Get(apiListOfCities)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 399 {
		return nil, fmt.Errorf("http error: %s", resp.Status())
	}

	if result, ok := resp.Result().(*[]City); ok {
		return *result, nil
	}

	return nil, errors.New("can't convert response to array of offices")
}
