package cdekapi

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (c Client) Calculator(ctx context.Context, request *CalculatorRequest) (*CalculatorResult, error) {

	request = c.calculatorPrepare(request)
	resp, err := c.getRequest(ctx).
		SetResult(&CalculatorResponse{}).
		SetError(&ErrorsApiV1{}).
		SetBody(request).
		Post(APICalculatorURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 399 {
		return nil, fmt.Errorf("http error: %s", resp.Status())
	}

	if ret, ok := resp.Result().(*CalculatorResponse); ok {
		if len(ret.Errors) > 0 {
			return nil, errors.New(ret.Errors[0].Text)
		}

		return &ret.Result, nil
	}

	return nil, errors.New("can't convert response to calculator response struct")
}

func (c Client) calculatorPrepare(request *CalculatorRequest) *CalculatorRequest {
	request.Version = "1.0"
	request.AuthLogin = &c.auth.ClientID
	request.Secure = &c.auth.Secret

	dateExecute := time.Now().Format("2006-01-02")
	request.DateExecute = &dateExecute

	return request
}
