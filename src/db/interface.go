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

type DataPersister interface {
	Close()

	GetPlayerByGuid(guid string) (models.Player, bool)
	SaveNewPlayer(name string, guid string, ipAddress string) (int, error)
	UpdatePlayer() error
	SetPlayerRole(guid string, role int) error

	PenAdd(guid string, size float64) error
	PenPlayerGetDailySize(guid string) (float64, error)
	PenPenOfTheDay() (string, []PenData, error)
	PenPenHallOfFame() ([]PenData, error)
	PenPenHallOfShame() ([]PenData, error)

	HandleRun(info models.PlayerRunInfo, checkpoints []int) error

	SetMapOptions(mapname, options string) error
	GetMapOptions(mapname string) (string, bool)
	DeleteMapOptions(mapname string) (bool, error)

	SaveGoto(mapname, jumpname string, posX, posY, posZ, angleV, angleH float64) error
	GetGoto(mapname, jumpname string) (GotoData, bool)
	GetGotoNames(mapname string) ([]string, error)
	DeleteGoto(mapname, jumpname string) (bool, error)
}

type GotoData struct {
	PosX   float64
	PosY   float64
	PosZ   float64
	AngleV float64
	AngleH float64
}
