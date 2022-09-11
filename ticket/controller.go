package ticket

import (
	"eventManager/commons"
	events "eventManager/event"
	users "eventManager/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Lista Tickets existentes
func List(context *gin.Context) {
	tickets, err := list()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var list []Response
		for _, ticket := range tickets {
			user, _ := users.Get(ticket.UserId)
			event, _ := events.Get(ticket.EventId)
			list = append(list, createResponse(ticket, *user, *event))
		}
		commons.WriteResponseByArray(context, "tickets", list)
	}
}

// Criação de um ticket com base no user + event
func Create(context *gin.Context) {
	var ticketRequest Request
	if commons.BindRequest(context, &ticketRequest) {
		err := create(ticketRequest.toModel())
		commons.WriteResponseBy(context, err)
	}
}
