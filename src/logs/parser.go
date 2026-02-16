package logs

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
)

func InitLogParser(myLogChannel chan string, c *appcontext.AppContext) {
	logfile := c.UrtConfig.LogFile
	logs, err := tail.TailFile(logfile, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})

	if err != nil {
		panic(err)
	}

	go func() {
		logrus.Infof("Sending playersDump command after logfile initialized.")
		c.RconCommand("playersDump")
	}()

	for line := range logs.Lines {
		logrus.Tracef("Channel: %v | New line in server log file: %s", myLogChannel, line.Text)
		myLogChannel <- line.Text
		logrus.Trace("New line after channel push")
	}
}
