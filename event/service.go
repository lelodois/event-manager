package event

import (
	"errors"
	"eventManager/commons"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Migração da tabela Events
func Migrate() {
	err := commons.DbInstance.Connection.AutoMigrate(&Event{})
	if err != nil {
		panic(fmt.Sprintf("Could not migrate event error: %s", err))
	}
}

// Criação do evento na base de dados
func create(event Event) error {
	if event.Date.Before(time.Now()) {
		return errors.New("date must be in the future")
	}
	transaction := commons.DbInstance.Connection.Create(&event)
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

// Lista de todos os eventos existentes na base de dados
func list() ([]Event, error) {
	var events []Event
	transaction := commons.DbInstance.Connection.Find(&events)
	return events, transaction.Error
}

// Busca evento específico pelo id
func Get(eventId int) (*Event, error) {
	event := &Event{}
	transaction := commons.DbInstance.Connection.First(&event, eventId)
	return event, transaction.Error
}

// Diminui a disponibilidade de um evento (available = available -1)
func DecreaseAvailable(event *Event, transaction *gorm.DB) error {
	limitResult := transaction.Model(event).
		Where("(available - 1) >= 0 and id = ?", event.Id).
		Updates(map[string]interface{}{
			"available":  gorm.Expr("available - 1"),
			"updated_at": time.Now()})

	if limitResult.Error != nil {
		return limitResult.Error
	}

	if limitResult.RowsAffected != 1 {
		return errors.New(fmt.Sprintf("Event: [%v] unavailable ", event.Id))
	}
	return nil
}
