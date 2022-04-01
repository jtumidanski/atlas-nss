package item

import (
	"atlas-nss/database"
	"atlas-nss/model"
	"gorm.io/gorm"
)

func getByShopId(shopId uint32) database.EntitySliceProvider[entity] {
	return func(db *gorm.DB) model.SliceProvider[entity] {
		return database.SliceQuery[entity](db, &entity{ShopId: shopId})
	}
}

func makeItem(e entity) (Model, error) {
	return Model{
		id:       e.ID,
		shopId:   e.ShopId,
		itemId:   e.ItemId,
		price:    e.Price,
		pitch:    e.Pitch,
		position: e.Position,
	}, nil
}
