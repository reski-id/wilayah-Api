package delivery

import "wilayah/domain"

type InserFormat struct {
	UserID      int    `json:"user_id" form:"user_id"`
	NIK         int    `json:"nik" form:"nik"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	ProvinsiID  int    `json:"provinsi_id" form:"provinsi_id"`
	KotaID      int    `json:"kota_id" form:"kota_id"`
	KelurahanID int    `json:"kel_id" form:"kel_id"`
	KecamatanID int    `json:"kec_id" form:"kec_id"`
}

func (i InserFormat) ToModel() domain.Member {
	return domain.Member{
		UserID:      i.UserID,
		NIK:         i.NIK,
		PhoneNumber: i.PhoneNumber,
		Address:     i.Address,
		ProvinsiID:  i.ProvinsiID,
		KotaID:      i.KotaID,
		KelurahanID: i.KelurahanID,
		KecamatanID: i.KecamatanID,
	}
}
