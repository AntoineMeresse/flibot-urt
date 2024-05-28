package db

type DataPersister interface {
	Close() 
	SaveNewPlayer(name string, guid string, ip_address string) error
	UpdatePlayer() error
}



