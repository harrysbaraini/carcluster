package speed

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"carcluster/message"
	"github.com/stianeikeland/go-rpio"
	"net/http"
	"time"
)

var (
	pin = rpio.Pin(2)

	wheelsCirc = 205 // cm => π * diameter_cm

	start = time.Now()
	pulseAmount = 1
	elapse = 0
)

type speedData struct {
	Speed int `json:"speed"`
	RPM int `json:"rpm"`
}

func toMilli(t time.Duration) int {
	return int(t / time.Millisecond)
}

// func milesPerHour(sp float64) float64 {
// 	sp / wheelsRevPerMile * 60.0
// }

func Init() {
	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge)
}

func Listen() {
	if pin.EdgeDetected() {
		pinValue := pin.Read()

		if pinValue == 1 {
			currentTime := time.Now()
			elapse = toMilli(currentTime.Sub(start))

			fmt.Println("Elapse: ", elapse)

			if elapse != 0 {
				start = currentTime

				rpm := int(1 / (float64(elapse) / 1000.0) * 60.0)
				speed := float64(rpm) * (float64(wheelsCirc) / 100.0) * 3.6

				fmt.Println("RPM: ", rpm)
				fmt.Println("SPEED: ", speed)

				data := speedData{int(speed), rpm}

				http.Post("http://127.0.0.1:5000/speed", "application/json", data.Marshal())
			}
		}
	}
}

func Emit(socket *websocket.Conn, body []byte) error {
	var data speedData

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	msg := message.New("speed:update", data)
	msg.Send(socket)

	return nil
}
