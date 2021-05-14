package shop

import (
	"atlas-nss/json"
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

		result := &DataContainer{
			Data: DataBody{
				Id:   "",
				Type: "",
				Attributes: Attributes{
					NPC: uint32(npcId),
					Items: []ItemAttributes{{
						ItemId:   3990000,
						Price:    500,
						Pitch:    0,
						Position: 1,
					}},
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
