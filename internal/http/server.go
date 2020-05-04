package http

import (
	"fmt"
	"net/http"

	"github.com/AlpacaLabs/auth/internal/config"
	"github.com/AlpacaLabs/auth/internal/services"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config  config.Config
	service services.Service
}

func NewServer(config config.Config, service services.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	r := mux.NewRouter()

	addr := fmt.Sprintf(":%d", s.config.HTTPPort)
	log.Infof("Listening for HTTP on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
