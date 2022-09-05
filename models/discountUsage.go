package models

type DiscountUsage struct {
	UserID     int      `gorm:"primaryKey"`
	User       User     `gorm:"foreignkey:UserID"`
	DiscountID int      `gorm:"primaryKey"`
	Discount   Discount `gorm:"foreignkey:DiscountID"`
}
