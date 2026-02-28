package commandslist

import (
	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

func Locate(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	target, err := cmd.Context.Players.GetPlayer(cmd.Params[0])
	if err != nil {
		cmd.RconText("%s", err.Error())
		return
	}

	if target.Ip == "" {
		cmd.RconText("^1No IP available for ^5%s^7.", target.Name)
		return
	}

	result, err := cmd.Context.Api.GetGeoIP(target.Ip)
	if err != nil {
		log.Errorf("[Locate] Error: %v", err)
		cmd.RconText("^1Failed to locate player.")
		return
	}

	if result.Status != "success" {
		cmd.RconText("^1Could not locate ^5%s^7.", target.Name)
		return
	}

	caller, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err == nil && caller.Role >= 90 {
		cmd.RconText("^5%s^7: ^3%s^7, ^3%s^7 (^6%s^7)", target.Name, result.Country, result.RegionName, result.Timezone)
	} else {
		cmd.RconText("^5%s^7: ^3%s", target.Name, result.Country)
	}
}
