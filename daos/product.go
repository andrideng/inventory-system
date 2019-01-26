package daos

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
)

// ProductDAO presists product data in database
type ProductDAO struct{}

// NewProductDAO creates a new ProductDAO
func NewProductDAO() *ProductDAO {
	return &ProductDAO{}
}

// Get reads the product eith specified SKU from the database.
func (dao *ProductDAO) Get(rs app.RequestScope, sku string) (*models.Product, error) {
	var product models.Product
	err := rs.Tx().Select().Model(sku, &product)
	return &product, err
}

// List retrives th product records from the database.
func (dao *ProductDAO) List(rs app.RequestScope) ([]models.Product, error) {
	products := []models.Product{}
	err := rs.Tx().Select().All(&products)
	return products, err
}

// Create saves a new product recrod in the database
// The Product.SKU will be populated with a generated SKU upon successful saving.
func (dao *ProductDAO) Create(rs app.RequestScope, product *models.Product) error {
	return rs.Tx().Model(product).Insert()
}
