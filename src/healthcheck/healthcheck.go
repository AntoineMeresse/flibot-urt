package healthcheck

import (
	"encoding/json"
	"os"
	"time"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/quake3_rcon"
	log "github.com/sirupsen/logrus"
)

const (
	HealthFilePath = "/tmp/flibot.health"
	Interval       = 30 * time.Second
)

type Status struct {
	Alive       bool      `json:"alive"`
	LastCheck   time.Time `json:"last_check"`
	UptimeStart time.Time `json:"uptime_start"`
}

func Run(c *appcontext.AppContext) {
	healthCheckRcon := quake3_rcon.Rcon{
		ServerIp:   c.UrtConfig.ServerConfig.Ip,
		ServerPort: c.UrtConfig.ServerConfig.Port,
		Password:   c.UrtConfig.ServerConfig.Password,
	}
	healthCheckRcon.Connect()
	defer healthCheckRcon.CloseConnection()

	start := time.Now()

	for {
		response := healthCheckRcon.RconCommand("serverinfo")
		alive := response != ""

		status := Status{
			Alive:       alive,
			LastCheck:   time.Now(),
			UptimeStart: start,
		}

		if err := write(status); err != nil {
			log.Errorf("[healthcheck] Failed to write health file: %v", err)
		} else {
			log.Tracef("[healthcheck] alive=%v", alive)
		}

		time.Sleep(Interval)
	}
}

func write(status Status) error {
	data, err := json.Marshal(status)
	if err != nil {
		return err
	}
	return os.WriteFile(HealthFilePath, data, 0644)
}
