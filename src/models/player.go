package models

import (
	"strings"
	"sync"
)

type Player struct {
	Guid string
	Name string
	Role int
}

type Players struct {
	Mutex sync.RWMutex
	List map[string]Player
}

func (players *Players) AddPlayer(playerNumber string, player Player) {
	// server.Players.Store(playerNumber, player)
	players.Mutex.Lock()
	players.List[playerNumber] = player
	players.Mutex.Unlock()
}

func (players *Players) RemovePlayer(playerNumber string) {
	// server.Players.Delete(playerNumber)
	players.Mutex.Lock()
	delete(players.List, playerNumber)
	players.Mutex.Unlock()
}

func (players *Players) GetPlayer(searchCriteria string)  (found bool, player *Player)  {
	// server.Players.Delete(playerNumber)
	
	players.Mutex.RLock()
	defer players.Mutex.RUnlock()

	for _, player := range(players.List) {
		if (strings.Contains(player.Name, searchCriteria)) {
			return true, &player
		}
	}
	
	return false, nil;
}

