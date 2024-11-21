package models

type Brand struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" validate:"required,min=3"`
	Vouchers []Voucher `json:"vouchers"`
}

type Voucher struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	BrandID     uint   `json:"brand_id" validate:"required"`
	Brand       Brand  `json:"brand" validate:"-"`
	Name        string `json:"name" validate:"required"`
	CostInPoint int    `json:"cost_in_point" validate:"required"`
}

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CustomerID  uint      `json:"customer_id"`
	Vouchers    []Voucher `json:"vouchers" gorm:"many2many:transaction_vouchers;"`
	TotalPoints int       `json:"total_points"`
}
