package server

import (
	"main/api"
	"main/config"

	"github.com/gin-gonic/gin"
)

var (
	APIv1           *gin.RouterGroup
	registerApiDocs func(routter *gin.RouterGroup)
)

func registerRoutes(router *gin.Engine, config *config.Configuration) {
	router.RedirectTrailingSlash = true

	api.CreateSession(APIv1)
}
