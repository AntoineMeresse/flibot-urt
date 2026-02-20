package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/quake3_rcon"
)

var blockedRconCommands = []string{"quit", "exit", "shutdown", "killserver"}

func Rcon(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) < 1 {
		cmd.Context.RconText(false, cmd.PlayerId, cmd.Usage)
		return
	}
	command := strings.ToLower(cmd.Params[0])
	for _, blocked := range blockedRconCommands {
		if command == blocked {
			cmd.Context.RconText(false, cmd.PlayerId, "^1Command %s is not allowed.", command)
			return
		}
	}
	rcon := NewRconClient(cmd)
	defer rcon.CloseConnection()

	result := rcon.RconCommand(strings.Join(cmd.Params, " "))
	_, lines := quake3_rcon.SplitReadInfos(result)
	for _, line := range lines {
		cmd.Context.RconText(false, cmd.PlayerId, "^7%s", line)
	}
}

func NewRconClient(cmd *appcontext.CommandsArgs) *quake3_rcon.Rcon {
	rcon := &quake3_rcon.Rcon{
		ServerIp:   cmd.Context.Rcon.ServerIp,
		ServerPort: cmd.Context.Rcon.ServerPort,
		Password:   cmd.Context.Rcon.Password,
	}
	rcon.Connect()
	return rcon
}
