package controllers

import (
	"net/http"

	"github.com/martinyonathann/bookstore_items-api/services"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	flagActive := c.Query("flagActive")

	items, err := services.GetAll(flagActive)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, items)
}
