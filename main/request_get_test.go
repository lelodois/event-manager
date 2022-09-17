package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type GetRequestParameter struct {
	testing *testing.T
	path    string
}

func (parameter GetRequestParameter) getFullPath() string {
	return fmt.Sprintf("%v/%v", server, parameter.path)
}

func (parameter GetRequestParameter) NewGetRequest() []byte {
	responseByte, err := parameter.doNewGetRequest()

	if err != nil {
		assert.FailNow(
			parameter.testing,
			fmt.Sprintf("response not found for get request: [%v]", parameter.path))
	}

	return responseByte
}

func (parameter GetRequestParameter) doNewGetRequest() ([]byte, error) {
	response, err := http.Get(parameter.getFullPath())

	if err != nil {
		assert.FailNow(
			parameter.testing,
			fmt.Sprintf("not send get request: [%v]", parameter.path))
	}

	return ioutil.ReadAll(response.Body)
}
