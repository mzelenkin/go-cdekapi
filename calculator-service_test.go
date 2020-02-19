package cdekapi

import (
	"context"
	//"reflect"
	"testing"
)

func TestClient_Calculator(t *testing.T) {
	c := New(APITestBaseURL).SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")
	data := CalculatorRequest{
		SenderCityId:   "44",
		ReceiverCityID: "44",
		Currency:       "RUB",
		TariffID:       1,
		Goods: []Good{
			{Weight: 0.3, Length: 5, Width: 20, Height: 10},
		},
		Services: []ServiceReq(nil),
	}
	result, err := c.Calculator(context.Background(), &data)

	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("#%#v\n", result)
	if result.Currency != "RUB" || result.PriceByCurrency != 380 {
		t.Fail()
	}
}
