package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/wex-coding-challenge/internal/config"
	"github.com/kevinsantana/wex-coding-challenge/internal/infra/database"
	"github.com/kevinsantana/wex-coding-challenge/internal/rest"
)

type HttpConfig struct {
	Cfg *config.Config
	Db  *database.Database
}

func Run(ctx context.Context, c HttpConfig) {
	health := rest.InitializeHealthWeb(c.Db)
	r := Router(health)

	ListenAndServe(ctx, r, c.Cfg)
}

// ListenAndServe starts http server and handles graceful shutdown
func ListenAndServe(ctx context.Context, srv *fiber.App, conf *config.Config) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	srvHost := net.JoinHostPort(conf.Server.Host, conf.Server.Port)

	go func() {
		log.WithField("host", conf.Server.Host).
			WithField("port", conf.Server.Port).
			Info("WEX TAG and Gateways Product server started")

		if err := srv.Listen(srvHost); err != nil && err != http.ErrServerClosed {
			log.WithError(err).
				Panicf("server error")
		}
	}()

	<-done

	log.Info("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(ctx, conf.Server.ShutdownTimeout)

	if err := srv.ShutdownWithContext(ctx); err != nil {
		log.WithError(err).Error("Server shutdown failed")
	}

	cancel()

	log.Info("Disconnecting other services")

	// disconnect database open connections

	log.Info("Server exited")

}
