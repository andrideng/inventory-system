package daos

import (
	"testing"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/testdata"
	"github.com/stretchr/testify/assert"
)

func TestProductDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewProductDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			product, err := dao.Get(rs, "ABCD")
			assert.Nil(t, err)
			if assert.NotNil(t, product) {
				assert.Equal(t, "ABCD", product.SKU)
			}
		})
	}
}
