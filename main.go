package main

import (
	"time"

	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/vote"
)

const Logfile = "/home/antoine/UrbanTerror43/q3ut4/games.log"
const WorkerNumber = 3 // Param ?

func main() {
	configureLogger()

	// Channels
	myLogChannel := make(chan string)
	voteChannel := make(chan models.Vote)
	
	server := &models.Server{VoteChannel: voteChannel}
	server.Init()

	defer server.Rcon.CloseConnection()
	defer server.DB.Close();

	// Initialize tail
	go logparser.InitLogparser(myLogChannel, Logfile)

	// Handle each line
	for i := 0; i < WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i, server)
	}

	// Initialize Vote system
	go vote.InitVoteSystem(voteChannel, server)

	keepRunning := true
	for keepRunning {
		time.Sleep(time.Second * 10)
		// Send server infos to bridge
	}
}
