package identifications

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model "github.com/javiroberts/key_value_server/webserver/model/identifications"
	service "github.com/javiroberts/key_value_server/webserver/services/indentifications"
	"net/http"
	"strconv"
)

func HandleIdentification(c *gin.Context) {
	var req model.IdentificationRequest
	err := c.BindJSON(&req)
	fmt.Println(req)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid parameters")
		return
	}

	err = service.HandleIdentification(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	c.String(http.StatusOK, "ok")
}

func GetIdentification(c *gin.Context) {
	idType := c.Param("idType")
	if idType == "" {
		c.String(http.StatusBadRequest, "invalid id type")
	}

	idNumber, err := strconv.ParseUint(c.Param("idNumber"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid id number")
	}

	res, err := service.GetIdentification(idType, idNumber)

	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("identification type %s - number %d not found", idType, idNumber))
		return
	}

	c.JSON(http.StatusOK, res)
}
