package server

import (
	"context"
	"fmt"
	"main/config"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func StartServer(ctx context.Context, addr string, port int) {
	router := gin.New()
	router.Use(gin.Recovery())

	server := &http.Server{
		Handler: router,
	}

	tcpSocket := fmt.Sprintf("%s:%d", addr, port)

	var listener net.Listener
	var err error

	if listener, err = net.Listen("tcp", tcpSocket); err != nil {
		return
	}

	if err := server.Serve(listener); err != nil {
		return
	}

	<-ctx.Done()
	log.Info(ctx, "Shutting down HTTP server")

	if err := server.Close(); err != nil {
		log.Errorf("Shutdown failed (%s)", err)
	}
}

func registerRoutes(router *gin.Engine, config *config.Configuration) {

}

func registerWebAppRoutes(router *gin.Engine, config *config.Configuration) {

}
