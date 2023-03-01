package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func Logout(c *gin.Context) {
	c.String(http.StatusOK, "logout ok")
}