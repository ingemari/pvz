package dto

type PvzCreateRequest struct {
	City string `json:"city"`
}

type PvzCreateResponse struct {
	Id      string `json:"id"`
	RegDate string `json:"registrationDate"`
	City    string `json:"city"`
}
