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
func InitializeRoutes(router *gin.Engine) {
	router.GET("/", Sayhello)
	router.GET("/go", Sayhello)
}
