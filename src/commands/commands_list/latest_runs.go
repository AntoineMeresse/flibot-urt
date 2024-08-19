package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/sirupsen/logrus"
)

func LatestRuns(cmd *models.CommandsArgs) {
	infos, err := cmd.Context.Api.GetLatestRuns()
	
	if err != nil {
		logrus.Error(err.Error())
		return;
	}

	if len(infos) > 0 {
		cmd.RconText("^7Latest runs :")
		for _, run := range(infos) {
			cmd.RconText("   ^5|--------> ^7%s ^5|^8 %s ^5|^7 (%s - Way %d)", run.RunTime, run.PlayerName, run.Mapname, run.Way)
		}
	}
}
