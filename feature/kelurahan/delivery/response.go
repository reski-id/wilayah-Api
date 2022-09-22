package delivery

import "wilayah/domain"

type KelurahanResponse struct {
	ID            int    `json:"id"`
	KelurahanName string `json:"kelurahan_name" form:"kelurahan_name"`
	KecamatanID   string `json:"kecamatan_id" form:"kecamatan_id"`
}

func FromModel(data domain.Kelurahan) KelurahanResponse {
	var res KelurahanResponse
	res.ID = data.ID
	res.KelurahanName = data.KelurahanName
	res.KecamatanID = data.KecamatanID
	return res
}

func FromModelToList(data []domain.Kelurahan) []KelurahanResponse {
	result := []KelurahanResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
