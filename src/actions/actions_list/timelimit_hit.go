package actionslist

import (
	"log/slog"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

func TimelimitHit(actionParams []string, c *appcontext.AppContext) {
	slog.Debug("Timelimit hit", "params", actionParams)
	v := models.Vote{Params: []string{"extend"}}
	c.NewVote(v)
}
