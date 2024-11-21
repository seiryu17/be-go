package main

import (
	"go-voucher-api/db"
	"go-voucher-api/routes"
	"go-voucher-api/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db := db.InitDB()

	utils.ApplyMigrations()

	r := gin.Default()

	r.POST("/brand", routes.CreateBrand(db))
	r.POST("/voucher", routes.CreateVoucher(db))
	r.GET("/voucher", routes.GetVoucher(db))
	r.GET("/voucher/brand", routes.GetVouchersByBrand(db))
	r.POST("/transaction/redemption", routes.CreateRedemption(db))
	r.GET("/transaction/redemption", routes.GetTransactionDetail(db))

	r.Run(":8080")
}
