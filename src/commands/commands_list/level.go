package commandslist

import (
	"slices"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Level(cmd *appcontext.CommandsArgs) {
	cmd.Context.Players.Mutex.RLock()
	players := make([]*models.Player, 0, len(cmd.Context.Players.PlayerMap))
	for _, player := range cmd.Context.Players.PlayerMap {
		players = append(players, player)
	}
	cmd.Context.Players.Mutex.RUnlock()

	slices.SortFunc(players, func(a, b *models.Player) int {
		return b.Role - a.Role
	})

	cmd.RconText("^7====== ^6Levels ^7======")
	for _, player := range players {
		cmd.RconText("^5%s ^7[^2%d^7]", player.Name, player.Role)
	}
}
