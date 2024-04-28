package controller

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	itemsLimit, _ := strconv.Atoi(c.Query("itemsLimit"))

	response, err := products.PaginateList(page, itemsLimit)
	if err != nil {
		switch typedErr := err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": typedErr.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": typedErr.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func ShowProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := products.Datails(id)

	if err != nil {
		switch typedErr := err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": typedErr.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": typedErr.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := products.Create(params, true)

	if err != nil {
		switch typedErr := err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": typedErr.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": typedErr.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	params.ID = id

	product, err := products.Update(params, true)

	if err != nil {
		switch typedErr := err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": typedErr.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": typedErr.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
