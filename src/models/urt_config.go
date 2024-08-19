package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Ip string 
	Port string 
	Password string
}

type ApiConfig struct {
	Url string
	ApiKey string
}

type UrtConfig struct {
	ServerConfig ServerConfig
	BasePath string
	DownloadPath string
	GotosPath string
	MapRepository string
	LogFile string
	WorkerNumber int
	ApiConfig ApiConfig
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

	u.ApiConfig.Url = os.Getenv("ujmUrl") 
	u.ApiConfig.ApiKey = os.Getenv("ujmApiKey") 

	u.LogFile = os.Getenv("logFilePath") 
	u.initWorkerNumber()	
}

func (u *UrtConfig) initWorkerNumber() {
	workerValue, found := os.LookupEnv("botWorkerNumber")
	if !found {
		log.Debug("Worker number not specify in conf. Will use default: 1")
		u.WorkerNumber = 1
	}

	value, err := utils.ExtractNumber(workerValue);

	if err == nil && value > 0 && value < 100 {
		if value != 1 {
			u.WorkerNumber = value;
			log.Debugf("Worker number has been modify in configuration to: %d (Default: 1)", value)
		}
	} else {
		log.Error("Please specify a number between 1 & 99 for the env variable: botWorkerNumber")
	}
}
