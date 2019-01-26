package apis

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// productService specifies the interface for the product service method needed by productResource.
	productService interface {
		Get(rs app.RequestScope, sku string) (*models.Product, error)
		List(rs app.RequestScope) ([]models.Product, error)
		Create(rs app.RequestScope, model *models.Product) (*models.Product, error)
	}

	// productResource defines the handlers for the CRUD APIs.
	productResource struct {
		service productService
	}
)

// ServerProductResource sets up the routing of product endpoints and the coresponding handlers.
func ServerProductResource(rg *routing.RouteGroup, service productService) {
	r := &productResource{service}
	rg.Get("/products/<sku>", r.get)
	rg.Get("/products", r.list)
	rg.Post("/products", r.create)
}

func (r *productResource) get(c *routing.Context) error {
	sku := c.Param("sku")
	response, err := r.service.Get(app.GetRequestScope(c), sku)
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}

func (r *productResource) list(c *routing.Context) error {
	response, err := r.service.List(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(Response{Data: response})
}

func (r *productResource) create(c *routing.Context) error {
	var model models.Product
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}
