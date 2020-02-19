package cdekapi

type ErrorsApiV1 struct {
	Errors []ErrorApiV1 `json:"error"`
}

type ErrorApiV1 struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
