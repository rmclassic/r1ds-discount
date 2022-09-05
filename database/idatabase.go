package database

import "discount/models"

type IDatabase interface {
	Init()
	GetForUpdate(interface{}) error

	DiscountAdd(*models.Discount) error
	DiscountGet(*models.Discount) error
	DiscountUpdate(*models.Discount) error

	DiscountUsageGet(models.DiscountUsage, *[]models.DiscountUsage) error
	DiscountUsageAdd(*models.DiscountUsage) error

	UserAdd(*models.User) error
	UserGet(*models.User) error
	UserGetByPhoneNumber(string) (*models.User, error)
}
