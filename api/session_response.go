package api

import (
	"main/config"

	"github.com/gin-gonic/gin"
)

func GetSessionResponse(authToken string, config *config.Configuration) gin.H {
	if authToken == "" {
		return gin.H{}
	} else {
		return gin.H{}
	}
}
