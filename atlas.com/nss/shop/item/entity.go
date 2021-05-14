package item

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&entity{})
}

type entity struct {
	ID       uint32 `gorm:"primaryKey;autoIncrement;not null"`
	ShopId   uint32 `gorm:"not null"`
	ItemId   uint32 `gorm:"not null"`
	Price    uint32 `gorm:"not null"`
	Pitch    uint32 `gorm:"not null"`
	Position uint32 `gorm:"not null"`
}

func (e entity) TableName() string {
	return "shopItems"
}
