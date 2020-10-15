package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func indexPage(c *gin.Context) {
	c.String(http.StatusOK, "ok")

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
	if json.User == "test" {
		logs.WithFields(logs.Fields{
			"Status": 0,
		}).Info("Ok")
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
		})
	} else {
		logs.WithFields(logs.Fields{
			"User": json.User,
		}).Info("Ok")
		c.JSON(http.StatusOK, gin.H{
			"User": json.User,
		})
	}

}

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	version1 := api.Group("/v1")

	version1.POST("/test", testJSON)
	version1.GET("/", indexPage)
}
