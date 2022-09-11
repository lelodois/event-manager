package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Disponibiza um bind de um modelo e retorna erro caso não consiga
func BindRequest(context *gin.Context, instance interface{}) bool {
	if err := context.ShouldBindJSON(&instance); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

// Disponibiza um bind de sucesso ou erro caso
func WriteResponseBy(context *gin.Context, err error) {
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusCreated, http.StatusCreated)
	}
}

// Escreve a resposta em json com o array recebido
func WriteResponseByArray(context *gin.Context, arrayName string, any interface{}) {
	context.JSON(http.StatusOK, gin.H{arrayName: any})
}
