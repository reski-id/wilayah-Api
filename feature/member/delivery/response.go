package delivery

import "wilayah/domain"

type MemberResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	NIK         int    `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ProvinsiID  int    `json:"provinsi_id" form:"provinsi_id"`
	KotaID      int    `json:"kota_id" form:"kota_id"`
	KecamatanID int    `json:"kec_id" form:"kec_id"`
	KelurahanID int    `json:"kel_id" form:"kel_id"`
}

func FromModel(data domain.Member) MemberResponse {
	var res MemberResponse
	res.ID = data.ID
	res.UserID = data.UserID
	res.NIK = data.NIK
	res.PhoneNumber = data.PhoneNumber
	res.Address = data.Address
	res.ProvinsiID = data.ProvinsiID
	res.KotaID = data.KotaID
	res.KecamatanID = data.KecamatanID
	res.KelurahanID = data.KelurahanID
	return res
}

func FromModelToList(data []domain.Member) []MemberResponse {
	result := []MemberResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
