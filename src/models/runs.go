package models

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
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
	History    map[string][]int
}

type PlayerRunInfo struct {
	Server       string `json:"server"`
	ServerName   string `json:"server_name"`
	Fps          string `json:"fps"`
	Mapname      string `json:"mapname"`
	Playername   string `json:"playername"`
	Guid         string `json:"guid"`
	Way          string `json:"way"`
	Time         string `json:"time"`
	Demopath     string `json:"demopath"`
	Playernumber string `json:"playernumber"`
	Utj          string `json:"g_utj"`
	PlayerIp     string
}

func (p *PlayerRunInfo) GetDemoName() string {
	s := strings.Split(p.Demopath, "/")
	return s[len(s)-1]
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

	info, ok := runs.PlayerRuns[playerNumber]
	if !ok {
		log.Warnf("AddCheckpoint: no active run for player %s", playerNumber)
		return
	}
	info.appendCheckpoint(time)
}

func (runs *RunsInfo) RunCanceled(playerNumber string) {
	log.Debugf("RunCanceled %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	delete(runs.PlayerRuns, playerNumber)
}

func (runs *RunsInfo) RunStopped(playerNumber string, playerGuid string, time string) {
	log.Debugf("RunCanceled %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	info, ok := runs.PlayerRuns[playerNumber]
	if !ok {
		log.Warnf("RunStopped: no active run for player %s", playerNumber)
		return
	}
	var checkpoints []int
	checkpoints = append(checkpoints, info.checkpoint...)
	runId := fmt.Sprintf("%s-%s-%s", playerGuid, info.way, time)
	runs.History[runId] = checkpoints

	delete(runs.PlayerRuns, playerNumber)
}

func (runs *RunsInfo) RunGetCheckpoint(playerNumber string, playerGuid string, time string, way string) []int {
	log.Debugf("RunGetCheckpoint %s (guid: %s) -> %s", playerNumber, playerGuid, time)
	runId := fmt.Sprintf("%s-%s-%s", playerGuid, way, time)
	log.Debugf("Runid: %s", runId)
	checkpoints, exist := runs.History[runId]

	if !exist {
		return []int{}
	}

	delete(runs.History, runId)
	return checkpoints
}
