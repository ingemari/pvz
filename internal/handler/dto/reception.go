package dto

type ReceptionRequst struct {
	PvzID string
}

type ReceptionResponse struct {
	Id       string
	DateTime string
	PvzID    string
	Status   string
}
