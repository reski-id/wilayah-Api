package delivery

import "wilayah/domain"

type MemberResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	NIK         int    `json:"nik"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	ProvinsiID  string `json:"provinsi_id"`
	KotaID      string `json:"kota_id"`
	KelID       string `json:"kel_id"`
	KecID       string `json:"kec_id"`
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
	res.KecID = data.KecID
	return res
}

func FromModelToList(data []domain.Member) []MemberResponse {
	result := []MemberResponse{}
	for val := range data {
		result = append(result, FromModel(data[val]))
	}
	return result
}
