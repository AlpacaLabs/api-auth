package app

import (
	"sync"

	"github.com/AlpacaLabs/go-kontext"

	"github.com/AlpacaLabs/api-auth/internal/grpc"

	"github.com/AlpacaLabs/api-auth/internal/configuration"
	"github.com/AlpacaLabs/api-auth/internal/db"
	"github.com/AlpacaLabs/api-auth/internal/http"
	"github.com/AlpacaLabs/api-auth/internal/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config configuration.Config
}

func NewApp(c configuration.Config) App {
	return App{
		config: c,
	}
}

func (a App) Run() {
	dbConn, err := db.Connect(a.config.SQLConfig)
	if err != nil {
		log.Fatalf("failed to dial database: %v", err)
	}
	dbClient := db.NewClient(dbConn)

	accountConn, err := kontext.Dial(a.config.AccountGRPCAddress)
	if err != nil {
		log.Fatalf("failed to dial Account service: %v", err)
	}

	mfaConn, err := kontext.Dial(a.config.MFAGRPCAddress)
	if err != nil {
		log.Fatalf("failed to dial MFA service: %v", err)
	}

	passwordConn, err := kontext.Dial(a.config.PasswordGRPCAddress)
	if err != nil {
		log.Fatalf("failed to dial Password service: %v", err)
	}

	svc := service.NewService(a.config, dbClient, accountConn, mfaConn, passwordConn)

	var wg sync.WaitGroup

	wg.Add(1)
	httpServer := http.NewServer(a.config, svc)
	httpServer.Run()

	wg.Add(1)
	grpcServer := grpc.NewServer(a.config, svc)
	grpcServer.Run()

	wg.Wait()
}
