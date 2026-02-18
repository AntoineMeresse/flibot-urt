package commandslist

import (
	"slices"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

const adminRole = 70

func Admins(cmd *appcontext.CommandsArgs) {
	cmd.Context.Players.Mutex.RLock()
	var admins []*models.Player
	for _, player := range cmd.Context.Players.PlayerMap {
		if player.Role >= adminRole {
			admins = append(admins, player)
		}
	}
	cmd.Context.Players.Mutex.RUnlock()

	slices.SortFunc(admins, func(a, b *models.Player) int {
		return b.Role - a.Role
	})

	cmd.RconText("^7====== ^6Admins ^7======")
	if len(admins) == 0 {
		cmd.RconText("No admins currently on the server.")
		return
	}
	for _, player := range admins {
		cmd.RconText("^5%s ^7[^2%d^7]", player.Name, player.Role)
	}
}
