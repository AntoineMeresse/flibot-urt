package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/db"
	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

const Logfile = "/home/antoine/UrbanTerror43/q3ut4/games.log"
const WorkerNumber = 3 // Param ?


func initLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {                                                     
			fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)  + " | "      
			//return frame.Function, fileName                                        
			return "", fmt.Sprintf("%20.20s", fileName)                                                      
		},    
		DisableLevelTruncation: true,
		PadLevelText: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	initLogger()

	// Variables
	myLogChannel := make(chan string)
	keepRunning := true
	
	rcon, rconErr := getRcon()
	db, dbErr := db.InitDb("test.db")

	if rconErr == nil && dbErr == nil {
		rcon.Connect()
		
		defer rcon.CloseConnection()
		defer db.Close()

		server := models.Server{Rcon : rcon, Db: db}
		server.Init()

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
		return quake3_rcon.Rcon{}, errors.New("could not load .env file")
	}

	serverIp := os.Getenv("serverip") 
	serverPort := os.Getenv("serverport") 
	password := os.Getenv("password") 
	
	return quake3_rcon.Rcon{ServerIp: serverIp, ServerPort: serverPort, Password: password}, nil
}