package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

// PostRegister - регистрация нового пользователя
func (c *Controller) PostRegister(ctx *gin.Context) {
	var req view.PostRegistrationRequestData
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	_, err := c.uc.UserRegistration(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(c.HandleError(ctx, err))
		return
	}

	ctx.JSON(http.StatusCreated, view.RegistrationResult{
		Data: view.RegistrationResultData{
			Result: true,
		},
	})
}
