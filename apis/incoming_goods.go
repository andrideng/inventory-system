package apis

import (
	"strconv"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// incomingGoodsService specifies the interface for incoming goods service needed by incomingGoodsResource.
	incomingGoodsService interface {
		Get(rs app.RequestScope, id int64) (*models.IncomingGoods, error)
		List(rs app.RequestScope) ([]models.IncomingGoods, error)
		Create(rs app.RequestScope, model *models.IncomingGoods) (*models.IncomingGoods, error)
		Update(rs app.RequestScope, id int64, model *models.IncomingGoods) (*models.IncomingGoods, error)
	}

	// incomingGoodsResource defines the handlers for the CRUD APIs.
	incomingGoodsResource struct {
		service incomingGoodsService
	}
)

// ServerIncomingGoodsResource sets up the routing of incoming goods endpoints and the corresponding handlers.
func ServerIncomingGoodsResource(rg *routing.RouteGroup, service incomingGoodsService) {
	r := &incomingGoodsResource{service}
	rg.Get("/incoming-goods", r.list)
	rg.Post("/incoming-goods", r.create)
	rg.Put("/incoming-goods/<id>", r.update)
}

func (r *incomingGoodsResource) list(c *routing.Context) error {
	response, err := r.service.List(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	return c.Write(Response{Data: response})
}

func (r *incomingGoodsResource) create(c *routing.Context) error {
	var model models.IncomingGoods
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(Response{Data: response})
}

func (r *incomingGoodsResource) update(c *routing.Context) error {
	id, err := (strconv.Atoi(c.Param("id")))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)
	model, err := r.service.Get(rs, int64(id))
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, int64(id), model)
	if err != nil {
		return err
	}

	return c.Write(Response{Data: response})
}
