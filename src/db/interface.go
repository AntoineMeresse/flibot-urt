package db

import (
	"database/sql"
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
	return p.Name.String
}

func (p PenData) GetDate() string {
	return utils.FormatTimeToDate(p.Date)
}

type DataPersister interface {
	Close()
	
	// Player
	SaveNewPlayer(name string, guid string, ip_address string) error
	UpdatePlayer() error

	// Role
	
	// Pen 
	Pen_add(guid string, size float64) error
	Pen_PenOfTheDay() (string, []PenData, error)
	Pen_PenHallOfFame() ([]PenData, error)
	Pen_PenHallOfShame() ([]PenData, error)
}



