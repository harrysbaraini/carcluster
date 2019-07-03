package server

import (
	"carcluster/speed"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	// Set WebSocket settings
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	socket *websocket.Conn
)

func main() {
	http.HandleFunc("/", socketHandler)
	http.HandleFunc("/speed", speedHandler)

	fmt.Println("Service running on port 5000...")

	http.ListenAndServe(":5000", nil)
}

func listenSensors() {
	speed.Init()

	for {
		speed.Listen()
		time.Sleep(50)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	s, err := upgrader.Upgrade(w, r, nil)


	if err != nil {
		fmt.Println("<<< ERROR >>>", err)
	}

	socket = s
}

func speedHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = speed.Emit(socket, b)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write([]byte("OK"))
}