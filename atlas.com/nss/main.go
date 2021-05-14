package main

import (
	"atlas-nss/database"
	"atlas-nss/logger"
	"atlas-nss/rest"
	"atlas-nss/shop/item"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l := logger.CreateLogger()

	db := database.ConnectToDatabase(l)

	item.Initialize(l, db)

	rest.CreateRestService(l, db)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
}
