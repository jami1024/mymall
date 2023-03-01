package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {}

type LoginResponse struct {}

func Login(c *gin.Context) {
	req := LoginRequest{}
	c.ShouldBindJSON(&req)


	// code
	resp := LoginResponse{}
	c.JSON(http.StatusOK, &resp)
}