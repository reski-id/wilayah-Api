package delivery

import "wilayah/domain"

type KotaResponse struct {
	ID           int    `json:"id"`
	KotaName     string `json:"kota_name"`
	ProvinsiCode string `json:"provinsi_code"`
}

func FromModel(data domain.Kota) KotaResponse {
	var res KotaResponse
	res.ID = data.ID
	res.KotaName = data.KotaName
	res.ProvinsiCode = data.ProvinsiCode
	return res
}

func FromModelToList(data []domain.Kota) []KotaResponse {
	result := []KotaResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
