package models

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

type Player struct {
	Id      string
	Number  string
	Guid    string
	Name    string
	Ip      string
	Aliases []string
	Role    int
}

type DumpPlayer struct {
	PlayerNumber int    `json:"playernumber"`
	Name         string `json:"name"`
	GUID         string `json:"guid"`
}

type Players struct {
	Mutex     sync.RWMutex
	PlayerMap map[string]*Player
}

func (players *Players) AddPlayer(playerNumber string, player *Player) {
	logrus.Debugf("AddPlayer: %s -> %v", playerNumber, player)
	players.Mutex.Lock()
	player.Number = playerNumber
	players.PlayerMap[playerNumber] = player
	players.Mutex.Unlock()
}

func (p *Player) hasInfoChange(infos map[string]string) bool {
	return p.Name != infos["name"] || p.Guid != infos["cl_guid"]
}

func (players *Players) UpdatePlayer(currentPlayer *Player, infos map[string]string) bool {
	logrus.Debug("UpdatePlayer called")
	if currentPlayer.hasInfoChange(infos) {
		players.Mutex.Lock()
		if name, ok := infos["name"]; ok {
			currentPlayer.Name = utils.DecolorString(name)
		}
		if guid, ok := infos["cl_guid"]; ok {
			currentPlayer.Guid = guid
		}
		if ipAndPort, ok := infos["ip"]; ok {
			ipAddress := strings.Split(ipAndPort, ":")[0]
			currentPlayer.Ip = ipAddress
		}
		players.Mutex.Unlock()
		return true
	}
	return false
}

func (players *Players) UpdatePlayerRights(playerNumber string, level int) {
	currentPlayer := players.PlayerMap[playerNumber]
	if currentPlayer == nil {
		logrus.Warnf("Player %s not found. Can't update rights.", playerNumber)
		return
	}
	players.Mutex.Lock()
	currentPlayer.Role = level
	players.Mutex.Unlock()
}

func (players *Players) RemovePlayer(playerNumber string) {
	players.Mutex.Lock()
	delete(players.PlayerMap, playerNumber)
	players.Mutex.Unlock()
}

func (players *Players) GetPlayer(searchCriteria string) (*Player, error) {
	players.Mutex.RLock()
	defer players.Mutex.RUnlock()

	var matchingPlayers []Player
	var alreadyAdded bool

	for playerNumber, player := range players.PlayerMap {
		alreadyAdded = false
		if utils.IsDigitOnly(searchCriteria) {
			if playerNumber == searchCriteria {
				matchingPlayers = append(matchingPlayers, *player)
				alreadyAdded = true
			}
		}

		if !alreadyAdded && strings.Contains(strings.ToLower(player.Name), strings.ToLower(searchCriteria)) {
			matchingPlayers = append(matchingPlayers, *player)
		}
	}

	if len(matchingPlayers) == 1 {
		return &matchingPlayers[0], nil
	} else if len(matchingPlayers) == 0 {
		return nil, fmt.Errorf("no player found using (%s)", searchCriteria)
	} else {
		var playerList []string

		for _, p := range matchingPlayers {
			playerList = append(playerList, fmt.Sprintf("%s [%s]", p.Name, p.Number))
		}

		playersDisplay := strings.Join(playerList, ", ")

		return nil, fmt.Errorf("multiple match (%s) using (%s)", playersDisplay, searchCriteria)
	}
}
