package commandslist

import (
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func Time(cmd *appcontext.CommandsArgs) {
	now := time.Now()
	cmd.RconText("^7Server time: ^5%s", now.Format("2006-01-02 15:04:05"))
}
