package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func configureLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + " | "
			//return frame.Function, fileName
			return "", fmt.Sprintf("%25.25s", fileName)
		},
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	log.SetOutput(os.Stdout)
}
