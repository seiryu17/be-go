package routes

import (
	"go-voucher-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRedemption(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var redemption models.Transaction
		if err := c.BindJSON(&redemption); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		totalPoints := 0
		for _, voucher := range redemption.Vouchers {
			totalPoints += voucher.CostInPoint
		}
		redemption.TotalPoints = totalPoints

		db.Create(&redemption)
		c.JSON(201, redemption)
	}
}

func GetTransactionDetail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.DefaultQuery("transactionId", "")
		var transaction models.Transaction
		if err := db.First(&transaction, transactionID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(200, transaction)
	}
}
