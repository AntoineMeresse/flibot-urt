package models

import (
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

type Server struct {
	Players []Player
	Rcon quake3_rcon.Rcon
}