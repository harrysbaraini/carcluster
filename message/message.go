package message

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Error error `json:"error"`
}

type MsgInterface interface {
	Marshal() ([]byte, error)
	Send(socket *websocket.Conn) error
}

func (m *Message) Marshal() ([]byte, error) {
	msgStr, err := json.Marshal(m)

	return msgStr, err
}

func (m *Message) Send(socket *websocket.Conn) error {
	msgStr, err := m.Marshal()

	if err != nil {
		fmt.Println("<<< ERROR >>>", err)
		return err
	}

	err = socket.WriteMessage(websocket.TextMessage, msgStr)
	if err != nil {
		fmt.Println("<<< ERROR >>>", err)
		return err
	}

	fmt.Println("Message sent: ", msgStr)

	return nil
}

func New(eventType string, data interface{}) Message {
	m := Message{eventType, "", nil}

	dataStr, err := json.Marshal(data)

	if err != nil {
		m.Error = err
		return m
	}

	m.Data = string(dataStr)
	m.Error = nil
	return m
}