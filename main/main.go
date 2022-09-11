package main

import (
	"eventManager/commons"
	"eventManager/event"
	"eventManager/ticket"
	"eventManager/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// Start da app, cria a conexão com a base, realiza as migrações e configura as rotas rest
func main() {
	log.Println("Starting app")
	commons.CreateDatabaseConnection()
	event.Migrate()
	ticket.Migrate()
	user.Migrate()

	router := gin.Default()
	configRouter(router)
	routerError := router.Run()
	if routerError != nil {
		panic(fmt.Sprintf("Cannot config gin: %s ", routerError.Error()))
	}
	log.Println("App ready for new requests")

}

// Configura as rotas rest
func configRouter(router *gin.Engine) {
	router.POST("/user", user.Create)
	router.GET("/user", user.List)

	router.POST("/event", event.Create)
	router.GET("/event", event.List)

	router.POST("/ticket", ticket.Create)
	router.GET("/ticket", ticket.List)
}
