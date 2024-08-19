package commandslist

import (
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

func LatestMaps(cmd *models.CommandsArgs) {
	infos, err := cmd.Context.Api.GetLatestMaps()
	
	if err != nil {
		logrus.Error(err.Error())
		return;
	}

	if len(infos) > 0 {
		cmd.RconText("^7Latest maps:")
		for _, mapInfo := range(infos) {
			dateRelease := strings.ReplaceAll(mapInfo.Date, " 00:00:00 GMT", "")
			cmd.RconText("   ^5|-------->^7 ^7%s ^5|^8 %s ^5|^7 (%s)", mapInfo.Filename, strings.Join(mapInfo.Mappers, ", "), dateRelease)
		}
	}
}
