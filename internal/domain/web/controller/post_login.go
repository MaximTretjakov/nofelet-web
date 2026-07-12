package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

// PostAuth - авторизация пользователя
func (c *Controller) PostAuth(ctx *gin.Context) {
	var req view.PostRegistrationRequestData
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	token, err := c.uc.CreateToken(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(c.HandleError(err))
		return
	}

	ctx.JSON(http.StatusOK, view.AuthResult{
		Data: view.AuthResultData{
			Token: token,
		},
	})
}
