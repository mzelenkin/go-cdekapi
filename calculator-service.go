package cdekapi

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (c Client) Calculator(ctx context.Context, request *CalculatorRequest) (CalculatorResult, error) {

	var ret CalculatorResult

	request = c.calculatorPrepare(request)
	resp, err := c.getRequest(ctx).
		SetResult(&CalculatorResponse{}).
		SetError(&ErrorsApiV1{}).
		SetBody(request).
		Post(APICalculatorURL)

	if err != nil {
		return ret, err
	}

	if resp.StatusCode() > 399 {
		return ret, fmt.Errorf("http error: %s", resp.Status())
	}

	if result, ok := resp.Result().(*CalculatorResponse); ok {
		if len(result.Errors) > 0 {
			return ret, errors.New(result.Errors[0].Text)
		}

		return result.Result, nil
	}

	return ret, errors.New("can't convert response to calculator response struct")
}

func (c Client) calculatorPrepare(request *CalculatorRequest) *CalculatorRequest {
	request.Version = "1.0"
	request.AuthLogin = &c.auth.ClientID
	request.Secure = &c.auth.Secret

	dateExecute := time.Now().Format("2006-01-02")
	request.DateExecute = &dateExecute

	return request
}
