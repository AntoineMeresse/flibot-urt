package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func Bounties(cmd *appcontext.CommandsArgs) {
	bounties, err := cmd.Context.Api.GetBounties()

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	if len(bounties) == 0 {
		return
	}

	cmd.RconText("^7Bounties:")
	for _, b := range bounties {
		status := "^1[OPEN]"
		if b.Done {
			status = "^2[DONE]"
		}
		cmd.RconText("   ^5|-------->^7 %s ^5%s^7 (way %d) | beat: ^8%s^7 | until: ^8%s", status, b.Filename, b.WayNumber, b.TimeToBeat, b.Until)
	}
}
