package delivery

import "wilayah/domain"

type InserFormat struct {
	KotaName   string `json:"kota_name" form:"kota_name"`
	ProvinsiID int    `json:"ProvinsiID" form:"ProvinsiID"`
}

func (i InserFormat) ToModel() domain.Kota {
	return domain.Kota{
		KotaName:   i.KotaName,
		ProvinsiID: i.ProvinsiID,
	}
}
