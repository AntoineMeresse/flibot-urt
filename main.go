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
	
	context := &models.Context{VoteChannel: voteChannel}
	context.Init()

	defer context.Rcon.CloseConnection()
	defer context.DB.Close();

	// Initialize tail
	go logparser.InitLogparser(myLogChannel, context.UrtConfig.LogFile)

	// Handle each line
	for i := 0; i < context.UrtConfig.WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i, context)
	}

	// Initialize Vote system
	go vote.InitVoteSystem(voteChannel, context)

	// Because we're only using go routines, if we don't have this block program isn't keep alived.
	keepRunning := true
	for keepRunning {
		time.Sleep(time.Second * 10)
		// Send server infos to bridge
	}
}
