package main

import (
	"cmd/internal/config"
	"cmd/internal/router"
	"cmd/internal/service"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func main() {
	cfg := config.InitConfig()

	port := os.Getenv("PORT")
	zap.S().Infof("start on port %v", port)
	if err := http.ListenAndServe(port, router.InitRouter(
		service.NewService(cfg.SFTP, cfg.Client))); err != nil {
		panic(err)
	}
}
