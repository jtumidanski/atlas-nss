package item

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

func Initialize(l logrus.FieldLogger, db *gorm.DB) {
	d, e := os.LookupEnv("ITEM_DIR")
	if !e {
		d = "/data/items"
	}
	items, err := readDataDirectory(l, d)
	if err != nil {
		l.Fatal(err.Error())
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err := deleteAll(tx)
		if err != nil {
			l.WithError(err).Errorf("Unable to truncate shop items for initialization.")
		}

		for _, i := range items {
			err := create(tx, i.ShopId, i.ItemId, i.Price, i.Pitch, i.Position)
			if err != nil {
				l.WithError(err).Errorf("Unable to insert item %d for shop %d.", i.ItemId, i.ShopId)
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.WithError(err).Errorf("Unable to initialize shop database.")
	}
}

func GetByShopId(db *gorm.DB) func(shopId uint32) ([]*Model, error) {
	return func(shopId uint32) ([]*Model, error) {
		return getByShopId(db, shopId)
	}
}