package delivery

import "wilayah/domain"

type InserFormat struct {
	KecamatanName string `json:"kecamatan_name" form:"kecamatan_name"`
	KotaID        int    `json:"kota_id" form:"kota_id"`
}

func (i InserFormat) ToModel() domain.Kecamatan {
	return domain.Kecamatan{
		KecamatanName: i.KecamatanName,
		KotaID:        i.KotaID,
	}
}
