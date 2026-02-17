package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/AntoineMeresse/flibot-urt/src/models"

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

func ReadSchema(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}

	return string(data)
}

type PositionData struct {
	X     float64
	Y     float64
	Z     float64
	Angle float64
}

type DataPersister interface {
	Close()

	GetPlayerByGuid(guid string) (models.Player, bool)
	SaveNewPlayer(name string, guid string, ipAddress string) (int, error)
	UpdatePlayer() error

	PenAdd(guid string, size float64) error
	PenPlayerGetDailySize(guid string) (float64, error)
	PenPenOfTheDay() (string, []PenData, error)
	PenPenHallOfFame() ([]PenData, error)
	PenPenHallOfShame() ([]PenData, error)

	HandleRun(info models.PlayerRunInfo, checkpoints []int) error

	PositionSave(mapname, location string, x, y, z, angle float64) error
	PositionGet(mapname, location string) (PositionData, error)
	PositionList(mapname string) ([]string, error)
	PositionDelete(mapname, location string) bool
}
