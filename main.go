package main

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/vote"
)

func main() {
	configureLogger()

	// Channels
	myLogChannel := make(chan string)
	voteChannel := make(chan models.Vote)

	c := &appcontext.AppContext{VoteChannel: voteChannel}
	c.Init()

	defer c.Rcon.CloseConnection()
	defer c.DB.Close()

	// Initialize tail
	go logparser.InitLogParser(myLogChannel, c)

	// Handle each line
	for i := 0; i < c.UrtConfig.WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i, c)
	}

	// Initialize Vote system
	go vote.InitVoteSystem(voteChannel, c)

	// Because we're only using go routines, if we don't have this block program isn't keep alived.
	for {
		time.Sleep(time.Second * 10)
		// Send server infos to bridge
	}
}
