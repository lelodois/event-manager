package event

import (
	"gorm.io/gorm"
	"time"
)

// Modelo de domínio Event - persistente
type Event struct {
	gorm.Model
	Id        int `gorm:"primary_key;auto_increment;not_null"`
	Capacity  int
	Available int
	Name      string
	Date      time.Time
	Price     float32
}

// Requisição para criação de um evento
type Request struct {
	Name     string    `json:"name" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
	Price    float32   `json:"ticket_price" binding:"required"`
	Capacity int       `json:"capacity" binding:"required"`
}

// Resposta dos eventos existentes
type Response struct {
	Id        int       `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Price     float32   `json:"ticket_price" binding:"required"`
	Capacity  int       `json:"capacity" binding:"required"`
	Available int       `json:"available" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

// Criação da resposta com base no modelo de domínio Event
func createResponse(event Event) Response {
	return Response{
		Id:        event.Id,
		Name:      event.Name,
		Date:      event.Date,
		Available: event.Available,
		Price:     event.Price,
		Capacity:  event.Capacity,
		CreatedAt: event.CreatedAt,
	}
}

// Criação do domínio Event com base na requisição
func (request Request) toModel() Event {
	return Event{
		Name:      request.Name,
		Date:      request.Date,
		Capacity:  request.Capacity,
		Price:     request.Price,
		Available: request.Capacity,
	}
}
