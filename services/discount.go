package services

import (
	"discount/database"
	"discount/models"
)

func SubmitDiscount(db database.IDatabase, discountCode string, phoneNumber string) error {
	discount := models.Discount{
		Code: discountCode,
	}

	(phoneNumber)

	if err := db.GetForUpdate(&discount); err != nil {
		return err
	}

}
