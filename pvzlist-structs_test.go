package cdekapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestOfficesItemUnmarshal(t *testing.T) {
	jsonItem := `{
         "code": "MSK40",
         "postalCode": "119361",
         "name": "На Юго-Западной",
         "countryCode": "1",
         "countryCodeIso": "RU",
         "countryName": "Россия",
         "regionCode": "81",
         "regionName": "Москва",
         "cityCode": "44",
         "city": "Москва",
         "workTime": "Пн-Пт 10:00-20:00, Сб 10:00-16:00, Вс 10:00-16:00",
         "address": "ул. Марии Поливановой, 9",
         "fullAddress": "Россия, Москва, Москва, ул. Марии Поливановой, 9",
         "phone": "+74954304800",
         "note": "",
         "coordX": "37.4532166",
         "coordY": "55.6786919",
         "type": "PVZ",
         "ownerCode": "cdek",
         "isDressingRoom": true,
         "haveCashless": true,
         "allowedCod": true,
         "nearestStation": "",
         "metroStation": "Юго-Западная",
         "site": "https://www.cdek.ru/contacts/g_moskva_ul_marii_polivanovoy_d9_st_m_yugo-zapadna.html",
         "email": "m.lukasheva@cdek.ru",
         "addressComment": "Метро юго-Западная(Сокольническая линия), выход последний вагон из центра. Из метро прямо на автобусную остановку. Автобусы 630.226. Ехать до остановки \"Озерная17\". Выйдя из автобуса, двигайтесь прямо против движения автотранспорта. На первом пешеходном п",
         "weightLimit": {
            "weightMin": "0",
            "weightMax": "30"
         },
         "workTimeYList": [
            {
               "day": "1",
               "periods": "10:00/20:00"
            },
            {
               "day": "2",
               "periods": "10:00/20:00"
            },
            {
               "day": "3",
               "periods": "10:00/20:00"
            },
            {
               "day": "4",
               "periods": "10:00/20:00"
            },
            {
               "day": "5",
               "periods": "10:00/20:00"
            },
            {
               "day": "6",
               "periods": "10:00/16:00"
            },
            {
               "day": "7",
               "periods": "10:00/16:00"
            }
         ],
         "phoneDetailList": [
            {
               "number": "+74954304800"
            }
         ]
      }`

	// То, что должно получиться %)
	ethalon := Office{
		Code:           "MSK40",
		PostalCode:     "119361",
		Name:           "На Юго-Западной",
		CountryCode:    1,
		CountryCodeIso: "RU",
		CountryName:    "Россия",
		RegionCode:     81,
		RegionName:     "Москва",
		CityCode:       44,
		City:           "Москва",
		WorkTime:       "Пн-Пт 10:00-20:00, Сб 10:00-16:00, Вс 10:00-16:00",
		Address:        "ул. Марии Поливановой, 9",
		FullAddress:    "Россия, Москва, Москва, ул. Марии Поливановой, 9",
		AddressComment: "Метро юго-Западная(Сокольническая линия), выход последний вагон из центра. Из метро прямо на автобусную остановку. Автобусы 630.226. Ехать до остановки \"Озерная17\". Выйдя из автобуса, двигайтесь прямо против движения автотранспорта. На первом пешеходном п",
		Phone:          "+74954304800",
		Email:          "m.lukasheva@cdek.ru",
		Note:           "",
		CoordX:         37.4532166,
		CoordY:         55.6786919,
		Type:           "PVZ",
		OwnerCode:      "cdek",
		IsDressingRoom: true,
		HaveCashless:   true,
		AllowedCod:     true,
		NearestStation: "",
		MetroStation:   "Юго-Западная",
		Site:           "https://www.cdek.ru/contacts/g_moskva_ul_marii_polivanovoy_d9_st_m_yugo-zapadna.html",
		OfficeImage:    []OfficeImage(nil),
		WorkTimeByDay:  []WorkTimeY(nil), WeightLimit: WeightLimit{WeightMin: 0, WeightMax: 30}, PhoneDetailList: []PhoneDetail{PhoneDetail{Number: "+74954304800"}}}
	// Объявляем структуру в которую будет происходить размаршаливание XML
	item := Office{}

	err := json.Unmarshal([]byte(jsonItem), &item)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", item)
	// Сверка результатов
	if reflect.DeepEqual(ethalon, item) != true {
		t.Fail()
	}
}
