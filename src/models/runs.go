package models

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

type CompareTarget struct {
	Name        string
	Runtime     int
	Checkpoints []int
}

type RunPlayerInfo struct {
	way            string
	checkpoint     []int
	compareTargets []CompareTarget
	compareIdxs    []int
}

type RunsInfo struct {
	RunMutex     sync.RWMutex
	PlayerRuns   map[string]*RunPlayerInfo
	History      map[string][]int
	CpEnabled    map[string]bool
	CpTargetIdxs map[string][]int
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

func (runs *RunsInfo) RunningPlayerNumbers() []string {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	numbers := make([]string, 0, len(runs.PlayerRuns))
	for number := range runs.PlayerRuns {
		numbers = append(numbers, number)
	}
	return numbers
}

func (runs *RunsInfo) AnyRunning() int {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	return len(runs.PlayerRuns)
}

func (runs *RunsInfo) IsRunning(playerNumber string) bool {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	_, ok := runs.PlayerRuns[playerNumber]
	return ok
}

func (runs *RunsInfo) RunStart(playerNumber string, wayName string, targets []CompareTarget) {
	log.Debugf("Starting run %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	idxs := runs.CpTargetIdxs[playerNumber]
	if len(idxs) == 0 {
		idxs = []int{0}
	}
	// clamp indices that exceed available targets
	valid := make([]int, 0, len(idxs))
	for _, idx := range idxs {
		if idx < len(targets) {
			valid = append(valid, idx)
		}
	}
	if len(valid) == 0 {
		valid = []int{0}
	}
	runs.PlayerRuns[playerNumber] = &RunPlayerInfo{way: wayName, checkpoint: []int{}, compareTargets: targets, compareIdxs: valid}
}

// GetTargetLimit returns how many DB rows to fetch at run start (max saved index + 1, min 1).
func (runs *RunsInfo) GetTargetLimit(playerNumber string) int {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	idxs := runs.CpTargetIdxs[playerNumber]
	max := 0
	for _, idx := range idxs {
		if idx > max {
			max = idx
		}
	}
	return max + 1
}

func (runs *RunsInfo) GetCurrentWay(playerNumber string) (string, bool) {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	info, ok := runs.PlayerRuns[playerNumber]
	if !ok {
		return "", false
	}
	return info.way, true
}

func (runs *RunsInfo) UpdateCompareTargets(playerNumber string, targets []CompareTarget, idxs []int) {
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()
	if info, ok := runs.PlayerRuns[playerNumber]; ok {
		info.compareTargets = targets
		info.compareIdxs = idxs
	}
	runs.CpTargetIdxs[playerNumber] = idxs
}

func (runs *RunsInfo) GetCompareTargets(playerNumber string) ([]CompareTarget, bool) {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	info, ok := runs.PlayerRuns[playerNumber]
	if !ok {
		return nil, false
	}
	return info.compareTargets, true
}

func (runs *RunsInfo) GetCompareIdxs(playerNumber string) []int {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	info, ok := runs.PlayerRuns[playerNumber]
	if ok {
		return info.compareIdxs
	}
	return runs.CpTargetIdxs[playerNumber]
}

func (runs *RunsInfo) ToggleCp(playerNumber string) bool {
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()
	runs.CpEnabled[playerNumber] = !runs.CpEnabled[playerNumber]
	return runs.CpEnabled[playerNumber]
}

func (runs *RunsInfo) EnableCp(playerNumber string) {
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()
	runs.CpEnabled[playerNumber] = true
}

func (runs *RunsInfo) IsCpEnabled(playerNumber string) bool {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()
	return runs.CpEnabled[playerNumber]
}

func (runs *RunsInfo) GetCpMsgs(playerNumber string) []string {
	runs.RunMutex.RLock()
	defer runs.RunMutex.RUnlock()

	info, ok := runs.PlayerRuns[playerNumber]
	if !ok || len(info.compareTargets) == 0 || len(info.checkpoint) == 0 {
		return nil
	}

	l := len(info.checkpoint)
	cpIdx := l - 1
	currentTime := info.checkpoint[cpIdx]

	parts := []string{fmt.Sprintf("^5CP %d:", cpIdx+1)}
	for _, idx := range info.compareIdxs {
		if idx >= len(info.compareTargets) {
			continue
		}
		target := info.compareTargets[idx]
		if cpIdx >= len(target.Checkpoints) {
			continue
		}
		bestTime := target.Checkpoints[cpIdx]
		diff := bestTime - currentTime

		var sign string
		absDiff := diff
		if diff == 0 {
			sign = "^7"
		} else if diff < 0 {
			sign = "^1-"
			absDiff = -diff
		} else {
			sign = "^2+"
		}
		parts = append(parts, fmt.Sprintf("^7[^3#%d^7: ^8%s^7 %s%s^7]", idx+1, target.Name, sign, FormatMs(absDiff)))
	}

	if len(parts) == 1 {
		return nil
	}
	return parts
}

func FormatMs(ms int) string {
	minutes := ms / 60000
	seconds := (ms % 60000) / 1000
	millis := ms % 1000
	if minutes > 0 {
		return fmt.Sprintf("%d:%02d.%03d", minutes, seconds, millis)
	}
	return fmt.Sprintf("%d.%03d", seconds, millis)
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

func (runs *RunsInfo) ClearRuns() {
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()
	runs.PlayerRuns = map[string]*RunPlayerInfo{}
	runs.History = map[string][]int{}
	runs.CpEnabled = map[string]bool{}
	runs.CpTargetIdxs = map[string][]int{}
}

func (runs *RunsInfo) RunCanceled(playerNumber string) {
	log.Debugf("RunCanceled %s", playerNumber)
	runs.RunMutex.Lock()
	defer runs.RunMutex.Unlock()

	delete(runs.PlayerRuns, playerNumber)
	delete(runs.CpEnabled, playerNumber)
	delete(runs.CpTargetIdxs, playerNumber)
}

func (runs *RunsInfo) RunStopped(playerNumber string, playerGuid string, time string) {
	log.Debugf("RunStopped %s", playerNumber)
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
