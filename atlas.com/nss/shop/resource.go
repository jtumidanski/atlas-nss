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
	getShop = "get_shop"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	nsr := router.PathPrefix("/npcs").Subrouter()
	nsr.HandleFunc("/{npcId}/shop", registerGetShop(l, db)).Methods(http.MethodGet)
}

type npcIdHandler func(npcId uint32) http.HandlerFunc

func parseNpcId(l logrus.FieldLogger, next npcIdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		monsterId, err := strconv.Atoi(vars["npcId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing npcId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(uint32(monsterId))(w, r)
	}
}

func registerGetShop(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(getShop, func(span opentracing.Span) http.HandlerFunc {
		return parseNpcId(l, func(npcId uint32) http.HandlerFunc {
			return handleGetShop(l, db)(span)(npcId)
		})
	})
}

func handleGetShop(l logrus.FieldLogger, db *gorm.DB) func(span opentracing.Span) func(npcId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(npcId uint32) http.HandlerFunc {
		return func(npcId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				items, err := item.GetByShopId(db)(npcId)
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
							NPC:   npcId,
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
}
