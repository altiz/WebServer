package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func indexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)

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
	logs.WithFields(logs.Fields{
		"User": json.User,
	}).Info("Ok")
	c.JSON(http.StatusOK, gin.H{
		"User": json.User,
	})

}

func InitializeRoutes(router *gin.Engine) {
	version1 := router.Group("/v1")

	version1.POST("/test", testJSON)
	version1.GET("/", indexPage)
}
