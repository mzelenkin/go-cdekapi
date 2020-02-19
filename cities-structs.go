package cdekapi

// City is response List of Cities request
type City struct {
	Code            string      `json:"code"`
	City            string      `json:"city"`
	FiasGUID        string      `json:"fias_guid"`
	KladrCode       string      `json:"kladr_code"`
	CountryCode     string      `json:"country_code"`
	Country         string      `json:"country"`
	Region          string      `json:"region"`
	RegionCode      string      `json:"region_code"`
	FiasRegionGUID  string      `json:"fias_region_guid"`
	KladrRegionCode string      `json:"kladr_region_code"`
	SubRegion       string      `json:"sub_region"`
	PostalCodes     []string    `json:"postal_codes"`
	Longitude       float64     `json:"longitude"`
	Latitude        float64     `json:"latitude"`
	TimeZone        string      `json:"time_zone"`
	PaymentLimit    float64     `json:"payment_limit"`
	Errors          []CityError `json:"errors"`
}

type CityError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
