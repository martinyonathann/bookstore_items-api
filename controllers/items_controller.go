package controllers

import (
	"net/http"
	"strconv"

	"github.com/martinyonathann/bookstore_items-api/domain/items"
	"github.com/martinyonathann/bookstore_items-api/services"
	"github.com/martinyonathann/bookstore_items-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func getBookId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid email address")
	}
	return userID, nil
}
func GetAllBook(c *gin.Context) {
	flagActive := c.Query("flagActive")

	items, err := services.ItemsService.GetAll(flagActive)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, items)
}
func GetBookById(c *gin.Context) {
	bookID, idErr := getBookId(c.Param("items_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	result, getErr := services.ItemsService.GetItemByID(bookID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, result)
}
func CreateBook(c *gin.Context) {
	var book items.Item
	if err := c.ShouldBindJSON(&book); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.ItemsService.CreateBook(book)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
