package serviceAccount

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Profile(c *gin.Context) {
	 c.String(http.StatusOK, "profile ok")
}