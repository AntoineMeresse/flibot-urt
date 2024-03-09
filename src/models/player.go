package models

import (
	"fmt"
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
	Mutex sync.RWMutex
	List map[string]Player
}

func (players *Players) AddPlayer(playerNumber string, player Player) {
	players.Mutex.Lock()
	players.List[playerNumber] = player
	players.Mutex.Unlock()
}

func (players *Players) RemovePlayer(playerNumber string) {
	players.Mutex.Lock()
	delete(players.List, playerNumber)
	players.Mutex.Unlock()
}

func (players *Players) GetPlayer(searchCriteria string)  (*Player, error)  {
	players.Mutex.RLock()
	defer players.Mutex.RUnlock()

	matchingPlayers := []Player{}
	var alreadyAdded bool;

	for playerNumber, player := range(players.List) {
		alreadyAdded = false;
		if utils.IsDigitOnly(searchCriteria) {
			if (playerNumber == searchCriteria) {
				matchingPlayers = append(matchingPlayers, player);
				alreadyAdded = true;
			}
		}

		if (!alreadyAdded && strings.Contains(strings.ToLower(player.Name), strings.ToLower(searchCriteria))) {
			matchingPlayers = append(matchingPlayers, player);
		}
	}
	

	if len(matchingPlayers) == 1 {
		return &matchingPlayers[0], nil
	} else if len(matchingPlayers) == 0 {
		return nil, fmt.Errorf("no player found using (%s)", searchCriteria)
	} else {
		playerList := []string{}

		for _, p := range matchingPlayers {
			playerList = append(playerList, fmt.Sprintf("%s [%s]", p.Name, p.Id))
		}

		playersDisplay := strings.Join(playerList, ", ")

		return nil, fmt.Errorf("multiple match (%s) using (%s)", playersDisplay , searchCriteria)
	}
}

