package main

import (
	"encoding/json"
	"eventManager/user"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SaveUser(testing *testing.T, newUser user.Request) int {
	PostRequestParameter{testing: testing, path: "user", data: newUser}.SendPostRequest()
	return getUserId(testing, newUser)
}

func getUserId(testing *testing.T, newUser user.Request) int {
	var userResponse *Users
	response := GetRequestParameter{testing: testing, path: "user"}.NewGetRequest()

	err := json.Unmarshal(response, &userResponse)
	if err != nil {
		assert.FailNow(testing, fmt.Sprintf("not parser user error: %v", err.Error()))
	}

	for _, item := range userResponse.Data {
		if item.Name == newUser.Name && item.Email == newUser.Email {
			return item.Id
		}
	}

	assert.FailNow(testing, fmt.Sprintf("not found user by name: %v", newUser.Name))
	return 0
}

type Users struct {
	Data []user.Response `json:"users" binding:"required"`
}
