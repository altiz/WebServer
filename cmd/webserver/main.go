package main

import (
	routes "WebServer/cmd/webserver/routes"
	"WebServer/version"

	gin "github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

var router *gin.Engine

func setupRouter() *gin.Engine {
	logs.SetFormatter(&logs.JSONFormatter{})

	logs.WithFields(logs.Fields{
		"commit":     version.Commit,
		"build time": version.BuildTime,
		"release":    version.Release,
	}).Info("Starting the service...")

	router := gin.Default()

	//router.LoadHTMLGlob("templates/*")

	routes.InitializeRoutes(router)
	return router
}

func main() {
	router := setupRouter()
	router.Run()

}
