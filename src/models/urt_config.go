package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Ip string 
	Port string 
	Password string
}

type UrtConfig struct {
	ServerConfig ServerConfig
	BasePath string
	DownloadPath string
	GotosPath string
	MapRepository string
}

func (u *UrtConfig) loadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		panic("Error trying to load env variables")
	}

	u.BasePath = os.Getenv("urtPath")
	if u.BasePath != "" {
		path := strings.TrimSuffix(u.BasePath, "/")
		u.DownloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		u.GotosPath = fmt.Sprintf("%s/%s", path, "q3ut4/gotos")
	}
	u.MapRepository = os.Getenv("urtRepo")
	
	u.ServerConfig.Ip = os.Getenv("serverip") 
	u.ServerConfig.Port =  os.Getenv("serverport") 
	u.ServerConfig.Password = os.Getenv("password") 
}