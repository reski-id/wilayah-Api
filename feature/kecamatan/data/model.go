package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Kecamatan struct {
	gorm.Model
	KecamatanName string
	KotaID        string
}

type Kota struct {
	gorm.Model
	KotaName     string
	ProvinsiCode string
}

func (s *Kecamatan) ToModel() domain.Kecamatan {
	return domain.Kecamatan{
		ID:            int(s.ID),
		KecamatanName: s.KecamatanName,
		KotaID:        s.KotaID,
	}
}

func ParseToArr(arr []Kecamatan) []domain.Kecamatan {
	var res []domain.Kecamatan

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Kecamatan) Kecamatan {
	var res Kecamatan
	res.KecamatanName = data.KecamatanName
	res.KotaID = data.KotaID
	return res
}
