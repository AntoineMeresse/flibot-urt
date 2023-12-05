package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
	"github.com/joho/godotenv"
)

const Logfile = "/home/antoine/UrbanTerror43/q3ut4/games.log"
const WorkerNumber = 3 // Param ?

func main() {
	fmt.Printf("Flibot starting\n")

	// Variables
	myLogChannel := make(chan string)
	keepRunning := true
	
	rcon, err := getRcon()

	if err == nil {
		rcon.Connect()
		defer rcon.CloseConnection()

		server := models.Server{Rcon : rcon}

		// Initialize tail
		go logparser.InitLogparser(myLogChannel, Logfile)

		// Handle each line
		for i := 0; i < WorkerNumber; i++ {
			go logparser.HandleLogsWorker(myLogChannel, i, server)
		}

		for keepRunning {
			time.Sleep(time.Second * 10)
			// Send server infos to bridge
		}
	}
}

func getRcon() (quake3_rcon.Rcon, error) {
	err := godotenv.Load()
	if err != nil {
		return quake3_rcon.Rcon{}, errors.New("Could not load .env file")
	}

	serverIp := os.Getenv("serverip") 
	serverPort := os.Getenv("serverport") 
	port, err := strconv.Atoi(serverPort)
	password := os.Getenv("password") 
	

	return quake3_rcon.Rcon{ServerIp: serverIp, ServerPort: port, Password: password}, nil
}