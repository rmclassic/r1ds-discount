package mysql

import "discount/models"

func (db *MysqlDatabase) DiscountUsageGet(conds models.DiscountUsage, discounts *[]models.DiscountUsage) error {
	return db.db.Preload("Discount").Preload("User").Where(conds).Find(discounts).Error
}

func (db *MysqlDatabase) DiscountUsageAdd(usage *models.DiscountUsage) error {
	return db.db.Create(usage).Error
}
