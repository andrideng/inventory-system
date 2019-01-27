package main

import (
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"

	"github.com/andrideng/inventory-system/apis"
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/daos"
	"github.com/andrideng/inventory-system/errors"
	"github.com/andrideng/inventory-system/services"
	routing "github.com/go-ozzo/ozzo-routing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Sirupsen/logrus"
)

func main() {
	// - load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// - load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// - create the logger
	logger := logrus.New()

	// - connect to the database
	db, err := dbx.MustOpen(app.Config.Dialect, app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// - wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// - start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/api")

	// - set basic route
	rg.Get("/", func(c *routing.Context) error {
		return c.Write("Welcome To Inventory System API " + app.Version)
	})
	// - set health check route
	rg.Get("/ping", func(c *routing.Context) error {
		return c.Write("PONG!")
	})

	// - products endpoint
	productDAO := daos.NewProductDAO()
	apis.ServerProductResource(rg, services.NewProductService(productDAO))

	// - incoming goods endpoint
	incomingGoodsDAO := daos.NewIncomingGoodsDAO()
	apis.ServerIncomingGoodsResource(rg, services.NewIncomingGoodsService(incomingGoodsDAO, productDAO))

	// - outgoing goods endpoint
	outgoingGoodsDAO := daos.NewOutgoingGoodsDAO()
	apis.ServerOutgoingGoodsResource(rg, services.NewOutgoingGoodsService(outgoingGoodsDAO, productDAO))

	// - reports endpoint
	apis.ServerReportResource(rg, services.NewReportService(productDAO, outgoingGoodsDAO))

	// - import csv endping
	apis.ServerImportCsvResource(rg, services.NewImportCsvService(productDAO, incomingGoodsDAO))

	return router
}
