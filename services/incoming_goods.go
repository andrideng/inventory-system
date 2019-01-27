package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
)

// incomingGoodsDAO specifies the interface of the incoming goods DAO needed by IncomingGoodsService
type incomingGoodsDAO interface {
	// Get return incoming goods with the specified incoming goods id.
	Get(rs app.RequestScope, id int64) (*models.IncomingGoods, error)
	// List return the list of the incoming goods.
	List(rs app.RequestScope) ([]models.IncomingGoods, error)
	// Create for add a new incoming goods.
	Create(rs app.RequestScope, incomingGoods *models.IncomingGoods) error
	// Update for update a existing incoming goods.
	Update(rs app.RequestScope, id int64, incomingGoods *models.IncomingGoods) error
}

// IncomingGoodsService provides services related with incoming goods.
type IncomingGoodsService struct {
	dao     incomingGoodsDAO
	prodDao productDAO
}

// NewIncomingGoodsService creates an new IncomingGoodsService with the given incoming goods DAO.
func NewIncomingGoodsService(dao incomingGoodsDAO, prodDao productDAO) *IncomingGoodsService {
	return &IncomingGoodsService{dao, prodDao}
}

// Get return the spcified incoming goods
func (s *IncomingGoodsService) Get(rs app.RequestScope, id int64) (*models.IncomingGoods, error) {
	return s.dao.Get(rs, id)
}

// List returns the list of incoming goods
func (s *IncomingGoodsService) List(rs app.RequestScope) ([]models.IncomingGoods, error) {
	return s.dao.List(rs)
}

// Create creates new incoming goods
func (s *IncomingGoodsService) Create(rs app.RequestScope, model *models.IncomingGoods) (*models.IncomingGoods, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	// - find sku
	product, err := s.prodDao.Get(rs, model.SKU)
	if err != nil {
		return nil, err
	}
	// - note
	model.Note = time.Now().Format("02/Jan/2006") + " terima " + strconv.Itoa(int(model.AmountReceived))
	// - status
	if model.AmountOrders == model.AmountReceived {
		model.Status = "COMPLETE"
	} else {
		model.Status = "PENDING"
	}
	// - caluclate total
	model.Total = float64(model.AmountOrders) * model.PurchasePrice
	// - handle created_at, updated_at
	model.CreatedAt = rs.CurrentDateTime()
	model.UpdatedAt = rs.CurrentDateTime()
	// - perform create new incoming goods
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	// - update product stocks
	product.Amount += model.AmountReceived
	product.UpdatedAt = rs.CurrentDateTime()
	// - perform update product
	if _, err := s.prodDao.Update(rs, model.SKU, product); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}

// Update udpates the incoming goods with the specified ID.
func (s *IncomingGoodsService) Update(rs app.RequestScope, id int64, model *models.IncomingGoods) (*models.IncomingGoods, error) {
	if err := model.ValidateUpdate(); err != nil {
		return nil, err
	}
	// - check if ID incoming goods exist
	incomingGoods, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}

	// - check status, return if already complete
	if incomingGoods.Status == "COMPLETE" {
		return nil, errors.New("Incoming Goods is Complete")
	}

	diff := incomingGoods.AmountOrders - incomingGoods.AmountReceived
	if model.AmountReceived > diff {
		return nil, errors.New("Amount Received is greater than must be received number of amount")
	}

	// - update status if complete
	if model.AmountReceived == diff {
		model.Status = "COMPLETE"
	}

	// - update note
	model.Note = incomingGoods.Note + " ; " +
		time.Now().Format("02/Jan/2006") + " terima " +
		strconv.Itoa(int(model.AmountReceived))
	// - update updated_at
	model.UpdatedAt = rs.CurrentDateTime()

	// - update amount received
	model.AmountReceived += incomingGoods.AmountReceived

	// - perform update incoming goods
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}

	// - get product
	product, err := s.prodDao.Get(rs, incomingGoods.SKU)
	if err != nil {
		return nil, err
	}
	// - upate stock
	product.Amount += model.AmountReceived
	product.UpdatedAt = rs.CurrentDateTime()
	// - perform update product
	if _, err := s.prodDao.Update(rs, incomingGoods.SKU, product); err != nil {
		return nil, err
	}

	return s.dao.Get(rs, id)
}
