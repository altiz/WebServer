package handlers

import (
	_ "WebServer/cmd/webserver/docs" // docs is generated by Swag CLI, you have to import it.
	"WebServer/cmd/webserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// swagger embed files
)

// HomeHandlers godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Success 200 {object} models.Data_resp
// @Router /home/ [post]
func HomeHandlers(c *gin.Context) {
	var version models.Version
	version.BuildTime = version.BuildTime
	version.Commit = version.Commit
	version.Release = version.Release
	c.JSON(http.StatusOK, version)

}
