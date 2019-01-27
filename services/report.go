package services

import (
	"time"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/andrideng/inventory-system/util"
)

// ReportService provides services related with reports.
type ReportService struct {
	productDAO       productDAO
	outgoingGoodsDAO outgoingGoodsDAO
}

// NewReportService creates a new ReportService with the given report DAO.
func NewReportService(productDAO productDAO, outgoingGoodsDAO outgoingGoodsDAO) *ReportService {
	return &ReportService{productDAO, outgoingGoodsDAO}
}

// ValueOfGoods calculate value of goods
func (s *ReportService) ValueOfGoods(rs app.RequestScope) (*models.ResponseValueOfGoods, error) {
	var vog []models.ValueOfGoods
	products, err := s.productDAO.List(rs)
	if err != nil {
		return nil, err
	}
	var totalItems, totalPrice int64
	for _, product := range products {
		total := product.Amount * product.AveragePurchasePrice
		totalItems += product.Amount
		totalPrice += total
		vog = append(vog, models.ValueOfGoods{
			SKU:                  product.SKU,
			Name:                 product.Name,
			Amount:               product.Amount,
			AveragePurchasePrice: product.AveragePurchasePrice,
			Total:                total,
		})
	}
	response := &models.ResponseValueOfGoods{
		PrintDate:    time.Now().Format("02 January 2006"),
		TotalSKU:     int64(len(products)),
		TotalProduct: totalItems,
		TotalValue:   totalPrice,
		ValueOfGoods: vog,
	}

	return response, nil
}

// Sales create sales report
func (s *ReportService) Sales(rs app.RequestScope, params *util.QueryParam) (*models.ResponseSales, error) {
	var sales []models.Sales
	outgoingGoods, err := s.outgoingGoodsDAO.List(rs, params)
	if err != nil {
		return nil, err
	}
	var totalItems, totalSales, totalOmzet, netProfit int64
	for _, item := range outgoingGoods {
		product, err := s.productDAO.Get(rs, item.SKU)
		if err != nil {
			return nil, err
		}
		profit := int64(item.Total) - (int64(item.Amount) * product.AveragePurchasePrice)
		netProfit += profit
		t, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}
		// - calcaulate total omzet
		totalOmzet += int64(item.Total)
		totalItems += item.Amount
		if item.OrderID != "" {
			totalSales++
		}
		sales = append(sales, models.Sales{
			OrderID:     item.OrderID,
			Time:        t.Format("2006-01-02 03.04.05"),
			SKU:         item.SKU,
			ProductName: product.Name,
			Amount:      item.Amount,
			SellPrice:   int64(item.Price),
			Total:       int64(item.Total),
			BuyPrice:    product.AveragePurchasePrice,
			Profit:      profit,
		})
	}
	response := &models.ResponseSales{
		Sales:      sales,
		PrintDate:  time.Now().Format("02 January 2006"),
		TotalOmzet: totalOmzet,
		NetProfit:  netProfit,
		TotalSales: totalSales,
		TotalItems: totalItems,
	}

	return response, nil
}
