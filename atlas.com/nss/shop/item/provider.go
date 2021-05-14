package item

import "gorm.io/gorm"

func getByShopId(db *gorm.DB, shopId uint32) ([]*Model, error) {
	var results []entity
	err := db.Where(&entity{ShopId: shopId}).Find(&results).Error
	if err != nil {
		return make([]*Model, 0), err
	}

	var items []*Model
	for _, i := range results {
		items = append(items, makeItem(&i))
	}
	return items, nil
}

func makeItem(e *entity) *Model {
	return &Model{
		id:       e.ID,
		shopId:   e.ShopId,
		itemId:   e.ItemId,
		price:    e.Price,
		pitch:    e.Pitch,
		position: e.Position,
	}
}
