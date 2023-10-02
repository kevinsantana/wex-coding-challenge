package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/wex-coding-challenge/internal/config"
)

func Run() {

	r := Router()

	ListenAndServe(r)
}

// ListenAndServe starts http server and handles graceful shutdown
func ListenAndServe(srv *fiber.App) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		log.Info("Gracefully shutting down...")

		_ = srv.Shutdown()
	}()

	log.WithField("host", config.Configuration.Host).
		WithField("port", config.Configuration.Port).
		Info("WEX TAG and Gateways Product server started")

	srvHost := net.JoinHostPort(config.Configuration.Host, config.Configuration.Port)

	if err := srv.Listen(srvHost); err != nil && err != http.ErrServerClosed {
		log.Panicf("server error: %v", err)
	}
}
