package commandslist

import (
	"github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ServerStatus(cmd *context.CommandsArgs) {
	cmd.RconText("Status !")
	infos, err := cmd.Context.Api.GetServerStatus()

	if err != nil {
		cmd.RconText(err.Error())
		return
	}

	for _, serverlist := range infos {
		for servername, serverInfos := range serverlist {
			cmd.RconText("^7  |^8----------------------------------------------------------------- ")
			cmd.RconText("^7  |--->^5%s^7 - %s ( ^2%d^7 / 24 )", servername, serverInfos.Mapname, serverInfos.NbPlayers)
			if len(serverInfos.Ingame) > 0 {
				ingame := []string{"^7  |---------> ^2In game^7: "}
				ingame = append(ingame, serverInfos.Ingame...)
				for _, ingamePlayerLine := range utils.ToShorterChunkArraySep(ingame, ", ", true) {
					cmd.RconText(ingamePlayerLine)
				}
			}
			if len(serverInfos.Spec) > 0 {
				inspec := []string{"^7  |---------> ^1In spec^7: "}
				inspec = append(inspec, serverInfos.Spec...)
				for _, specPlayerLine := range utils.ToShorterChunkArraySep(inspec, ", ", true) {
					cmd.RconText(specPlayerLine)
				}
			}
		}
	}
}
