package delivery

import "wilayah/domain"

type KecamatanResponse struct {
	ID            int    `json:"id"`
	KecamatanName string `json:"kecamatan_name"`
	KotaID        int    `json:"kota_id"`
}

func FromModel(data domain.Kecamatan) KecamatanResponse {
	var res KecamatanResponse
	res.ID = data.ID
	res.KecamatanName = data.KecamatanName
	res.KotaID = data.KotaID
	return res
}

func FromModelToList(data []domain.Kecamatan) []KecamatanResponse {
	result := []KecamatanResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
