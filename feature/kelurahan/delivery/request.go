package delivery

import "wilayah/domain"

type InserFormat struct {
	KelurahanName string `json:"kelurahan_name" form:"kelurahan_name"`
	KecamatanID   string `json:"kecamatan_id" form:"kecamatan_id"`
}

func (i InserFormat) ToModel() domain.Kelurahan {
	return domain.Kelurahan{
		KelurahanName: i.KelurahanName,
		KecamatanID:   i.KecamatanID,
	}
}
