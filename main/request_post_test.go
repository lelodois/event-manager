package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type PostRequestParameter struct {
	testing *testing.T
	path    string
	data    interface{}
}

func (parameter PostRequestParameter) SendPostRequest() PostRequestParameter {
	response, requestError := parameter.doNewPostRequest()

	if requestError != nil {
		assert.FailNow(parameter.testing, "Could not send the user request", requestError.Error())
	}

	if response == nil {
		assert.FailNow(parameter.testing, "Could not send the user request")
	}

	if response.StatusCode != 200 && response.StatusCode != 201 {
		assert.FailNow(parameter.testing, "Could not save the user")
	}
	return parameter
}

func (parameter PostRequestParameter) doNewPostRequest() (*http.Response, error) {
	marshal, errParser := json.Marshal(parameter.data)
	if errParser != nil {
		assert.FailNow(parameter.testing,
			fmt.Sprintf("Could not parser the struct[%v] to request: [%v]",
				parameter.data, parameter.path))
	}

	return http.Post(parameter.getFullPath(), contentTypeJson, bytes.NewBuffer(marshal))
}
func (parameter PostRequestParameter) getFullPath() string {
	return fmt.Sprintf("%v/%v", server, parameter.path)
}
