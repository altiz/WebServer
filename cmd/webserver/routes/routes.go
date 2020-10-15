package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sayhello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})

}
func SayGo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "go",
	})

}
func testJSON(c *gin.Context) {
	type data struct {
		User string `json:"user"`
	}

	var json data
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hello": json.User,
	})

}

func InitializeRoutes(router *gin.Engine) {
	version1 := router.Group("/v1")

	version1.POST("/test", testJSON)
}
