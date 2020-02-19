package cdekapi

//Good Package dimensions
type Good struct {
	Weight float64 `json:"weight"`
	Length int     `json:"length,omitempty"`
	Width  int     `json:"width,omitempty"`
	Height int     `json:"height,omitempty"`
	Volume float64 `json:"volume,omitempty"`
}

//ServiceReq List of additional service
type ServiceReq struct {
	ID    int `json:"id"`
	Param int `json:"param,omitempty"`
}

type CalculatorRequest struct {
	Version     string  `json:"version"`
	AuthLogin   *string `json:"authLogin,omitempty"`
	Secure      *string `json:"secure,omitempty"`
	DateExecute *string `json:"dateExecute,omitempty"`
	Lang        string  `json:"lang"`

	SenderCityId   string       `json:"senderCityId"`
	ReceiverCityID string       `json:"receiverCityId"`
	Currency       string       `json:"currency,omitempty"`
	TariffID       int          `json:"tariffId"`
	Goods          []Good       `json:"goods"`
	Services       []ServiceReq `json:"services,omitempty"`
}

type CalculatorResponse struct {
	ErrorsApiV1
	Result CalculatorResult
}

type CalculatorResult struct {
	Price             float64 `json:"price,string"`
	DeliveryPeriodMin int     `json:"deliveryPeriodMin"`
	DeliveryPeriodMax int     `json:"deliveryPeriodMax"`
	DeliveryDateMin   string  `json:"deliveryDateMin"`
	DeliveryDateMax   string  `json:"deliveryDateMax"`
	TariffID          int     `json:"tariffId"`
	CashOnDelivery    float64 `json:"cashOnDelivery"`
	PriceByCurrency   float64 `json:"priceByCurrency"`
	Currency          string  `json:"currency"`
	PercentVAT        int     `json:"percentVAT"`
	Services          []AdditionalService
}

type AdditionalService struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Rate  float64 `json:"rate"`
}
