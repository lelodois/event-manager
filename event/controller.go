package event

import (
	"eventManager/commons"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Retorna um json com os eventos ( "events" : [])
func List(context *gin.Context) {
	events, err := list()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var list []Response
		for _, user := range events {
			list = append(list, createResponse(user))
		}
		commons.WriteResponseByArray(context, "events", list)
	}
}

// Cria um evento com base no Request
func Create(context *gin.Context) {
	var eventRequest Request
	if commons.BindRequest(context, &eventRequest) {
		err := create(eventRequest.toModel())
		commons.WriteResponseBy(context, err)
	}
}
