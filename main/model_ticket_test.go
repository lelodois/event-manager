package main

import (
	"encoding/json"
	"eventManager/event"
	"eventManager/ticket"
	"eventManager/user"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SaveTicket(testing *testing.T, newTicket ticket.Request) {
	PostRequestParameter{testing: testing, path: "ticket", data: newTicket}.SendPostRequest()
}

func GetTicket(testing *testing.T, event event.Request, user user.Request) *ticket.Response {
	var ticketResponse *Ticket
	response := GetRequestParameter{testing: testing, path: "ticket"}.NewGetRequest()

	err := json.Unmarshal(response, &ticketResponse)
	if err != nil {
		assert.FailNow(testing, fmt.Sprintf("not parser ticket error: %v", err.Error()))
	}

	for _, item := range ticketResponse.Data {
		if item.UserResponse.Name == user.Name && item.EventResponse.Name == event.Name {
			return &item
		}
	}

	assert.FailNow(testing,
		fmt.Sprintf("not found ticket for user: %v and event: %v", user.Name, event.Name))
	return nil
}

type Ticket struct {
	Data []ticket.Response `json:"tickets" binding:"required"`
}
