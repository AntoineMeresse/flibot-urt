package models

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

type RunCompare struct {
	playerName string
	checkpoint []int
}

type RunPlayerInfo struct {
	way        string
	checkpoint []int
	runCompare RunCompare
}

type RunsInfo struct {
	RunMutex   sync.RWMutex
	PlayerRuns map[string]*RunPlayerInfo
}

func (runs *RunsInfo) RunStart(playerNumber string, wayName string) {
	log.Debugf("Starting run %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	runs.PlayerRuns[playerNumber] = &RunPlayerInfo{way: wayName, checkpoint: []int{}}
}

func (i *RunPlayerInfo) appendCheckpoint(time string) {
	if v, err := strconv.Atoi(time); err == nil {
		i.checkpoint = append(i.checkpoint, v)
	} else {
		log.Errorf("Error converting time to int %v", err)
	}
}

func (runs *RunsInfo) AddCheckpoint(playerNumber string, time string) {
	log.Debugf("AddCheckpoint %s -> %s", playerNumber, time)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	runs.PlayerRuns[playerNumber].appendCheckpoint(time)
}

func (runs *RunsInfo) RunCanceled(playerNumber string) {
	log.Debugf("RunCanceled %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	delete(runs.PlayerRuns, playerNumber)
}
