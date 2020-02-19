package cdekapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCalculatorStruct(t *testing.T) {
	response := []byte(`{
    "result": {
        "price": "2555",
        "deliveryPeriodMin": 2,
        "deliveryPeriodMax": 3,
        "deliveryDateMin": "2018-07-28",
        "deliveryDateMax": "2018-07-29",
        "tariffId": "137",
        "cashOnDelivery": "30000.00",
        "percentVAT": 18,
        "priceByCurrency": 2555,
        "currency": "RUB",
        "services": [
            {
                "id": 2,
                "title": "Insurance",
                "price": 15
            },
            {
                "id": 30,
                "title": "Home fitting",
                "price": 0
            }
        ]
    }
}`)
	result := CalculatorResponse{}
	err := json.Unmarshal(response, &result)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v\n", result)
}
