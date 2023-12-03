package logs

import (
	"github.com/nxadm/tail"
)

func InitLogparser(myLogChannel chan string, logfile string) {
	logs, err := tail.TailFile(logfile, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})

	if err != nil {
		panic(err)
	}

	for line := range logs.Lines {
		myLogChannel <- line.Text
	}
}