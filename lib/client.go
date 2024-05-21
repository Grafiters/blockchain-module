package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ClientLib struct {
	Server string
}

func NewClient(server string) ClientLib {
	return ClientLib{
		Server: server,
	}
}

func (cl ClientLib) Post(path string, body bytes.Buffer) (interface{}, error) {
	var target interface{}
	serverUrl := cl.Server + "/" + path
	requestClient, err := http.Post(serverUrl, "application/json", &body)
	if err != nil {
		return nil, err
	}

	defer requestClient.Body.Close()

	err = json.NewDecoder(requestClient.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return target, nil
}
