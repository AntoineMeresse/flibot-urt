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
	level := log.DebugLevel
	if v := os.Getenv("logLevel"); v != "" {
		if parsed, err := log.ParseLevel(v); err == nil {
			level = parsed
		}
	}
	log.SetLevel(level)
	log.SetReportCaller(true)
	if os.Getenv("logFormat") == "dokploy" {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
			DisableColors:    true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + " | "
				return "", fileName
			},
			DisableLevelTruncation: true,
			PadLevelText:           true,
		})
	} else {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fileName := " " + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + " | "
				return "", fmt.Sprintf("%25.25s", fileName)
			},
			DisableLevelTruncation: true,
			PadLevelText:           true,
		})
	}
	log.SetOutput(os.Stdout)
}
