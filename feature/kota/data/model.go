package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Kota struct {
	gorm.Model
	KotaName     string
	ProvinsiCode string
}

type Provinsi struct {
	gorm.Model
	ProvinsiName string
	ProvinsiCode string
}

func (s *Kota) ToModel() domain.Kota {
	return domain.Kota{
		ID:           int(s.ID),
		KotaName:     s.KotaName,
		ProvinsiCode: s.ProvinsiCode,
	}
}

func ParseToArr(arr []Kota) []domain.Kota {
	var res []domain.Kota

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Kota) Kota {
	var res Kota
	res.KotaName = data.KotaName
	res.ProvinsiCode = data.ProvinsiCode
	return res
}
