package cdekapi

type CitiesFilterParamName string

const (
	// Array of country codes in the format ISO_3166-1_alpha-2
	CitiesFilterCountryCodes CitiesFilterParamName = "country_codes"

	// Region code in the CDEK IS
	CitiesFilterRegionCode CitiesFilterParamName = "region_code"

	// Region code according to the Russian Classifier of Countries of the World
	CitiesFilterKladrRegionCode CitiesFilterParamName = "kladr_region_code"

	// Region code according to the Federal Information Address System
	CitiesFilterFiasRegionCode CitiesFilterParamName = "fias_region_guid"

	// City code according to the Russian Classifier of Countries of the World
	CitiesFilterKladrCityCode CitiesFilterParamName = "kladr_code"

	// City code according to the Federal Information Address System
	CitiesFilterFiasCityCode CitiesFilterParamName = "fias_guid"

	// 	City code according to the Federal Information Address System, UUID
	CitiesFilterPostalCode CitiesFilterParamName = "postal_code"

	// City code from the CDEK IS
	CitiesFilterCode CitiesFilterParamName = "code"

	// City name
	CitiesFilterCity CitiesFilterParamName = "city"

	// Limitation on the number of results displayed. Required, if "page" is specified. Default value: 1,000
	CitiesFilterSize CitiesFilterParamName = "size"

	// Number of the results page. Default value: 0
	CitiesFilterPage CitiesFilterParamName = "page"

	// Localization. Default: “rus”
	CitiesFilterLang CitiesFilterParamName = "lang"

	// Cash-on-delivery amount limit, possible values: 1 – no limit; 0 – cash on delivery is not accepted;
	// positive value – the cash-on-delivery amount does not exceed this value.
	CitiesFilterPaymentLimit CitiesFilterParamName = "payment_limit"
)

type CitiesFilter struct {
	params map[string]string
}

func (f *CitiesFilter) AddParam(name CitiesFilterParamName, value string) *CitiesFilter {
	if f.params == nil {
		f.params = map[string]string{}
	}

	paramName := string(name)
	f.params[paramName] = value

	return f
}

func (f CitiesFilter) BuildQueryParams() map[string]string {
	return f.params
}
