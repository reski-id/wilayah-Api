package delivery

import "wilayah/domain"

type ProvinsiResponse struct {
	ID           int    `json:"id"`
	ProvinsiName string `json:"provinsi_name"`
	ProvinsiCode string `json:"provinsi_code"`
}

func FromModel(data domain.Provinsi) ProvinsiResponse {
	var res ProvinsiResponse
	res.ID = data.ID
	res.ProvinsiName = data.ProvinsiName
	res.ProvinsiCode = data.ProvinsiCode
	return res
}

func FromModelToList(data []domain.Provinsi) []ProvinsiResponse {
	result := []ProvinsiResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
