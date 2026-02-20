package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/sirupsen/logrus"
)

func Birthday(cmd *appcontext.CommandsArgs) {
	birthdays, err := cmd.Context.Api.GetBirthdays()

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	if len(birthdays) == 0 {
		return
	}

	cmd.RconText("^7Birthdays:")
	for _, b := range birthdays {
		cmd.RconText("   ^5|-------->^7 ^7%s ^5|^8 %d years ago", b.Filename, b.Years)
	}
}
