package services

import (
	"discount/database"
	"discount/models"
)

func SubmitDiscount(db database.IDatabase, discountCode string, phoneNumber string) error {
	discount := models.Discount{
		Code: discountCode,
	}

	wallet, err := GetUserWalletByPhoneNumber(phoneNumber)
	if err != nil {
		return err
	}

	if err := db.GetForUpdate(&discount); err != nil {
		return err
	}

	wallet.Balance += discount.Amount

	err = ChargeUserWallet(wallet.UserID, discount.Amount)
	if err != nil {
		return err
	}

	discount.QuantityLeft--
	err = db.DiscountUpdate(&discount)
	if err != nil {
		return err
	}

	go db.DiscountUsageAdd(&models.DiscountUsage{
		DiscountID: discount.ID,
		Discount:   discount,
		UserID:     wallet.UserID,
		User:       wallet.User,
	})

	return nil
}

func AddDiscount(db database.IDatabase, discount *models.Discount) error {
	return db.DiscountAdd(discount)
}
