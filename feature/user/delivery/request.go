package delivery

import "wilayah/domain"

type InsertFormat struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		UserName: i.UserName,
		Password: i.Password,
		FullName: i.FullName,
		Email:    i.Email,
	}
}

type LoginFormat struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		UserName: lf.UserName,
		Password: lf.Password,
	}
}
