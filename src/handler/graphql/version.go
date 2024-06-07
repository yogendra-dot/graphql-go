package graphql

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, "1.0")
}