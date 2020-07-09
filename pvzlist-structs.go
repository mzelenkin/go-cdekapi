package cdekapi

type Offices struct {
	Offices []Office
}

type Office struct {
	Code               string           `json:"code"`
	Name               string           `json:"name"`
	Location           OfficeLocation   `json:"location"`
	AddressComment     string           `json:"address_comment"`
	NearestStation     string           `json:"nearest_station"`
	MetroStation       string           `json:"nearest_metro_station"`
	WorkTime           string           `json:"work_time"`
	Phones             []PhoneDetail    `json:"phones"`
	Email              string           `json:"email"`
	Note               string           `json:"note"`
	Type               string           `json:"type"`
	OwnerCode          string           `json:"owner_сode"`
	TakeOnly           bool             `json:"take_only"`
	IsDressingRoom     bool             `json:"is_dressing_room"`
	HaveCashless       bool             `json:"have_cashless"`
	HaveCash           bool             `json:"have_cash"`
	AllowedCod         bool             `json:"allowed_cod"`
	Site               string           `json:"site"`
	OfficeImage        []OfficeImage    `json:"office_image_list"`
	WorkTimeList       []WorkTimeItem   `json:"work_time_list"`
	WorkTimeExceptions []WorkTimeExcept `json:"work_time_exceptions"`
	WeightMin          float64          `json:"weight_min"`
	WeightMax          float64          `json:"weight_max"`

	Errors []ErrorApiV2
}

type OfficeLocation struct {
	CountryCode string  `json:"country_code"`
	RegionCode  int     `json:"region_code"`
	RegionName  string  `json:"region"`
	CityCode    int     `json:"city_code"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postal_сode"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"adress"`
	AddressFull string  `json:"address_full"`
}

// OfficeImage contains all photos of the office (except for a photo showing how to get to it)
type OfficeImage struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

// WorkTimeItem opening hours for every day
type WorkTimeItem struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

// WorkTimeExcept
type WorkTimeExcept struct {
	Date      string `json:"date"`
	Time      string `json:"time"`
	IsWorking bool   `json:"is_working"`
}

type PhoneDetail struct {
	Number     string `json:"number"`
	Additional string `json:"additional"`
}

type ErrorApiV2 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
