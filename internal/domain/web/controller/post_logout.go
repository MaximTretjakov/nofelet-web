package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostLogout - разлогин пользователя
func (c *Controller) PostLogout(ctx *gin.Context) {
	// Временное решение
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
