package models

type (
	// Sales represents a sales for reporting
	Sales struct {
		OrderID     string `json:"order_id"`
		Time        string `json:"time"`
		SKU         string `json:"sku"`
		ProductName string `json:"product_name"`
		Amount      int64  `json:"amount"`
		SellPrice   int64  `json:"sell_price"`
		Total       int64  `json:"total"`
		BuyPrice    int64  `json:"buy_price"`
		Profit      int64  `json:"profit"`
	}

	// ResponseSales represents response for sales model
	ResponseSales struct {
		Sales      []Sales `json:"sales"`
		PrintDate  string  `json:"print_date"`
		TotalOmzet int64   `json:"total_omzet"`
		NetProfit  int64   `json:"net_profit"`
		TotalSales int64   `json:"total_sales"`
		TotalItems int64   `json:"total_items"`
	}
)
