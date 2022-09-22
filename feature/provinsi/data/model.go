package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Provinsi struct {
	gorm.Model
	ProvinsiName string
	ProvinsiCode string
}

func (s *Provinsi) ToModel() domain.Provinsi {
	return domain.Provinsi{
		ID:           int(s.ID),
		ProvinsiName: s.ProvinsiName,
		ProvinsiCode: s.ProvinsiCode,
	}
}

func ParseToArr(arr []Provinsi) []domain.Provinsi {
	var res []domain.Provinsi

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Provinsi) Provinsi {
	var res Provinsi
	res.ProvinsiName = data.ProvinsiName
	res.ProvinsiCode = data.ProvinsiCode
	return res
}
