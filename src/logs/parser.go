package logs

import (
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
)

func InitLogparser(myLogChannel chan string, logfile string) {
	logs, err := tail.TailFile(logfile, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})

	if err != nil {
		panic(err)
	}

	for line := range logs.Lines {
		logrus.Tracef("Channel: %v | New line in server log file: %s", myLogChannel, line.Text)
		myLogChannel <- line.Text
		logrus.Trace("New line after channel push")
	}
}