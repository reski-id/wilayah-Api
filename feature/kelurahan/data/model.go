package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Kelurahan struct {
	gorm.Model
	KelurahanName string
	KecamatanID   int
	Kecamatan     Kecamatan `gorm:"foreignKey:KecamatanID; references:ID; constraint:OnDelete:CASCADE"`
}

type Kecamatan struct {
	gorm.Model
	KecamatanName string
}

func (s *Kelurahan) ToModel() domain.Kelurahan {
	return domain.Kelurahan{
		ID:            int(s.ID),
		KelurahanName: s.KelurahanName,
		KecamatanID:   s.KecamatanID,
	}
}

func ParseToArr(arr []Kelurahan) []domain.Kelurahan {
	var res []domain.Kelurahan

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Kelurahan) Kelurahan {
	var res Kelurahan
	res.KelurahanName = data.KelurahanName
	res.KecamatanID = data.KecamatanID
	return res
}
