package models

type (
	// ValueOfGoods represents a value of goods for reporting
	ValueOfGoods struct {
		SKU                  string `json:"sku"`
		Name                 string `json:"name"`
		Amount               int64  `json:"amount"`
		AveragePurchasePrice int64  `json:"average_purhcase_price"`
		Total                int64  `json:"total"`
	}

	// ResponseValueOfGoods represents the response
	ResponseValueOfGoods struct {
		ValueOfGoods []ValueOfGoods `json:"value_of_goods"`
		PrintDate    string         `json:"print_date"`
		TotalSKU     int64          `json:"total_sku"`
		TotalProduct int64          `json:"total_product"`
		TotalValue   int64          `json:"total_value"`
	}
)
