package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Message struct {
	//struct
	ID              int
	Message         string
	MessageParentID int
	IDUser          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MessageUseCase interface {
	//usecase
	AddMessage(IDUser int, useMessage Message) (Message, error)
	DelMessage(IDMessage int) (bool, error)
}

type MessageHandler interface {
	//handler
	InsertMessage() echo.HandlerFunc
	DeleteMessage() echo.HandlerFunc
}

type MessageData interface {
	//data
	Insert(insertMessage Message) Message
	Delete(IDMessage int) bool
}
