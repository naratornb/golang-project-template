package config

import (
	cm "cmd/internal/common"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"kline-uat-git001-sys.kasikornline.com/go-packages/klvault"
	"kline-uat-git001-sys.kasikornline.com/go-packages/zapinstance"
)

type Config struct {
	SFTP   cm.SFTP
	Client cm.HttpClient
}

func InitConfig() *Config {
	initEnvFile()
	initLogsAndVault()
	return &Config{
		SFTP:   cm.NewSFTP(),
		Client: cm.NewHttpClient(),
	}
}

func initEnvFile() {
	err := godotenv.Load()
	if err != nil {
		zap.S().Errorf("cannot load evn file: %v", err)
		log.Fatal(err)
	}
}

func initLogsAndVault() {
	logs := zapinstance.Init(os.Getenv("APP_ENV"))
	zapinstance.ReplaceGlobals(logs)
	err := klvault.ReadSecret()
	if err != nil {
		zap.S().Errorf("cannot read secret from vault %v", err)
		log.Fatal(err)
	}
}
