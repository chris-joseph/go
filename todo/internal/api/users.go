package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todo/pkg/models"
)

func (a App) Register(c *gin.Context) {
	newUser, err := models.ValidateRegisterRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	err = a.userSvc.CreateAccount(newUser)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	c.String(http.StatusCreated, "")
}

func (a App) Login(c *gin.Context) {
	loginRequest, err := models.ValidateLoginRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	token, err := a.userSvc.Login(loginRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Code, err)
		return
	}

	response := &models.LoginResponse{Token: token}

	c.JSON(http.StatusOK, response)
}
