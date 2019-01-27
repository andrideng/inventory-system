package apis

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
	"github.com/andrideng/inventory-system/util"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// reportService specifies the interface for the report service method needed by reportReource.
	reportService interface {
		// ValueOfGoods
		ValueOfGoods(rs app.RequestScope) (*models.ResponseValueOfGoods, error)
		// Sales
		Sales(rs app.RequestScope, params *util.QueryParam) (*models.ResponseSales, error)
	}

	// reportResource defines the handlers.
	reportResource struct {
		service reportService
	}
)

// ServerReportResource sets up the routing of report endpoints and corresponding handlers.
func ServerReportResource(rg *routing.RouteGroup, service reportService) {
	r := &reportResource{service}
	rg.Get("/reports/value-of-goods", r.valueOfGoods)
	rg.Get("/reports/sales", r.sales)
}

func (r *reportResource) valueOfGoods(c *routing.Context) error {
	response, err := r.service.ValueOfGoods(app.GetRequestScope(c))
	if err != nil {
		return err
	}
	// - generate to csv
	f, err := os.Create("./reports/" +
		time.Now().Format("02_Jan_2006_03_04_05") +
		"-goods-of-values.csv")
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	var record [][]string
	record = append(record, []string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})
	for i, obj := range response.ValueOfGoods {
		record = append(record, []string{
			obj.SKU,
			obj.Name,
			strconv.Itoa(int(obj.Amount)),
			strconv.Itoa(int(obj.AveragePurchasePrice)),
			strconv.Itoa(int(obj.Total)),
		})
		w.Write(record[i])
	}
	w.Flush()

	return c.Write(Response{Data: response})
}

func (r *reportResource) sales(c *routing.Context) error {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if startDate == "" || endDate == "" {
		return errors.New("Must fill start_date and end_date")
	}
	params := &util.QueryParam{
		StartDate: startDate,
		EndDate:   endDate,
	}
	response, err := r.service.Sales(app.GetRequestScope(c), params)
	if err != nil {
		return err
	}
	// - generate to csv
	f, err := os.Create("./reports/" +
		time.Now().Format("02_Jan_2006_03_04_05") +
		"-sales.csv")
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	var record [][]string
	record = append(record, []string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual",
		"Total", "Harga Beli", "Laba"})
	for i, obj := range response.Sales {
		record = append(record, []string{
			obj.OrderID,
			obj.Time,
			obj.SKU,
			obj.ProductName,
			strconv.Itoa(int(obj.Amount)),
			strconv.Itoa(int(obj.SellPrice)),
			strconv.Itoa(int(obj.Total)),
			strconv.Itoa(int(obj.BuyPrice)),
			strconv.Itoa(int(obj.Profit)),
		})
		w.Write(record[i])
	}
	w.Flush()
	return c.Write(Response{Data: response})
}
