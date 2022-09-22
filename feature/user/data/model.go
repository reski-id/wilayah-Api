package data

import (
	"wilayah/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
		Password:  u.Password,
		FullName:  u.FullName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.UserName
	res.Password = data.Password
	res.FullName = data.FullName
	return res
}
