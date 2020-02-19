package cdekapi

type Offices struct {
	Offices []Office `json:"pvz"`
}

type Office struct {
	Code            string        `json:"code"`
	PostalCode      string        `json:"postalCode"`
	Name            string        `json:"name"`
	CountryCode     int           `json:"countryCode,string"`
	CountryCodeIso  string        `json:"countryCodeIso"`
	CountryName     string        `json:"countryName"`
	RegionCode      int           `json:"regionCode,string"`
	RegionName      string        `json:"regionName"`
	CityCode        int           `json:"cityCode,string"`
	City            string        `json:"city"`
	WorkTime        string        `json:"workTime"`
	Address         string        `json:"address"`
	FullAddress     string        `json:"fullAddress"`
	AddressComment  string        `json:"addressComment"`
	Phone           string        `json:"phone"`
	Email           string        `json:"email"`
	Note            string        `json:"note"`
	CoordX          float64       `json:"coordX,string"`
	CoordY          float64       `json:"coordY,string"`
	Type            string        `json:"type"`
	OwnerCode       string        `json:"ownerCode"`
	IsDressingRoom  bool          `json:"isDressingRoom"`
	HaveCashless    bool          `json:"haveCashless"`
	AllowedCod      bool          `json:"allowedCod"`
	NearestStation  string        `json:"nearestStation"`
	MetroStation    string        `json:"metroStation"`
	Site            string        `json:"site"`
	OfficeImage     []OfficeImage `json:"officeImageList"`
	WorkTimeByDay   []WorkTimeY   `json:"workTimeY"`
	WeightLimit     WeightLimit   `json:"weightLimit"`
	PhoneDetailList []PhoneDetail `json:"phoneDetailList"`
}

// OfficeImage contains all photos of the office (except for a photo showing how to get to it)
type OfficeImage struct {
	Number int
	URL    string `json:"url"`
}

// WorkTimeY opening hours for every day
type WorkTimeY struct {
	Day     int    `json:"day"`
	Periods string `json:"periods"`
}

// WeightLimit is weight limits for a pickup point(the tag is used only if limits are set)
type WeightLimit struct {
	WeightMin float64 `json:"weightMin,string"`
	WeightMax float64 `json:"weightMax,string"`
}

type PhoneDetail struct {
	Number string `json:"number"`
}
