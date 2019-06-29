package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Service running on port 5000...")

	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("<<< ERROR >>>", err)
	}
	for {
		// Read the message, then try to transform it to valid JSON
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println("<<< ERROR >>>", err)
			return
		}

		msgData := Message{}
		err = json.Unmarshal(msg, &msgData)

		if err != nil {
			fmt.Println("<<< ERROR >>> Could not unmarshal message", msg)
			return
		}

		fmt.Println("MESSAGE RECEIVED:", string(msg))

		// =========================================
		// Process messages
		// =========================================

		var responseMsg Message

		if msgData.Type == "speed:update" {
			responseMsg = OnSpeedUpdate(msgData)
		}

		if msgData.Type == "fuel:update" {
			// responseMsg = OnFuelUpdate(msgData)
		}

		// =========================================
		// Emit responses
		// =========================================

		responseStr, err := json.Marshal(responseMsg)
		if err != nil {
			fmt.Println("<<< ERROR >>>", err)
		}

		err = socket.WriteMessage(msgType, responseStr)
		if err != nil {
			fmt.Println("<<< ERROR >>>", err)
			return
		}
	}
}