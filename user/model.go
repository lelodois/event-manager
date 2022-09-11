package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id      int     `gorm:"primary_key;auto_increment;not_null"`
	Name    string  `gorm:"not_null"`
	Balance float32 `gorm:"not_null"`
}

type Request struct {
	Name    string  `json:"name" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}

type Response struct {
	Id        int       `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Balance   float32   `json:"balance" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

func createResponse(user User) Response {
	return Response{
		Id:        user.Id,
		Name:      user.Name,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}

func (request Request) toModel() User {
	return User{
		Name:    request.Name,
		Balance: request.Balance,
	}
}
