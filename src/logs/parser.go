package logs

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/nxadm/tail"
)

func InitLogParser(myLogChannel chan string, c *appcontext.AppContext) {
	logfile := c.UrtConfig.LogFile
	logs, err := tail.TailFile(logfile, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})

	if err != nil {
		panic(err)
	}

	go func() {
		slog.Info("Sending playersDump command after logfile initialized")
		c.RconCommand("playersDump")
	}()

	for line := range logs.Lines {
		slog.Debug("New line in server log file", "line", line.Text)
		myLogChannel <- line.Text
		slog.Debug("New line after channel push")
	}
}
