package mysql

import "discount/models"

func (db *MysqlDatabase) DiscountUsageGet(discountId int, discounts *[]models.DiscountUsage) error {
	return db.db.Where("discount_id = ?", discountId).Find(discounts).Error
}

func (db *MysqlDatabase) DiscountUsageAdd(usage *models.DiscountUsage) error {
	return db.db.Create(usage).Error
}
