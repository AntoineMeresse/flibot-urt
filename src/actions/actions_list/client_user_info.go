package actionslist

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientUserinfo(action_params []string, context *models.Context) {
	log.Debugf("Client User Info: %v", action_params)
	if (len(action_params) > 1) {
		playerNumber := action_params[0]
		infoString := strings.Join(action_params[1:], "")
		infos := splitInfos(infoString)

		log.Debugf("Infos: \n%v\n", infos)
		
		context.Players.AddPlayer(playerNumber, generatePlayer(playerNumber, infos))
	}
	
}

func splitInfos(infos string) map[string]string {
	res := make(map[string]string)

	infoSplit := utils.CleanEmptyElements(strings.Split(infos, "\\"))

	for i:=0; i < len(infoSplit)-1; i+=2 {
		res[infoSplit[i]] = infoSplit[i+1]
	}

	return res;
}

func generatePlayer(playerNumber string, infos map[string]string) models.Player {
	player := models.Player{};
	
	if name, ok := infos["name"]; ok {
		player.Name = utils.DecolorString(name);
	}

	if guid, ok := infos["cl_guid"]; ok {
		player.Guid = guid;
	}

	player.Id = playerNumber
	player.Role = 100; // Todo: change with db rights

	return player                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     
}