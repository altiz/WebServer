package handlers

import (
	_ "WebServer/cmd/webserver/docs" // docs is generated by Swag CLI, you have to import it.
	"WebServer/cmd/webserver/models"
	"net/http"

	logs "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	// swagger embed files
)

// ChangeClientCard godoc
// @Summary Сохраняет данные по делам в биллинговой системе
// @Description Передача данных по делам из Jeffit во внешнюю систему
// @Accept  json
// @Produce json
// Param [param_name] [param_type] [data_type] [required/mandatory] [description]
// @Param q body models.TDelo false "name search by q"
// @Success 200 {object} models.TData_resp
// @Router /changeclientcard/ [post]
func ChangeClientCard(c *gin.Context) {

	var json models.TDelo
	var resp models.TData_resp
	logs.SetFormatter(&logs.JSONFormatter{})

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	logs.WithFields(logs.Fields{
		"Client":      json.Client,
		"Decisions":   json.Decisions,
		"ServiceName": json.ServiceName,
	}).Info("Starting the service...")

	c.JSON(http.StatusOK, resp)
}
