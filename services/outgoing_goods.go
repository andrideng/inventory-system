package services

import (
	"errors"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/andrideng/inventory-system/util"
)

// outgoingGoodsDAO specifies the interface of the outgoing goods DAO nedded by OutogingGoodsService.
type outgoingGoodsDAO interface {
	// Get return outgoing goods with the specified outgoing goods id.
	Get(rs app.RequestScope, id int64) (*models.OutgoingGoods, error)
	// List return the list of the outgoing goods.
	List(rs app.RequestScope, params *util.QueryParam) ([]models.OutgoingGoods, error)
	// Create for add a new outgoing goods
	Create(rs app.RequestScope, outgoingGoods *models.OutgoingGoods) error
}

// OutgoingGoodsService provides services related with outgoing goods.
type OutgoingGoodsService struct {
	dao     outgoingGoodsDAO
	prodDao productDAO
}

// NewOutgoingGoodsService creates a new OutgoingGoodsService with the given outgoing goods DAO.
func NewOutgoingGoodsService(dao outgoingGoodsDAO, prodDao productDAO) *OutgoingGoodsService {
	return &OutgoingGoodsService{dao, prodDao}
}

// Get return the sepicied outgoing goods based on id
func (s *OutgoingGoodsService) Get(rs app.RequestScope, id int64) (*models.OutgoingGoods, error) {
	return s.dao.Get(rs, id)
}

// List returns the list of incoming goods
func (s *OutgoingGoodsService) List(rs app.RequestScope, params *util.QueryParam) ([]models.OutgoingGoods, error) {
	return s.dao.List(rs, params)
}

// Create creates new outgoing goods
func (s *OutgoingGoodsService) Create(rs app.RequestScope, model *models.OutgoingGoods) (*models.OutgoingGoods, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	// - find sku
	product, err := s.prodDao.Get(rs, model.SKU)
	if err != nil {
		return nil, err
	}

	// - validate if outgoing stock is lower than current stock
	if product.Amount < model.Amount {
		return nil, errors.New("Product " + model.SKU + " Stock not enough")
	}

	// - calculate total
	model.Total = float64(model.Amount) * model.Price

	// - created_at, updated_at
	model.CreatedAt = rs.CurrentDateTime()
	model.UpdatedAt = rs.CurrentDateTime()

	// - perform create outgoing goods
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}

	// - update product (deduct stock)
	product.Amount -= model.Amount
	// - update product (updated_at)
	product.UpdatedAt = rs.CurrentDateTime()
	// - perform update product
	if _, err := s.prodDao.Update(rs, model.SKU, product); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}
