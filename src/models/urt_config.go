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
	ChannelId      int64
	ServerName     string
	BridgeUrl      string
	BridgeApiKey   string
}

type UrtConfig struct {
	ServerConfig       ServerConfig
	BasePath           string
	DownloadPath       string
	DemoPath           string
	MapRepository      string
	LogFile            string
	WorkerNumber       int
	ApiConfig          ApiConfig
	DbUri              string
	ResetOptions       []string
	PortGotoPath       string
	PortMapOptionsPath string
	WelcomeMessage         string
	DailyPbPenCoinLimit    int
	SchemaPath             string
	TranslateUrl   string
	TranslateLangs []string
}

func (u *UrtConfig) LoadConfig() {
	// Defaults
	viper.SetDefault("serverport", "27960")
	viper.SetDefault("botWorkerNumber", 1)
	viper.SetDefault("dailyPbPenCoinLimit", 2)
	viper.SetDefault("schemaPath", "./sqlc/postgres/schema.sql")

	viper.SetDefault("translateLangs", []string{"fr", "en", "es", "it", "de"})
	viper.SetDefault("resetOptions", []string{
		"sv_fps 125",
		"g_maxGameClients 0",
		"g_oldtriggers 0",
		"g_gear QS",
		"g_allownoclip 1",
		"g_flagreturntime 0",
		"g_nodamage 1",
		"g_novest 1",
		"g_gravity 800",
	})

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
	viper.BindEnv("downloadPath", "downloadPath")
	viper.BindEnv("demoPath", "demoPath")
	viper.BindEnv("botWorkerNumber", "botWorkerNumber")
	viper.BindEnv("discordWebhook", "discordWebhook")
	viper.BindEnv("channelId", "channelId")
	viper.BindEnv("serverName", "serverName")
	viper.BindEnv("bridgeUrl", "bridgeUrl")
	viper.BindEnv("bridgeApiKey", "bridgeApiKey")
	viper.BindEnv("resetOptions", "resetOptions")
	viper.BindEnv("portGotoPath", "portGotoPath")
	viper.BindEnv("portMapOptionsPath", "portMapOptionsPath")
	viper.BindEnv("welcomeMessage", "welcomeMessage")
	viper.BindEnv("dailyPbPenCoinLimit", "dailyPbPenCoinLimit")
	viper.BindEnv("schemaPath", "schemaPath")
	viper.BindEnv("translateUrl", "translateUrl")
	viper.BindEnv("translateLangs", "translateLangs")

	u.BasePath = viper.GetString("urtPath")
	if u.BasePath != "" {
		path := strings.TrimSuffix(u.BasePath, "/")
		u.DownloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		u.DemoPath = fmt.Sprintf("%s/%s", path, "q3ut4/serverdemos")
	}

	if v := viper.GetString("downloadPath"); v != "" {
		u.DownloadPath = v
	}
	if v := viper.GetString("demoPath"); v != "" {
		u.DemoPath = v
	}
	u.MapRepository = viper.GetString("urtRepo")

	u.ServerConfig.Ip = viper.GetString("serverip")
	u.ServerConfig.Port = viper.GetString("serverport")
	u.ServerConfig.Password = viper.GetString("password")

	u.ApiConfig.Url = viper.GetString("ujmUrl")
	u.ApiConfig.ApiKey = viper.GetString("ujmApiKey")
	u.ApiConfig.DiscordWebhook = viper.GetString("discordWebhook")
	u.ApiConfig.ChannelId = viper.GetInt64("channelId")
	u.ApiConfig.BridgeUrl = viper.GetString("bridgeUrl")
	u.ApiConfig.BridgeApiKey = viper.GetString("bridgeApiKey")
	u.ApiConfig.ServerName = viper.GetString("serverName")
	if u.ApiConfig.ServerName == "" {
		u.ApiConfig.ServerName = "Server"
	}

	u.LogFile = viper.GetString("logFilePath")
	u.DbUri = viper.GetString("dbUri")
	u.PortGotoPath = viper.GetString("portGotoPath")
	u.PortMapOptionsPath = viper.GetString("portMapOptionsPath")
	u.WelcomeMessage = viper.GetString("welcomeMessage")
	u.DailyPbPenCoinLimit = viper.GetInt("dailyPbPenCoinLimit")
	u.SchemaPath = viper.GetString("schemaPath")
	u.TranslateUrl = viper.GetString("translateUrl")
	u.TranslateLangs = viper.GetStringSlice("translateLangs")

	u.ResetOptions = parseStringSliceOption(viper.Get("resetOptions"))

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

func parseStringSliceOption(raw any) []string {
	var items []string
	switch v := raw.(type) {
	case []string:
		items = v
	case string:
		items = strings.Split(v, ",")
	case []interface{}:
		for _, item := range v {
			if s, ok := item.(string); ok {
				items = append(items, s)
			}
		}
	}
	result := make([]string, 0, len(items))
	for _, opt := range items {
		if trimmed := strings.TrimSpace(opt); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
