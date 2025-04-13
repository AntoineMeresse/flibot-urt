package models

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type RunCompare struct {
	playerName string
	checkpoint []int
}

type RunInfo struct {
	way        string
	checkpoint []int
	runCompare RunCompare
}

type RunsInfo struct {
	RunMutex   sync.RWMutex
	PlayerRuns map[string]RunInfo
}

func (runs *RunsInfo) RunStart(playerNumber string, wayName string) {
	log.Debugf("Starting run %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	runs.PlayerRuns[playerNumber] = RunInfo{way: wayName}
	log.Debugf("Run map %v", runs.PlayerRuns)
}
