package commands

import (
	"fmt"
	"log/slog"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

type Command struct {
	Function     func(*appcontext.CommandsArgs)
	Level        int
	Usage        string
	sendToBridge bool
}

type commandInfo struct {
	command  Command
	isValid  bool
	isGlobal bool
	name     string
	message  string
	params   []string
}

func (info *commandInfo) sendCommandToBridge() error {
	// TODO: implement logic
	if info.command.sendToBridge {
		slog.Debug("Sending command to bridge", "message", info.message)
	} else {
		slog.Debug("Not sending command to bridge", "message", info.message)
	}
	return fmt.Errorf("sendCommandToBridge not implemented yet")
}

func isCommand(text string) bool {
	return utils.IsVoteCommand(text) || (len(text) > 1 && (text[0] == '!' || text[0] == '@'))
}

func isCommandGlobal(text string) bool {
	return text[0] == '@'
}

func replaceShortcutByKnownCommand(cmd *string) {
	if val, ok := Alias[*cmd]; ok {
		*cmd = val
	}
}

func extractCmdInfos(actionParams []string) (command commandInfo) {
	text := actionParams[2]
	message := strings.Join(actionParams[2:], " ")
	if isCommand(text) {
		var name string
		if utils.IsVoteCommand(text) {
			name = text
		} else {
			name = strings.ToLower(text[1:])
		}
		replaceShortcutByKnownCommand(&name)
		if command, ok := Commands[name]; ok {
			isGlobal := isCommandGlobal(text)
			params := utils.CleanEmptyElements(actionParams[3:])
			return commandInfo{command: command, isValid: true, isGlobal: isGlobal, name: name, params: params, message: message}
		}
	}
	return commandInfo{command: Command{sendToBridge: true}, message: message}
}

func checkPlayerRights(playerNumber string, command Command, c *appcontext.AppContext) (canAccess bool, required int, got int) {
	slog.Debug("-------------------------------------------------------------")

	player, err := c.Players.GetPlayer(playerNumber)
	var canUseCmd = false
	role := 0

	if err == nil {
		role = player.Role
		slog.Debug("checkPlayerRights | player", "player", player)
		canUseCmd = role >= command.Level
	}

	if command.Level == 0 {
		slog.Debug("Command that can be used by everyone.")
		canUseCmd = true
	}

	return canUseCmd, command.Level, role
}

func overrideParamsForCommands(commandName string, role int, cmdArgs *appcontext.CommandsArgs) {
	if commandName == "help" {
		var cmdList []string
		for key, value := range Commands {
			if value.Level <= role {
				cmdList = append(cmdList, key)
			}
		}
		cmdArgs.Params = utils.CleanEmptyElements(cmdList)
	}
}

func HandleCommand(actionParams []string, c *appcontext.AppContext) {
	playerNumber := actionParams[0]
	commandInfos := extractCmdInfos(actionParams)
	if commandInfos.isValid {
		canAccess, level, role := checkPlayerRights(playerNumber, commandInfos.command, c)
		if canAccess {
			displayCommandInfos(actionParams[2], playerNumber, commandInfos.params, commandInfos.isGlobal)
			args := appcontext.CommandsArgs{
				Context:  c,
				PlayerId: playerNumber,
				Params:   commandInfos.params,
				IsGlobal: commandInfos.isGlobal,
				Usage:    commandInfos.command.Usage,
			}
			overrideParamsForCommands(commandInfos.name, role, &args)
			commandInfos.command.Function(&args)
		} else {
			slog.Error(fmt.Sprintf("Player with id (%s) doesn't have enough rights to use command %s (required: %d | got: %d)",
				playerNumber, actionParams[2], level, role))
			c.RconText(false, playerNumber, msg.NOT_ENOUGH_RIGHTS, actionParams[2], level, role)
		}
	}
	err := commandInfos.sendCommandToBridge()
	if err != nil {
		//Todo: uncomment slog.Error(err.Error())
	}
}

func displayCommandInfos(commandName string, playerNumber string, commandParams []string, isGlobal bool) {
	slog.Debug("Command", "name", commandName)
	slog.Debug("    |-> isGlobal", "isGlobal", isGlobal)
	slog.Debug("    |-> Playernumber", "playerNumber", playerNumber)
	slog.Debug(fmt.Sprintf("    |-> Params (%d)", len(commandParams)), "params", commandParams)
}
