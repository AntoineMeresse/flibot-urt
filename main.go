package main

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/healthcheck"
	logparser "github.com/AntoineMeresse/flibot-urt/src/logs"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/portgotos"
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

	if c.UrtConfig.PortGotoPath != "" {
		portgotos.PortGotos(c.UrtConfig.PortGotoPath, c.DB)
	}
	if c.UrtConfig.PortMapOptionsPath != "" {
		portgotos.PortMapOptions(c.UrtConfig.PortMapOptionsPath, c.DB)
	}

	// Initialize tail
	go logparser.InitLogParser(myLogChannel, c)

	// Handle each line
	for i := 0; i < c.UrtConfig.WorkerNumber; i++ {
		go logparser.HandleLogsWorker(myLogChannel, i, c)
	}

	// Initialize Vote system
	go vote.InitVoteSystem(voteChannel, c)

	// Keep-alive loop: probes the server every 30s and writes health status to file.
	healthcheck.Run(c)
}
