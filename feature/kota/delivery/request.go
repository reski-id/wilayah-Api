package delivery

import "wilayah/domain"

type InserFormat struct {
	KotaName     string `json:"kota_name" form:"kota_name"`
	ProvinsiCode string `json:"provinsi_code" form:"provinsi_code"`
}

func (i InserFormat) ToModel() domain.Kota {
	return domain.Kota{
		KotaName:     i.KotaName,
		ProvinsiCode: i.ProvinsiCode,
	}
}
