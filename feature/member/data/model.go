package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	UserID      int
	NIK         int
	PhoneNumber string
	Address     string
	ProvinsiID  string
	KotaID      string
	KelID       string
	KecID       string
}

type Provinsi struct {
	gorm.Model
	ProvinsiName string
	ProvinsiCode string
}

type Kota struct {
	gorm.Model
	KotaName     string
	ProvinsiCode string
}

type Kecamatan struct {
	gorm.Model
	KecamatanName string
	KotaID        string
}

type Kelurahan struct {
	gorm.Model
	KelurahanName string
	KecamatanID   string
}
type User struct {
	gorm.Model
	Username string
	Password string
	FullName string
	Email    string
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
		KelID:       s.KelID,
		KecID:       s.KecID,
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
	res.KelID = data.KelID
	res.KecID = data.KecID
	return res
}
