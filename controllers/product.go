package controllers

import (
	"net/http"

	"github.com/fatihrizqon/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func GetProductById(c *gin.Context) {
	var products models.Product

	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "record not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func UpdateProductById(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unable to update selected record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "selected record has been updated",
		"data":    product,
	})
}

func DeleteProductById(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "record not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unable to delete selected record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "selected record has been deleted",
	})
}
