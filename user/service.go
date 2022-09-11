package user

import (
	"errors"
	"eventManager/commons"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Migração da tabela users
func Migrate() {
	err := commons.DbInstance.Connection.AutoMigrate(&User{})
	if err != nil {
		panic(fmt.Sprintf("Could not migrate user error: %s", err))
	}
}

// Criação do user na base de dados
func create(user User) error {
	transaction := commons.DbInstance.Connection.Create(&user)
	if transaction.Error != nil {
		return transaction.Error
	}
	return nil
}

// Listagem dos usuários existentes na base de dados
func list() ([]User, error) {
	var users []User
	transaction := commons.DbInstance.Connection.Find(&users)
	return users, transaction.Error
}

// Busca um usuário específico pelo id
func Get(id int) (*User, error) {
	user := &User{}
	transaction := commons.DbInstance.Connection.First(&user, id)
	return user, transaction.Error
}

// Diminui o saldo do usuário com base em um valor recebido
func DecreaseBalance(amount float32, user *User, transaction *gorm.DB) error {
	result := transaction.Model(user).
		Where("(balance - ?) >= 0 and id = ?", amount, user.Id).
		Updates(map[string]interface{}{
			"balance":    gorm.Expr("balance - ?", amount),
			"updated_at": time.Now()})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return errors.New(fmt.Sprintf("User: [%v] not has enough balance ", user.Id))
	}
	return nil
}
