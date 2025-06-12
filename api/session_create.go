package api

import (
	"main/config"

	"github.com/gin-gonic/gin"
)

func CreateSession(router *gin.RouterGroup) {
	var config *config.Config
	config.ClientSession()
}
