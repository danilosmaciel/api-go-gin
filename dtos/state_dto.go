package dtos

type StateDTO struct {
	Code   int32     `json:"code"`
	Sigla  string    `json:"sigla"`
	Cities []CityDTO `json:"cities"`
}
