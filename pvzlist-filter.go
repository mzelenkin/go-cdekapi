package cdekapi

type OfficesFilterParamName string

const (
	//OfficesFilterCityPostCode Почтовый индекс города, для которого необходим список ПВЗ
	OfficesFilterCityPostCode OfficesFilterParamName = "citypostcode"

	//OfficesFilterCityID Код города по базе СДЭК
	OfficesFilterCityID OfficesFilterParamName = "cityid"

	//OfficesFilterType Тип пункта выдачи, по умолчанию «PVZ».
	OfficesFilterType OfficesFilterParamName = "type"

	//OfficesFilterCountryID Код страны по базе СДЭК
	OfficesFilterCountryID OfficesFilterParamName = "countryid"

	//OfficesFilterCountryIso Код страны в формате ISO_3166-1_alpha-2
	OfficesFilterCountryIso OfficesFilterParamName = "countryiso"

	//OfficesFilterRegionID Код региона по базе СДЭК
	OfficesFilterRegionID OfficesFilterParamName = "regionid"

	//OfficesFilterHaveCashless Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	OfficesFilterHaveCashless OfficesFilterParamName = "havecashless"

	//OfficesFilterAllowedCod Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	OfficesFilterAllowedCod OfficesFilterParamName = "allowedcod"

	//OfficesFilterIsDressingRoom Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	OfficesFilterIsDressingRoom OfficesFilterParamName = "isdressingroom"

	//OfficesFilterWeightMax Максимальный вес, который может принять ПВЗ
	OfficesFilterWeightMax OfficesFilterParamName = "weightmax"

	//OfficesFilterLang Локализация ПВЗ. По-умолчанию "rus"
	OfficesFilterLang OfficesFilterParamName = "lang"

	//OfficesFilterTakeOnly Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
	OfficesFilterTakeOnly OfficesFilterParamName = "takeonly"
)

type OfficesFilter struct {
	params map[string]string
}

func (f *OfficesFilter) AddParam(name OfficesFilterParamName, value string) *OfficesFilter {
	if f.params == nil {
		f.params = map[string]string{}
	}

	paramName := string(name)
	f.params[paramName] = value

	return f
}

func (f OfficesFilter) BuildQueryParams() map[string]string {
	return f.params
}
