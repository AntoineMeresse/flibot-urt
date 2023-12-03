package main

import (
	"fmt"
	"time"

	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
)

const Logfile = "/home/antoine/UrbanTerror43/q3ut4/games.log"
const WorkerNumber = 3 // Param ?

func main() {
	fmt.Printf("Flibot starting\n")

	// Variables
	myLogChannel := make(chan string)
	keepRunning := true

	// Initialize tail
	go logparser.InitLogparser(myLogChannel, Logfile)

	// Handle each line
    for i := 0; i < WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i)
	}

	for keepRunning {
		time.Sleep(time.Second * 10)
		// Send server infos to bridge
	}
	
}