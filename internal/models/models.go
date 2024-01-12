package models

type LocationRequest struct {
	ReqID string `json:"reqId"`
	IP    string `json:"ip"`
}

type LocationResponse struct {
	// ReqID 		string `json:"reqId"`
	CountryCode string  `json:"countryCode"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}