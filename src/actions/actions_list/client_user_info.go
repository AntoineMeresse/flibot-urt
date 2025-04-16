package actionslist

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientUserinfo(actionParams []string, context *models.Context) {
	log.Debugf("Client User Info: %v", actionParams)
	if len(actionParams) > 1 {
		playerNumber := actionParams[0]
		infoString := strings.Join(actionParams[1:], "")
		infos := splitInfos(infoString)
		context.Players.UpdatePlayer(playerNumber, infos)
	}
}

func splitInfos(infos string) map[string]string {
	res := make(map[string]string)
	infoSplit := utils.CleanEmptyElements(strings.Split(infos, "\\"))
	for i := 0; i < len(infoSplit)-1; i += 2 {
		res[infoSplit[i]] = infoSplit[i+1]
	}
	return res
}
