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
	level := log.InfoLevel
	if v := os.Getenv("logLevel"); v != "" {
		if parsed, err := log.ParseLevel(v); err == nil {
			level = parsed
		}
	}
	log.SetLevel(level)
	log.SetReportCaller(true)
	if os.Getenv("logFormat") == "json" {
		log.SetFormatter(&log.JSONFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return "", path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			},
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
