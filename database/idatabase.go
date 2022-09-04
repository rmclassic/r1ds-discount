package database

import "discount/models"

type IDatabase interface {
	Init()
	GetForUpdate(interface{}) error

	DiscountAdd(*models.Discount) error
	DiscountGet(*models.Discount) error
	DiscountUpdate(*models.Discount) error
}
