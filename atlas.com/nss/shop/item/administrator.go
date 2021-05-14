package item

import (
	"gorm.io/gorm"
)

func create(db *gorm.DB, shopId uint32, itemId uint32, price uint32, pitch uint32, position uint32) error {
	a := &entity{
		ShopId: shopId,
		ItemId: itemId,
		Price: price,
		Pitch: pitch,
		Position: position,
	}
	return db.Create(a).Error
}

func deleteAll(db *gorm.DB) error {
	return db.Exec("DELETE FROM shop_items").Error
}