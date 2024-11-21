package routes

import (
	"fmt"
	"go-voucher-api/models"
	"go-voucher-api/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CreateVoucher(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var voucher models.Voucher

		if err := c.BindJSON(&voucher); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		fmt.Printf("Parsed Voucher: %+v\n", voucher)
		if err := validation.Validate.Struct(&voucher); err != nil {
			validationErrors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors[err.Field()] = err.Tag()
			}

			fmt.Printf("Validation Errors: %+v\n", validationErrors)

			c.JSON(400, gin.H{
				"error":   "Validation failed",
				"details": validationErrors,
			})
			return
		}

		if err := db.Create(&voucher).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create voucher"})
			return
		}

		c.JSON(201, voucher)
	}
}

func GetVoucher(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		voucherID := c.DefaultQuery("id", "")
		var voucher models.Voucher
		if err := db.First(&voucher, voucherID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Voucher not found"})
			return
		}
		c.JSON(200, voucher)
	}
}

func GetVouchersByBrand(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		brandID := c.DefaultQuery("id", "")
		var vouchers []models.Voucher
		if err := db.Where("brand_id = ?", brandID).Find(&vouchers).Error; err != nil {
			c.JSON(404, gin.H{"error": "No vouchers found for this brand"})
			return
		}
		c.JSON(200, vouchers)
	}
}
