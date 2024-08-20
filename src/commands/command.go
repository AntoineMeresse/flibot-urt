package commands

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/AntoineMeresse/flibot-urt/src/utils/msg"
)

type Command struct {
	Function interface{}
	Level int
	Usage string
}

func isCommand(text string) (bool) {
	return utils.IsVoteCommand(text) || (len(text) > 1 && (text[0] == '!' || text[0] == '@'))
}

func isCommandGlobal(text string) bool {
	return text[0] == '@'
}

func replaceShortcutByKnownCommand(cmd *string) {
	if val, ok := CommandsAlias[*cmd]; ok {
		*cmd = val;
	}
}

func extractCmdInfos(action_params []string) (iscommand bool, commandName string, command Command, isGlobal bool, params []string) {
	text := action_params[2]
	if isCommand(text) {
		var command string;
		if utils.IsVoteCommand(text) {
			command = text
		} else {
			command = strings.ToLower(text[1:])
		}
		replaceShortcutByKnownCommand(&command) 
		if cmd, ok := Commands[command]; ok {
			return true, command,  cmd, isCommandGlobal(text), utils.CleanEmptyElements(action_params[3:]) 
		}
	}
	return false, "", Command{}, false, nil
}

func checkPlayerRights(playerNumber string, command Command, context *models.Context) (canAccess bool, required int, got int) {
	log.Debugf("-------------------------------------------------------------")

	if (command.Level == 0) {
		log.Debug("Command that can be used by everyone.")
		return true, 0, 0
	}

	player, err := context.Players.GetPlayer(playerNumber)
	var canUseCmd bool = false;
	role := 0

	if (err == nil) {
		role = player.Role
		log.Debugf("checkPlayerRights | player (%v)", player)
		canUseCmd =  role >= command.Level
	}

	return canUseCmd, command.Level, role
}

func overrideParamsForCommands(commandName string, role int, cmdArgs *models.CommandsArgs) {
	if commandName == "help" {
		var cmdList []string;
		for key, value := range(Commands) {
			if (value.Level <= role) {	
				cmdList = append(cmdList, key)
			}
		}
		cmdArgs.Params =  utils.CleanEmptyElements(cmdList)
	}
}

func HandleCommand(action_params []string, context *models.Context) {
	playerNumber := action_params[0]
	isCommand, commandName, command, isGlobal, command_params := extractCmdInfos(action_params)
	if isCommand {
		canAccess, level, role := checkPlayerRights(playerNumber, command, context)
		if canAccess {
			displayCommandInfos(action_params[2], playerNumber, command_params, isGlobal)
			args := models.CommandsArgs{
				Context: context, 
				PlayerId: playerNumber, 
				Params: command_params, 
				IsGlobal: isGlobal,
				Usage: command.Usage,
			}
			overrideParamsForCommands(commandName, role, &args)
			command.Function.(func(*models.CommandsArgs))(&args)
		} else {
			log.Errorf("Player with id (%s) doesn't have enough rights to use command %s (required: %d | got: %d) ", 
				playerNumber, 
				action_params[2],
				level,
				role,
			)
			context.RconText(false, playerNumber, msg.NOT_ENOUGH_RIGHTS, action_params[2], level, role)
		}
	}
}

func displayCommandInfos(commandname string, playerNumber string, command_params []string, isGlobal bool) {
	log.Debugf("Command: %s", commandname)
	log.Debugf("    |-> isGlobal: %v", isGlobal)
	log.Debugf("    |-> Playernumber: %s", playerNumber)
	log.Debugf("    |-> Params (%d): %v\n", len(command_params), command_params)
}