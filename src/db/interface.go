package db

import "time"

type PenData struct {
	Name string    
	Size float64   
	Date time.Time 
}

type DataPersister interface {
	Close() 
	SaveNewPlayer(name string, guid string, ip_address string) error
	UpdatePlayer() error
	Pen_add(guid string, size float64) error
	Pen_PenOfTheDay() ([]PenData, error)
}



