package db

import (
	"database/sql"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

type PenData struct {
	Name sql.NullString
	Size float64
	Date time.Time
}

func (p PenData) GetName() string {
	if p.Name.String == "" {
		return "Unknown"
	}
	//x := models.Player{}
	//logrus.Debugf("%v", x)
	return p.Name.String
}

func (p PenData) GetDate() string {
	return utils.FormatTimeToDate(p.Date)
}

type DataPersister interface {
	Close()

	SaveNewPlayer(name string, guid string, ipAddress string) error
	InitRight(guid string) error
	UpdatePlayer() error

	PenAdd(guid string, size float64) error
	PenPenOfTheDay() (string, []PenData, error)
	PenPenHallOfFame() ([]PenData, error)
	PenPenHallOfShame() ([]PenData, error)

	HandleRun(info models.PlayerRunInfo, checkpoints []int) error
}
