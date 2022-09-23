package domain

import "github.com/labstack/echo/v4"

type Member struct {
	ID          int
	UserID      int
	NIK         int
	PhoneNumber string
	Address     string
	ProvinsiID  string
	KotaID      string
	KelID       string
	KecID       string
}

type MemberHandler interface {
	AddMember() echo.HandlerFunc
	GetMember() echo.HandlerFunc
	UpdateDataMember() echo.HandlerFunc
	DeleteDataMember() echo.HandlerFunc
}

type MemberUsecase interface {
	AddMember(newMember Member) (row int, err error)
	GetAllMember() ([]Member, error)
	UpdateMember(id int, UpdateMember Member) (row int, err error)
	DeleteMember(id int) (row int, err error)
}

type MemberData interface {
	InsertMember(newMember Member) (row int, err error)
	GetAll() ([]Member, error)
	Update(id int, updatedData Member) (row int, err error)
	Delete(id int) (row int, err error)
}
