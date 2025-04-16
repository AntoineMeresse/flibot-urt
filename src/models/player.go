package models

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

type Player struct {
	Id   string
	Guid string
	Name string
	Role int
}

type Players struct {
	Mutex     sync.RWMutex
	PlayerMap map[string]*Player
}

func (players *Players) AddPlayer(playerNumber string, player *Player) {
	logrus.Debugf("AddPlayer: %s -> %v", playerNumber, player)
	players.Mutex.Lock()
	players.PlayerMap[playerNumber] = player
	players.Mutex.Unlock()
}

func (p *Player) hasInfoChange(infos map[string]string) bool {
	return p.Name != infos["name"] || p.Guid != infos["cl_guid"]
}

func (players *Players) UpdatePlayer(playerNumber string, infos map[string]string) {
	currentPlayer := players.PlayerMap[playerNumber]

	if currentPlayer == nil {
		logrus.Warnf("Player %s not found. Creating it", playerNumber)
		currentPlayer = &Player{}
		players.AddPlayer(playerNumber, currentPlayer)
	}

	if currentPlayer.hasInfoChange(infos) {
		players.Mutex.Lock()
		if currentPlayer.Guid == "" {
			logrus.Debugf("Player %v has no guid. Init player with: %v", playerNumber, infos)
			currentPlayer.Role = 100 // TODO: fetch role
		}
		if name, ok := infos["name"]; ok {
			currentPlayer.Name = utils.DecolorString(name)
		}
		if guid, ok := infos["cl_guid"]; ok {
			currentPlayer.Guid = guid
		}
		players.Mutex.Unlock()
	}
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
			playerList = append(playerList, fmt.Sprintf("%s [%s]", p.Name, p.Id))
		}

		playersDisplay := strings.Join(playerList, ", ")

		return nil, fmt.Errorf("multiple match (%s) using (%s)", playersDisplay, searchCriteria)
	}
}
