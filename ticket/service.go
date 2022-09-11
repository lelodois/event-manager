package ticket

import (
	"errors"
	"eventManager/commons"
	events "eventManager/event"
	users "eventManager/user"
	"fmt"
	"log"
)

// Migração da tabela tickets
func Migrate() {
	err := commons.DbInstance.Connection.AutoMigrate(&Ticket{})
	if err != nil {
		panic(fmt.Sprintf("Could not migrate ticket error: %s", err))
	}
}

// Lista das tickets existentes na base de dados
func list() ([]Ticket, error) {
	var tickets []Ticket
	transaction := commons.DbInstance.Connection.Find(&tickets)
	return tickets, transaction.Error
}

// Verificação se existe algum ticket com o mesmo usuário e o evento que está sendo criado
func ExistsByUserAndEvent(newTicket Ticket) bool {
	ticket := &Ticket{}
	transaction := commons.DbInstance.Connection.Where(
		"user_id = ? and event_id = ? and deleted_at is null",
		newTicket.UserId, newTicket.EventId,
	).Find(&ticket)
	if transaction.Error != nil {
		return false
	}
	return ticket != nil && ticket.Id > 0
}

// Criação do ticket e validações de regra de negócio
func create(ticket Ticket) error {
	if ExistsByUserAndEvent(ticket) {
		return errors.New(fmt.Sprintf("Ticket already for user: [%v]", ticket.UserId))
	}

	event, eventErr := events.Get(ticket.EventId)
	if eventErr != nil {
		log.Println(fmt.Sprintf("Event not found for id: [%v]", ticket.EventId))
		return eventErr
	}
	user, userErr := users.Get(ticket.UserId)
	if userErr != nil {
		log.Println(fmt.Sprintf("User not found for id: [%v]", ticket.UserId))
		return userErr
	}

	if user.Balance < event.Price {
		return errors.New(fmt.Sprintf(
			"User: [%v] with balance: [%v] is not enough for event: [%v] with price [%v]",
			user.Id, user.Balance, event.Id, event.Price))
	}

	return reservation(ticket, event, user)
}
