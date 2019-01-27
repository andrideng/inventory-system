package services

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
)

// productDAO specifies the interface of the product DAO needed by ProductService
type productDAO interface {
	// Get returns the product with the specified product SKU.
	Get(rs app.RequestScope, sku string) (*models.Product, error)
	// List return the list of the products.
	List(rs app.RequestScope) ([]models.Product, error)
	// Create saves a new product in the storage.
	Create(rs app.RequestScope, product *models.Product) error
	// Update updates product in the storage.
	Update(rs app.RequestScope, sku string, product *models.Product) (*models.Product, error)
}

// ProductService provides services related with products.
type ProductService struct {
	dao productDAO
}

// NewProductService creates a new ProductService with the given product DAO.
func NewProductService(dao productDAO) *ProductService {
	return &ProductService{dao}
}

// Get retruns the product with the specified the product SKU.
func (s *ProductService) Get(rs app.RequestScope, sku string) (*models.Product, error) {
	return s.dao.Get(rs, sku)
}

// List returns the list of products
func (s *ProductService) List(rs app.RequestScope) ([]models.Product, error) {
	return s.dao.List(rs)
}

// Create creates a new product
func (s *ProductService) Create(rs app.RequestScope, model *models.Product) (*models.Product, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	model.CreatedAt = rs.CurrentDateTime()
	model.UpdatedAt = rs.CurrentDateTime()
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.SKU)
}

// Update updates a product based on sku.
func (s *ProductService) Update(rs app.RequestScope, sku string, model *models.Product) (*models.Product, error) {
	if _, err := s.dao.Get(rs, sku); err != nil {
		return nil, err
	}
	// - updated at
	model.UpdatedAt = rs.CurrentDateTime()
	if _, err := s.dao.Update(rs, sku, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, sku)
}
