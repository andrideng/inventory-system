package apis

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// outgoingGoodsService specifies the interface for outgoing goods service needed by outgoingGoodsResource
	outgoingGoodsService interface {
		Get(rs app.RequestScope, id int64) (*models.OutgoingGoods, error)
		List(rs app.RequestScope) ([]models.OutgoingGoods, error)
		Create(rs app.RequestScope, model *models.OutgoingGoods) (*models.OutgoingGoods, error)
	}

	// outgoingGoodsResource defines the handlers for the CRUD APIs.
	outgoingGoodsResource struct {
		service outgoingGoodsService
	}
)

// ServerOutgoingGoodsResource sets up the routing outgoing goods endpoints and the corresponding handlers.
func ServerOutgoingGoodsResource(rg *routing.RouteGroup, service outgoingGoodsService) {
	r := &outgoingGoodsResource{service}
	rg.Get("/outgoing-goods", r.list)
	rg.Post("/outgoing-goods", r.create)
}

func (r *outgoingGoodsResource) list(c *routing.Context) error {
	response, err := r.service.List(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}

func (r *outgoingGoodsResource) create(c *routing.Context) error {
	var model models.OutgoingGoods
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}
