package mysql

import (
	"discount/models"
)

func (db *MysqlDatabase) DiscountAdd(discount *models.Discount) error {
	return db.db.Create(discount).Error
}

func (db *MysqlDatabase) DiscountUpdate(discount *models.Discount) error {
	return db.db.Save(discount).Error
}

func (db *MysqlDatabase) DiscountGet(discount *models.Discount) error {
	return db.db.Find(discount, discount).Error
}
