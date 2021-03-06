package handlers

import (
	_ "WebServer/cmd/webserver/docs" // docs is generated by Swag CLI, you have to import it.
	"WebServer/cmd/webserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus" // swagger embed files
)

// indexPage godoc
// @Summary Текущая версия API
// @Produce html
// @Success 200 {body} ok
// @Router / [get]
func IndexPage(c *gin.Context) {
	c.String(http.StatusOK, "ok")

}

// testJSON godoc
// @Summary Retrieves user based on given ID
// @Description получить строку по идентификатору
// @Accept  json
// @Produce json
// Param [param_name] [param_type] [data_type] [required/mandatory] [description]
// @Param q body models.TData_req false "name search by q"
// @Success 200 {object} models.TData_resp
// @Router /test/ [post]
func TestJSON(c *gin.Context) {

	var json models.TData_req
	var resp models.TData_resp
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User == "test" {
		logs.WithFields(logs.Fields{
			"Status": 0,
		}).Info("Ok")
		resp.Status = 0
		c.JSON(http.StatusOK, resp)
	} else {
		logs.WithFields(logs.Fields{
			"User": json.User,
		}).Info("Ok")
		resp.Status = 0
		c.JSON(http.StatusOK, resp)
	}

}
