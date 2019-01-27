package daos

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
)

// IncomingGoodsDAO presists incoming goods data in database
type IncomingGoodsDAO struct{}

// NewIncomingGoodsDAO creates a new IncomingGoodsDAO
func NewIncomingGoodsDAO() *IncomingGoodsDAO {
	return &IncomingGoodsDAO{}
}

// Get reads the incoming goods with specified id from the database.
func (dao *IncomingGoodsDAO) Get(rs app.RequestScope, id int64) (*models.IncomingGoods, error) {
	var incomingGoods models.IncomingGoods
	err := rs.Tx().Select().Model(id, &incomingGoods)
	return &incomingGoods, err
}

// List retrives incoming goods record from the database.
func (dao *IncomingGoodsDAO) List(rs app.RequestScope) ([]models.IncomingGoods, error) {
	incomingGoods := []models.IncomingGoods{}
	err := rs.Tx().Select().All(&incomingGoods)
	return incomingGoods, err
}

// Create saves a new incoming goods record in the database.
// The IncomingGoods.ID will be populated with a generated ID upon successful saving.
func (dao *IncomingGoodsDAO) Create(rs app.RequestScope, incomingGoods *models.IncomingGoods) error {
	return rs.Tx().Model(incomingGoods).Insert()
}

// Update saves the changes to a incoming goods in the database.
func (dao *IncomingGoodsDAO) Update(rs app.RequestScope, id int64, incomingGoods *models.IncomingGoods) error {
	return rs.Tx().Model(incomingGoods).Exclude("ID").Update()
}
