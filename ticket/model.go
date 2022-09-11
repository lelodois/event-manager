package ticket

import (
	events "eventManager/event"
	users "eventManager/user"
	"gorm.io/gorm"
	"time"
)

// Modelo de domain Ticket
type Ticket struct {
	gorm.Model
	Id      int `gorm:"primary_key;auto_increment;not_null"`
	EventId int
	UserId  int
}

// Modelo de requisição para criação de um Ticket
type Request struct {
	UserId  int `json:"user_id" binding:"required"`
	EventId int `json:"event_id" binding:"required"`
}

// Modelo de resposta dos tickets criados
type Response struct {
	Id            int           `json:"id" binding:"required"`
	UserResponse  UserResponse  `json:"user" binding:"required"`
	EventResponse EventResponse `json:"event" binding:"required"`
	CreatedAt     time.Time     `json:"created_at" binding:"required"`
}

// Modelo de resposta dos usuários de um ticket
type UserResponse struct {
	Name    string  `json:"name" binding:"required"`
	Balance float32 `json:"balance" binding:"required"`
}

// Modelo de resposta dos eventos de um ticket
type EventResponse struct {
	Name      string  `json:"name" binding:"required"`
	Price     float32 `json:"ticket_price" binding:"required"`
	Capacity  int     `json:"capacity" binding:"required"`
	Available int     `json:"available" binding:"required"`
}

// Criação do modelo de resposta de um ticket
func createResponse(model Ticket, user users.User, event events.Event) Response {
	return Response{
		Id: model.Id,
		UserResponse: UserResponse{
			Name:    user.Name,
			Balance: user.Balance,
		},
		EventResponse: EventResponse{
			Name:      event.Name,
			Price:     event.Price,
			Capacity:  event.Capacity,
			Available: event.Available,
		},
		CreatedAt: model.CreatedAt,
	}
}

// Criação do modelo do domain ticket
func (request Request) toModel() Ticket {
	return Ticket{
		EventId: request.EventId,
		UserId:  request.UserId,
	}
}
