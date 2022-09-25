package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Kota struct {
	gorm.Model
	KotaName   string
	ProvinsiID int
	Provinsi   Provinsi `gorm:"foreignKey:ProvinsiID; references:ID; constraint:OnDelete:CASCADE"`
}

type Provinsi struct {
	gorm.Model
	ProvinsiName string
	ProvinsiCode string
}

func (s *Kota) ToModel() domain.Kota {
	return domain.Kota{
		ID:         int(s.ID),
		KotaName:   s.KotaName,
		ProvinsiID: s.ProvinsiID,
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
	res.ProvinsiID = data.ProvinsiID
	return res
}
