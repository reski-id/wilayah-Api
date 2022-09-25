package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	UserID      int
	User        User `gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE"`
	NIK         int
	PhoneNumber string
	Address     string
	ProvinsiID  int
	Provinsi    Provinsi `gorm:"foreignKey:ProvinsiID; references:ID; constraint:OnDelete:CASCADE"`
	KotaID      int
	Kota        Kota `gorm:"foreignKey:KotaID; references:ID; constraint:OnDelete:CASCADE"`
	KecamatanID int
	Kecamatan   Kecamatan `gorm:"foreignKey:KecamatanID; references:ID; constraint:OnDelete:CASCADE"`
	KelurahanID int
	Kelurahan   Kelurahan `gorm:"foreignKey:KelurahanID; references:ID; constraint:OnDelete:CASCADE"`
}

type Provinsi struct {
	gorm.Model
	ProvinsiName string
}

type Kota struct {
	gorm.Model
	KotaName string
}

type Kecamatan struct {
	gorm.Model
	KecamatanName string
}

type Kelurahan struct {
	gorm.Model
	KelurahanName string
}
type User struct {
	gorm.Model
	Username string
}

func (s *Member) ToModel() domain.Member {
	return domain.Member{
		ID:          int(s.ID),
		UserID:      s.UserID,
		NIK:         s.NIK,
		PhoneNumber: s.PhoneNumber,
		Address:     s.Address,
		ProvinsiID:  s.ProvinsiID,
		KotaID:      s.KotaID,
		KelurahanID: s.KelurahanID,
		KecamatanID: s.KecamatanID,
	}
}

func ParseToArr(arr []Member) []domain.Member {
	var res []domain.Member

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Member) Member {
	var res Member
	res.UserID = data.UserID
	res.NIK = data.NIK
	res.PhoneNumber = data.PhoneNumber
	res.ProvinsiID = data.ProvinsiID
	res.Address = data.Address
	res.KotaID = data.KotaID
	res.KelurahanID = data.KelurahanID
	res.KecamatanID = data.KecamatanID
	return res
}
