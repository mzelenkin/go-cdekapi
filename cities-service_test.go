package cdekapi

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_GetCities(t *testing.T) {
	filters := CitiesFilter{}
	filters.AddParam(CitiesFilterFiasCityCode, "7e42f157-87d6-4111-a406-17432e814723")
	//filters.AddParam(CitiesFilterCity, "Нижний Новгород")

	c := New(APITestBaseURL).SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")
	result, err := c.GetCities(context.Background(), filters)

	want := []City{
		City{
			Code:            "1002",
			City:            "Арзамас",
			FiasGUID:        "7e42f157-87d6-4111-a406-17432e814723",
			KladrCode:       "52",
			CountryCode:     "RU",
			Country:         "Россия",
			Region:          "Нижегородская",
			RegionCode:      "67",
			FiasRegionGUID:  "88cd27e2-6a8a-4421-9718-719a28a0a088",
			KladrRegionCode: "52",
			SubRegion:       "Арзамас",
			PostalCodes:     []string{"607220", "607221", "607222", "607223", "607224", "607225", "607226", "607227", "607228", "607230", "607232", "607233", "607278", "607279", "607280"},
			Longitude:       43.8157,
			Latitude:        55.3867,
			TimeZone:        "Europe/Moscow",
			PaymentLimit:    -1,
			Errors:          []CityError(nil),
		}}

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(result, want) == false {
		t.Errorf("\nRESU: %#v\nWANT: %#v", result, &want)
		t.Fail()
	}
}
