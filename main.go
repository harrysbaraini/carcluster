package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
	"os"
)

var (
	// Set WebSocket settings
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	pin = rpio.Pin(22)
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Service running on port 5000...")

	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge) // enable falling edge event detection

	// Open a WebSocket connection
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("<<< ERROR >>>", err)
	}

	// =========================================
	// Process messages
	// =========================================

	for {
		var responseMsg Message

		// DETECT SPEED SENSOR EVENT
		if pin.EdgeDetected() {
			fmt.Println("EVENT: SPEED SENSOR")

			responseMsg = OnSpeedUpdate()
		}

		// =========================================
		// Emit responses
		// =========================================

		responseStr, err := json.Marshal(responseMsg)
		if err != nil {
			fmt.Println("<<< ERROR >>>", err)
		}

		err = socket.WriteMessage(websocket.TextMessage, responseStr)
		if err != nil {
			fmt.Println("<<< ERROR >>>", err)
			return
		}
	}
}