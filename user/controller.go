package user

import (
	"eventManager/commons"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Listagem de todos os usuários existentes
func List(context *gin.Context) {
	users, err := list()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		var list []Response
		for _, user := range users {
			list = append(list, createResponse(user))
		}
		commons.WriteResponseByArray(context, "users", list)
	}
}

// Criação do usuário
func Create(context *gin.Context) {
	var userRequest Request
	if commons.BindRequest(context, &userRequest) {
		err := create(userRequest.toModel())
		commons.WriteResponseBy(context, err)
	}
}
