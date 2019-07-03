package sensors

import (
	"carcluster/speed"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func main() {
	initGpio()
	listenSensors()
}

func initGpio() {
	// Open and map memory to access gpio, check for errors
	fmt.Println("Opening GPIO...")

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("GPIO Opened")

	// Unmap gpio memory when done
	defer rpio.Close()
}

func listenSensors() {
	speed.Init()

	for {
		speed.Listen()
		time.Sleep(50)
	}
}