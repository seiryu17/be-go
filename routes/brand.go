package routes

import (
	"go-voucher-api/models"
	"go-voucher-api/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateBrand(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var brand models.Brand
		if err := c.BindJSON(&brand); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		if err := validation.Validate.Struct(&brand); err != nil {
			validationErrors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors[err.Field()] = err.Tag()
			}

			c.JSON(400, gin.H{
				"error":   "Validation failed",
				"details": validationErrors,
			})
			return
		}

		if err := db.Create(&brand).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create brand"})
			return
		}

		c.JSON(201, brand)
	}
}
