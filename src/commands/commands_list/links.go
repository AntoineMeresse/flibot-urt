package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func Website(cmd *models.CommandsArgs) {
	cmd.RconText("^2[Site]^7 UJM: urtjumpmaps.com")
}

func Discord(cmd *models.CommandsArgs) {
	cmd.RconText("^2[Discord]^7 Urban Terror Jumping: ^5discord.gg/B2SMhvhbC8^7")
}
