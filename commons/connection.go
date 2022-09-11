// Pacote comum entre os domínios
package commons

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DatabaseConfiguration struct {
	Connection *gorm.DB
}

// Conecta na base e atribui a deixa disponível a conexão para a pp
func CreateDatabaseConnection() {
	dsn := "user_test:pwd_test@tcp(127.0.0.1:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Cannot open error: %s dbConn in the database", err.Error()))
	}
	log.Println("Database connected")
	DbInstance = DatabaseConfiguration{Connection: dbConn}
}

// Compartilha a conexão criada no start para a app
var DbInstance DatabaseConfiguration
