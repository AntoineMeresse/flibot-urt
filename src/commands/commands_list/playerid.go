package commandslist

import (
	"fmt"
	"strconv"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)


func PlayerId(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		cmd.RconUsage()
		return
	}

	id, err := strconv.Atoi(cmd.Params[0])
	if err != nil {
		cmd.RconText("^1Invalid id: %s", cmd.Params[0])
		return
	}

	displayPlayerById(cmd, id)
}

func displayPlayerById(cmd *appcontext.CommandsArgs, id int) {
	r, found := cmd.Context.DB.GetPlayerById(id)
	if !found {
		cmd.RconText("^7No player found with id ^5%d", id)
		return
	}

	aliasesDisplay := r.Aliases
	if aliasesDisplay == "" {
		aliasesDisplay = "none"
	}

	cmd.RconText("^7PlayerId: ^5%d", r.Id)
	cmd.RconText("^7---> Name: ^3%s ^7(aliases: ^6%s^7)", r.Name, aliasesDisplay)
	cmd.RconText("^7---> Guid: ^5%s", r.Guid)
	geo := geoIPText(cmd, r.Ip)
	if geo == "" {
		cmd.RconText("^7---> Ip: ^5%s", r.Ip)
	} else {
		cmd.RconText("^7---> Ip: ^5%s ^7(^3%s)", r.Ip, geo)
	}

	sameIp, err := cmd.Context.DB.GetPlayersByIp(r.Ip)
	if err == nil {
		others := make([]string, 0)
		for _, p := range sameIp {
			if p.Id != r.Id {
				aliases := p.Aliases
				if aliases == "" {
					aliases = "none"
				}
				others = append(others, fmt.Sprintf("^5%d^7 / ^3%s ^7(^6%s^7)", p.Id, p.Name, aliases))
			}
		}
		if len(others) > 0 {
			cmd.RconText("^7---> Players with same Ip:")
			for _, line := range others {
				cmd.RconText("^7-----> %s", line)
			}
		}
	}
}
