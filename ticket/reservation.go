package ticket

import (
	"eventManager/commons"
	events "eventManager/event"
	users "eventManager/user"
	"fmt"
	"gorm.io/gorm"
	"log"
)

// Criação de uma transação que realiza a 1- diminuição do event.available, 2- diminuição do user.balance e 3-criação de ticket
func reservation(ticket Ticket, event *events.Event, user *users.User) error {

	return commons.DbInstance.Connection.Transaction(func(transaction *gorm.DB) error {
		availableErr := events.DecreaseAvailable(event, transaction)
		if availableErr != nil {
			return availableErr
		}

		balanceErr := users.DecreaseBalance(event.Price, user, transaction)
		if balanceErr != nil {
			return balanceErr
		}

		ticketErr := transaction.Create(&ticket)
		if ticketErr.Error != nil {
			return transaction.Error
		}
		log.Println(fmt.Sprintf(
			"decrease available eventId: [%v], "+
				"decrease amount: [%v] of user balance: [%v], "+
				"because create ticket: [%v]",
			event.Id, event.Price, user.Id, ticket.Id))
		return nil
	})
}
