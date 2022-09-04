package models

type DiscountUsage struct {
	UserID     int
	User       User `gorm:"foreignkey:UserID"`
	DiscountID int
	Discount   Discount `gorm:"foreignkey:DiscountID"`
}
