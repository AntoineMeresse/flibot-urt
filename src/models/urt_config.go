package models

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Ip       string
	Port     string
	Password string
}

type ApiConfig struct {
	Url            string
	ApiKey         string
	DiscordWebhook string
}

type UrtConfig struct {
	ServerConfig  ServerConfig
	BasePath      string
	DownloadPath  string
	GotosPath     string
	DemoPath      string
	MapRepository string
	LogFile       string
	WorkerNumber  int
	ApiConfig     ApiConfig
	DbUri         string
}

func (u *UrtConfig) LoadConfig() {
	// Defaults
	viper.SetDefault("serverport", "27960")
	viper.SetDefault("botWorkerNumber", 1)

	// Config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")
	if err := viper.ReadInConfig(); err != nil {
		log.Debugf("No config file found: %s", err)
	}

	// Env variables (override config file)
	viper.AutomaticEnv()

	// Bind env
	viper.BindEnv("dbUri", "dbUri")
	viper.BindEnv("serverip", "serverip")
	viper.BindEnv("serverport", "serverport")
	viper.BindEnv("password", "password")
	viper.BindEnv("logFilePath", "logFilePath")
	viper.BindEnv("urtRepo", "urtRepo")
	viper.BindEnv("ujmUrl", "ujmUrl")
	viper.BindEnv("ujmApiKey", "ujmApiKey")
	viper.BindEnv("urtPath", "urtPath")
	viper.BindEnv("botWorkerNumber", "botWorkerNumber")
	viper.BindEnv("discordWebhook", "discordWebhook")

	u.BasePath = viper.GetString("urtPath")
	if u.BasePath != "" {
		path := strings.TrimSuffix(u.BasePath, "/")
		u.DownloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		u.GotosPath = fmt.Sprintf("%s/%s", path, "q3ut4/gotos")
		u.DemoPath = fmt.Sprintf("%s/%s", path, "q3ut4/serverdemos")
	}
	u.MapRepository = viper.GetString("urtRepo")

	u.ServerConfig.Ip = viper.GetString("serverip")
	u.ServerConfig.Port = viper.GetString("serverport")
	u.ServerConfig.Password = viper.GetString("password")

	u.ApiConfig.Url = viper.GetString("ujmUrl")
	u.ApiConfig.ApiKey = viper.GetString("ujmApiKey")
	u.ApiConfig.DiscordWebhook = viper.GetString("discordWebhook")

	u.LogFile = viper.GetString("logFilePath")
	u.DbUri = viper.GetString("dbUri")

	log.Info("Db uri: ", u.DbUri)
	log.Info("Direct env: ", os.Getenv("dbUri"))

	u.initWorkerNumber()
}

func (u *UrtConfig) initWorkerNumber() {
	value := viper.GetInt("botWorkerNumber")

	if value <= 0 || value >= 100 {
		log.Error("Please specify a number between 1 & 99 for botWorkerNumber")
		u.WorkerNumber = 1
		return
	}

	u.WorkerNumber = value
	if value != 1 {
		log.Debugf("Worker number has been modify in configuration to: %d (Default: 1)", value)
	}
}

func (s ServerConfig) GetServerUrl() string {
	return fmt.Sprintf("%s:%s", s.Ip, s.Port)
}
