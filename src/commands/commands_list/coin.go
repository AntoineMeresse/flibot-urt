package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func Coin(cmd *appcontext.CommandsArgs) {
	cmd.RconText(utils.RandomValueFromSlice([]string{"Heads!", "Tails!"}))
}
