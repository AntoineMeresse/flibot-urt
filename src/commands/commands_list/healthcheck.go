package commandslist

import (
	"encoding/json"
	"os"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/healthcheck"
)

func HealthCheck(cmd *appcontext.CommandsArgs) {
	data, err := os.ReadFile(healthcheck.HealthFilePath)
	if err != nil {
		cmd.RconText("^1Health file not found: %s", err.Error())
		return
	}

	var status healthcheck.Status
	if err := json.Unmarshal(data, &status); err != nil {
		cmd.RconText("^1Failed to parse health file: %s", err.Error())
		return
	}

	aliveDisplay := "^2YES"
	if !status.Alive {
		aliveDisplay = "^1NO"
	}

	cmd.RconText("^7Server alive: %s", aliveDisplay)
	cmd.RconText("^7Last check: ^5%s", status.LastCheck.Format("2006-01-02 15:04:05"))
	cmd.RconText("^7Uptime since: ^5%s", status.UptimeStart.Format("2006-01-02 15:04:05"))
}
