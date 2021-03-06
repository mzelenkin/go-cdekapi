package cdekapi

import (
	"context"
	"errors"
	"fmt"
)

func (c Client) GetOffices(ctx context.Context, filter OfficesFilter) ([]Office, error) {
	resp, err := c.getRequest(ctx).
		SetResult(&[]Office{}).
		SetQueryParams(filter.BuildQueryParams()).
		Get(apiOfficesList)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 399 {
		return nil, fmt.Errorf("http error: %s", resp.Status())
	}

	if ret, ok := resp.Result().(*[]Office); ok {
		return *ret, nil
	}

	return nil, errors.New("can't convert response to array of offices")
}
