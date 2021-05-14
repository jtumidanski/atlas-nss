package shop

import (
	"atlas-nss/json"
	"atlas-nss/shop/item"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetShop(fl logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := fl.WithFields(logrus.Fields{"originator": "HandleGetKeyMap", "type": "rest_handler"})

		npcId, err := strconv.Atoi(mux.Vars(r)["npcId"])
		if err != nil {
			l.WithError(err).Errorf("Unable to properly parse npcId from path.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		items, err := item.GetByShopId(db)(uint32(npcId))
		if err != nil {
			l.WithError(err).Errorf("Unable to get the shops items.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		is := make([]ItemAttributes, 0)
		for _, i := range items {
			is = append(is, ItemAttributes{
				ItemId:   i.ItemId(),
				Price:    i.Price(),
				Pitch:    i.Pitch(),
				Position: i.Position(),
			})
		}

		result := &DataContainer{
			Data: DataBody{
				Id:   "",
				Type: "",
				Attributes: Attributes{
					NPC: uint32(npcId),
					Items: is,
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		err = json.ToJSON(result, w)
		if err != nil {
			l.WithError(err).Errorf("Writing response.")
		}
	}
}
