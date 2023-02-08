package router

import (
	"cmd/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	pathHealth   = "/batch/ping"
	pathMonitor  = "/batch/monitor"
	pathDownload = "/sftp/download"
	pathUpload   = "/sftp/upload"
)

func InitRouter(service service.Service) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat(pathHealth))
	r.Get(pathMonitor, service.Monitor)
	r.Get(pathDownload, service.Download)
	// r.Post(pathUpload, service.Upload)
	return r
}
