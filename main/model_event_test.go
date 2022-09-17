package main

import (
	"encoding/json"
	"eventManager/event"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SaveEvent(testing *testing.T, newEvent event.Request) int {
	PostRequestParameter{testing: testing, path: "event", data: newEvent}.SendPostRequest()
	return getEventId(testing, newEvent)
}

func getEventId(testing *testing.T, newEvent event.Request) int {
	var eventResponse *Events
	response := GetRequestParameter{testing: testing, path: "event"}.NewGetRequest()

	err := json.Unmarshal(response, &eventResponse)
	if err != nil {
		assert.FailNow(testing, fmt.Sprintf("not parser event error: %v", err.Error()))
	}

	for _, item := range eventResponse.Data {
		if item.Name == newEvent.Name {
			return item.Id
		}
	}

	assert.FailNow(testing, fmt.Sprintf("not found event by name: %v", newEvent.Name))
	return 0
}

type Events struct {
	Data []event.Response `json:"events" binding:"required"`
}
