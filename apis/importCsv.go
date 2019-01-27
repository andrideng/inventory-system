package apis

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	importCsvService interface {
		ImportProducts(rs app.RequestScope) ([]models.Product, error)
		ImportIncomingGoods(rs app.RequestScope) ([]models.IncomingGoods, error)
	}
	importCsvResource struct {
		service importCsvService
	}
)

// ServerImportCsvResource sets up the routing of import csv endpoints.
func ServerImportCsvResource(rg *routing.RouteGroup, service importCsvService) {
	r := &importCsvResource{service}
	rg.Get("/import-products", r.importProduct)
	rg.Get("/import-incoming-goods", r.importIncomingGoods)
}

func (r *importCsvResource) importProduct(c *routing.Context) error {
	response, err := r.service.ImportProducts(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}

func (r *importCsvResource) importIncomingGoods(c *routing.Context) error {
	response, err := r.service.ImportIncomingGoods(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}
