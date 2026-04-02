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
	UpdatePlayerOnJoin(guid, name, ip string, aliases []string) error
	SetPlayerRole(guid string, role int) error

	PenAdd(guid string, size float64) error
	PenSetAllAttempts(attempts int) error
	PenInitIfNotExists(guid string) error
	PenPlayerGetDailySize(guid string) (float64, error)
	PenGetAttempts(guid string) (int, error)
	PenIncrementAttempts(guid string) error
	PenDecrementAttempts(guid string) error
	PenPenOfTheDay() (string, []PenData, error)
	PenPenHallOfFame() ([]PenData, error)
	PenPenHallOfShame() ([]PenData, error)
	GetDailyPbPenCoinCount(guid string) (int, error)
	IncrementDailyPbPenCoinCount(guid string) error

	HandleRun(info models.PlayerRunInfo, checkpoints []int) error
	GetBestCheckpoints(mapname, way string) ([]int, string, error)
	GetTopCheckpoints(mapname, way string, limit int) ([]TopCheckpoint, error)

	SetMapOptions(mapname, options string) error
	GetMapOptions(mapname string) (string, bool)
	DeleteMapOptions(mapname string) (bool, error)

	SaveGoto(mapname, jumpname string, posX, posY, posZ, angleV, angleH float64) error
	GetGoto(mapname, jumpname string) (GotoData, bool)
	GetGotoNames(mapname string) ([]string, error)
	DeleteGoto(mapname, jumpname string) (bool, error)
	DeleteAllGotos(mapname string) (int, error)

	AddIgnore(guid, ignoredGuid string) error
	RemoveIgnore(guid, ignoredGuid string) error
	GetIgnoredGuids(guid string) ([]string, error)
	GetIgnoredPlayers(guid string) ([]IgnoredPlayer, error)

	UpsertPreferences(guid string, commands []string) error
	GetPreferences(guid string) (commands []string, found bool, err error)

	AddBan(guid, ip, reason string) error
	GetBan(guid, ip string) (reason string, banned bool, err error)
	RemoveBan(playerDbId int) error
	GetBans() ([]BanEntry, error)

	GetQuoteById(id int) (int, string, error)
	GetRandomQuote() (int, string, error)
	SearchQuotes(search string) ([]QuoteEntry, error)
	SaveQuote(text string) error
	DeleteQuote(id int) error

	LookupPlayers(search string) ([]LookupResult, error)
	GetPlayerById(id int) (LookupResult, bool)
	GetPlayersByIp(ip string) ([]LookupResult, error)

	RegisterServer(ip string, port int, rconpassword string, channelId int64, name string) error
}

type IgnoredPlayer struct {
	Id   int
	Name string
	Guid string
}

type QuoteEntry struct {
	Id   int
	Text string
}

type BanEntry struct {
	Id   int
	Name string
}

type TopCheckpoint struct {
	Name        string
	Runtime     int
	Checkpoints []int
}

type LookupResult struct {
	Id      int
	Name    string
	Aliases string
	Role    int
	Ip      string
	Guid    string
}

type GotoData struct {
	PosX   float64
	PosY   float64
	PosZ   float64
	AngleV float64
	AngleH float64
}
