package commandslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func GlobalPen(cmd *appcontext.CommandsArgs) {
	cmd.Context.Players.Mutex.RLock()
	numbers := make([]string, 0, len(cmd.Context.Players.PlayerMap))
	for number := range cmd.Context.Players.PlayerMap {
		numbers = append(numbers, number)
	}
	cmd.Context.Players.Mutex.RUnlock()

	adminName := cmd.PlayerId
	if player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId); err == nil {
		adminName = player.Name
	}
	cmd.RconGlobalText("^7With great power comes great responsibility: ^5Pen comp^7 started by ^3%s", adminName)

	go func() {
		for _, number := range numbers {
			cmd.RconCommand("spoof %s say !pen [auto pen]", number)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
