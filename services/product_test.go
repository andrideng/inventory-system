package services

import (
	"errors"
	"testing"
	"time"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/stretchr/testify/assert"
)

// CREATE MOCK
type mockProductDAO struct {
	records []models.Product
}

func newMockProductDAO() productDAO {
	return &mockProductDAO{
		records: []models.Product{
			{SKU: "ABCD", Name: "Product-01", Amount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{SKU: "EFGH", Name: "Product-02", Amount: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{SKU: "IJKL", Name: "Product-03", Amount: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
}

func (m *mockProductDAO) Get(rs app.RequestScope, sku string) (*models.Product, error) {
	for _, record := range m.records {
		if record.SKU == sku {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockProductDAO) List(rs app.RequestScope) ([]models.Product, error) {
	return m.records, nil
}

func (m *mockProductDAO) Create(rs app.RequestScope, product *models.Product) error {
	m.records = append(m.records, *product)
	return nil
}

// START TEST
func TestNewProductService(t *testing.T) {
	dao := newMockProductDAO()
	s := NewProductService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestProductService_List(t *testing.T) {
	s := NewProductService(newMockProductDAO())
	result, err := s.List(nil)
	if assert.Nil(t, err) {
		assert.Equal(t, 3, len(result))
	}
}

func TestAritstService_Get(t *testing.T) {
	s := NewProductService(newMockProductDAO())
	product, err := s.Get(nil, "EFGH")
	if assert.Nil(t, err) && assert.NotNil(t, product) {
		assert.Equal(t, "Product-02", product.Name)
	}

	product, err = s.Get(nil, "WXYZ")
	assert.NotNil(t, err)
}
