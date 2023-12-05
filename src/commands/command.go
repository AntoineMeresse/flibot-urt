package commands

import (
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/models"
)

type Command struct {
	Function interface{}
	Level int
	Usage string
}

func isCommand(text string) (bool) {
	return len(text) > 1 && (text[0] == '!' || text[1] == '@')
}

func isCommandGlobal(text string) bool {
	return text[0] == '@'
}

func replaceShortcutByKnownCommand(cmd *string) {
	if val, ok := CommandsShortcut[*cmd]; ok {
		*cmd = val;
	}
}

func extractCmdInfos(action_params []string) (iscommand bool, command Command, isGlobal bool, params []string) {
	text := action_params[2]
	if isCommand(text) {
		command := strings.ToLower(text[1:])
		replaceShortcutByKnownCommand(&command) 
		if cmd, ok := Commands[command]; ok {
			return true, cmd, isCommandGlobal(text), action_params[3:] 
		}
	}
	return false, Command{}, false, nil
}

func checkPlayerRights(playerNumber string,command Command) bool {
	playerRights := 100 // replace
	return playerRights >= command.Level;
}

func HandleCommand(action_params []string, server models.Server) {
	playerNumber := action_params[0]
	isCommand, command, isGlobal, command_params := extractCmdInfos(action_params)
	if isCommand && checkPlayerRights(playerNumber, command) {
		command.Function.(func(models.Server, string, []string, bool))(server, playerNumber, command_params, isGlobal)
	}
}