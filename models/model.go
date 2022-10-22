package models

type weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusResponse struct {
	Status weather `json:"status"`
}
