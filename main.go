package main

import (
	"time"

	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/vote"
)

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
	go logparser.InitLogparser(myLogChannel, server.UrtConfig.LogFile)

	// Handle each line
	for i := 0; i < server.UrtConfig.WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i, server)
	}

	// Initialize Vote system
	go vote.InitVoteSystem(voteChannel, server)

	// Because we're only using go routines, if we don't have this block program isn't keep alived.
	keepRunning := true
	for keepRunning {
		time.Sleep(time.Second * 10)
		// Send server infos to bridge
	}
}
