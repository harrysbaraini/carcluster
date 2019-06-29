package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Error error `json:"error, omitempty"`
}

type MessageInterface interface {
	FromResponse(v interface{}) error
	ToResponse(v interface{}) (Message, error)
}


func (m Message) FromResponse(v interface{}) error {
	err := json.Unmarshal([]byte(m.Data), &v)

	if err != nil {
		fmt.Println("<<< ERROR >>> Could not unmarshal data", err)
	}

	return err
}

func (m Message) ToResponse(v interface{}) Message {
	js, err := json.Marshal(v)

	if err != nil {
		m.Data = ""
		m.Error = err
		return m
	}


	m.Data = string(js)
	m.Error = nil
	return m
}