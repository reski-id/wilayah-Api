package domain

import "time"

type User struct {
	//struct
	ID        int
	UserName  string
	Password  string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCase interface {
	//usecase
	AddUser(newUser User) (User, error)
	LoginUser(userLogin User) (row int, data User, err error)
	GetProfile(id int) (User, error)
	DeleteUser(id int) (row int, err error)
	UpdateUser(id int, updateProfile User) (User, error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(userLogin User) (row int, data User, err error)
	GetSpecific(userID int) (User, error)
	Delete(userID int) (row int, err error)
	Update(userID int, updatedData User) User
}
