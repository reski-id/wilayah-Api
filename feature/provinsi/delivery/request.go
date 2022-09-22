package delivery

import "wilayah/domain"

type InserFormat struct {
	ProvinsiName string `json:"provinsi_name" form:"provinsi_name"`
	ProvinsiCode string `json:"provinsi_code" form:"provinsi_code"`
}

func (i InserFormat) ToModel() domain.Provinsi {
	return domain.Provinsi{
		ProvinsiName: i.ProvinsiName,
		ProvinsiCode: i.ProvinsiCode,
	}
}
