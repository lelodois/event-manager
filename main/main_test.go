package main

import (
	"eventManager/event"
	"eventManager/ticket"
	"eventManager/user"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	server          = "http://localhost:8080"
	contentTypeJson = "application/json"
)

func TestIntegrationTest(testing *testing.T) {
	go main()
	time.Sleep(time.Second * 2)

	// create user and recover by name
	newUser := user.Request{
		Name:    faker.FirstNameFemale() + " " + faker.LastName(),
		Email:   faker.Email(),
		Balance: 100.00,
	}
	userId := SaveUser(testing, newUser)
	fmt.Println(fmt.Sprintf("new user: [%v] with id: [%v]", newUser.Name, userId))

	// create event and recover by name
	newEvent := event.Request{
		Name: faker.Name(), Date: now.EndOfDay(), Price: 10.00, Capacity: 2,
	}
	eventId := SaveEvent(testing, newEvent)
	fmt.Println(fmt.Sprintf("new event: [%v] with id: [%v]", newEvent.Name, eventId))

	// create ticket and recover by event name and name user
	SaveTicket(testing, ticket.Request{UserId: userId, EventId: eventId})
	newTicket := GetTicket(testing, newEvent, newUser)
	fmt.Println(fmt.Sprintf("new ticket: [%v]", newTicket.Id))

	assert.NotEqual(testing, newTicket, nil)
	assert.GreaterOrEqual(testing, newTicket.Id, 1)
	// event with decreased available
	assert.Equal(testing, newTicket.EventResponse.Available, 1)
	// user with decreased balance
	assert.Equal(testing, newTicket.UserResponse.Balance, float32(90.00))
}
