package rest

import (
	"atlas-nss/shop"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes(db))
}

func ProduceRoutes(db *gorm.DB) func(l logrus.FieldLogger) http.Handler {
	return func(l logrus.FieldLogger) http.Handler {
		router := mux.NewRouter().PathPrefix("/ms/nss").Subrouter().StrictSlash(true)
		router.Use(CommonHeader)

		nsr := router.PathPrefix("/npcs").Subrouter()
		nsr.HandleFunc("/{npcId}/shop", shop.GetShop(l, db)).Methods(http.MethodGet)

		return router
	}
}
