package cdekapi

type OfficesFilterParamName string

const (
	//OfficesFilterCityPostCode Почтовый индекс города, для которого необходим список ПВЗ
	OfficesFilterCityPostCode OfficesFilterParamName = "postal_code"

	//OfficesFilterCityID Код города по базе СДЭК
	OfficesFilterCityID OfficesFilterParamName = "city_code"

	//OfficesFilterType Тип пункта выдачи, по умолчанию «PVZ».
	OfficesFilterType OfficesFilterParamName = "type"

	//OfficesFilterCountryIso Код страны в формате ISO_3166-1_alpha-2
	OfficesFilterCountryIso OfficesFilterParamName = "country_code"

	//OfficesFilterRegionID Код региона по базе СДЭК
	OfficesFilterRegionID OfficesFilterParamName = "region_code"

	//OfficesFilterHaveCashless Наличие терминала оплаты («1», «true» - есть; «0», «false» - нет.)
	OfficesFilterHaveCashless OfficesFilterParamName = "have_cashless"

	//OfficesFilterHaveCash Есть прием наличных («1», «true» - есть; «0», «false» - нет.)
	OfficesFilterHaveCash OfficesFilterParamName = "have_cashless"

	//OfficesFilterAllowedCod Разрешен наложенный платеж («1», «true» - да; «0», «false» - нет.)
	OfficesFilterAllowedCod OfficesFilterParamName = "allowed_cod"

	//OfficesFilterIsDressingRoom Наличие примерочной («1», «true» - есть; «0», «false» - нет.)
	OfficesFilterIsDressingRoom OfficesFilterParamName = "is_dressing_room"

	//OfficesFilterWeightMin Минимальный вес в кг, который принимает ПВЗ.
	//При переданном значении будут выводиться ПВЗ с минимальным весом до указанного значения
	OfficesFilterWeightMin OfficesFilterParamName = "weight_min"

	//OfficesFilterWeightMax Максимальный вес, который может принять ПВЗ
	OfficesFilterWeightMax OfficesFilterParamName = "weight_max"

	//OfficesFilterLang Локализация ПВЗ. По-умолчанию "rus"
	OfficesFilterLang OfficesFilterParamName = "lang"

	//OfficesFilterTakeOnly Является ли ПВЗ только пунктом выдачи («1», «true» - да; «0», «false» - нет.)
	OfficesFilterTakeOnly OfficesFilterParamName = "take_only"
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
