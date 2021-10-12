package shop

import (
	"atlas-nss/json"
	"atlas-nss/rest"
	"atlas-nss/shop/item"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	GetShop = "get_shop"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	nsr := router.PathPrefix("/npcs").Subrouter()
	nsr.HandleFunc("/{npcId}/shop", rest.RetrieveSpan(GetShop, HandleGetShop(l, db))).Methods(http.MethodGet)
}

func HandleGetShop(fl logrus.FieldLogger, db *gorm.DB) rest.SpanHandler {
	return func(span opentracing.Span) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			l := fl.WithFields(logrus.Fields{"originator": GetShop, "type": "rest_handler"})

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

			if len(items) == 0 {
				w.WriteHeader(http.StatusNotFound)
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
						NPC:   uint32(npcId),
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
}
