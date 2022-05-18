package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a App) HealthCheck(c *gin.Context) {

	c.String(http.StatusOK, "healthy")

}
