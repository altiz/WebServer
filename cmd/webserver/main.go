package main

import (
	routes "WebServer/cmd/webserver/routes"

	gin "github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

var router *gin.Engine

func setupRouter() *gin.Engine {
	logs.SetFormatter(&logs.JSONFormatter{})

	logs.WithFields(logs.Fields{
		"proc": "main",
	}).Info("service start")

	router := gin.Default()

	//router.LoadHTMLGlob("templates/*")

	logs.WithFields(logs.Fields{
		"proc": "main",
	}).Info("ruter start")
	routes.InitializeRoutes(router)
	return router
}

func main() {
	router := setupRouter()
	router.Run()

}
