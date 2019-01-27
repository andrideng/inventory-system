package daos

import (
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/andrideng/inventory-system/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// OutgoingGoodsDAO presists outgoing goods data in database
type OutgoingGoodsDAO struct{}

// NewOutgoingGoodsDAO creates a new OutgoingGoodsDAO
func NewOutgoingGoodsDAO() *OutgoingGoodsDAO {
	return &OutgoingGoodsDAO{}
}

// Get reads th outgoing goods with specified id from the database.
func (dao *OutgoingGoodsDAO) Get(rs app.RequestScope, id int64) (*models.OutgoingGoods, error) {
	var outgoingGoods models.OutgoingGoods
	err := rs.Tx().Select().Model(id, &outgoingGoods)
	return &outgoingGoods, err
}

// List retrives outgoing goods record from the database.
func (dao *OutgoingGoodsDAO) List(rs app.RequestScope, params *util.QueryParam) ([]models.OutgoingGoods, error) {
	outgoingGoods := []models.OutgoingGoods{}
	var err error
	if params.StartDate == "" || params.EndDate == "" {
		err = rs.Tx().Select().All(&outgoingGoods)
	} else {
		err = rs.Tx().Select().Where(
			dbx.NewExp(
				"created_at >= {:start_date} and created_at <= {:end_date}",
				dbx.Params{"start_date": params.StartDate, "end_date": params.EndDate},
			),
		).All(&outgoingGoods)

	}

	return outgoingGoods, err
}

// Create saves a new outgoing goods record in the database.
// The OutgoingGoods.ID will be populated with a generated ID upon successful saving.
func (dao *OutgoingGoodsDAO) Create(rs app.RequestScope, outgoingGoods *models.OutgoingGoods) error {
	return rs.Tx().Model(outgoingGoods).Insert()
}
