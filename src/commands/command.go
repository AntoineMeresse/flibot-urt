package commands

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

type Command struct {
	Function     interface{}
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
		log.Debugf("Sending command to bridge: %s", info.message)
	} else {
		log.Debugf("Not sending command to bridge: %s", info.message)
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
			return commandInfo{command: command, isValid: true, isGlobal: isGlobal, name: name, params: params, message: text}
		}
	}
	return commandInfo{command: Command{sendToBridge: true}}
}

func checkPlayerRights(playerNumber string, command Command, context *models.Context) (canAccess bool, required int, got int) {
	log.Debugf("-------------------------------------------------------------")

	if command.Level == 0 {
		log.Debug("Command that can be used by everyone.")
		return true, 0, 0
	}

	player, err := context.Players.GetPlayer(playerNumber)
	var canUseCmd bool = false
	role := 0

	if err == nil {
		role = player.Role
		log.Debugf("checkPlayerRights | player (%v)", player)
		canUseCmd = role >= command.Level
	}

	return canUseCmd, command.Level, role
}

func overrideParamsForCommands(commandName string, role int, cmdArgs *models.CommandsArgs) {
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

func HandleCommand(actionParams []string, context *models.Context) {
	playerNumber := actionParams[0]
	commandInfos := extractCmdInfos(actionParams)
	if commandInfos.isValid {
		canAccess, level, role := checkPlayerRights(playerNumber, commandInfos.command, context)
		if canAccess {
			displayCommandInfos(actionParams[2], playerNumber, commandInfos.params, commandInfos.isGlobal)
			args := models.CommandsArgs{
				Context:  context,
				PlayerId: playerNumber,
				Params:   commandInfos.params,
				IsGlobal: commandInfos.isGlobal,
				Usage:    commandInfos.command.Usage,
			}
			overrideParamsForCommands(commandInfos.name, role, &args)
			commandInfos.command.Function.(func(*models.CommandsArgs))(&args)
		} else {
			log.Errorf("Player with id (%s) doesn't have enough rights to use command %s (required: %d | got: %d) ",
				playerNumber, actionParams[2], level, role,
			)
			context.RconText(false, playerNumber, msg.NOT_ENOUGH_RIGHTS, actionParams[2], level, role)
		}
	}
	err := commandInfos.sendCommandToBridge()
	if err != nil {
		log.Error(err)
	}
}

func displayCommandInfos(commandName string, playerNumber string, commandParams []string, isGlobal bool) {
	log.Debugf("Command: %s", commandName)
	log.Debugf("    |-> isGlobal: %v", isGlobal)
	log.Debugf("    |-> Playernumber: %s", playerNumber)
	log.Debugf("    |-> Params (%d): %v\n", len(commandParams), commandParams)
}
