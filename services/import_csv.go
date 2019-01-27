package services

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/andrideng/inventory-system/app"
	"github.com/andrideng/inventory-system/models"
)

// ImportCsvService provides service related with imoprt csv.
type ImportCsvService struct {
	productDAO       productDAO
	incomingGoodsDAO incomingGoodsDAO
}

// NewImportCsvService creates a new Import CSV service.
func NewImportCsvService(productDAO productDAO, incomingGoodsDAO incomingGoodsDAO) *ImportCsvService {
	return &ImportCsvService{productDAO, incomingGoodsDAO}
}

// ImportProducts import ./storages/products.csv
func (s *ImportCsvService) ImportProducts(rs app.RequestScope) ([]models.Product, error) {
	csvFile, err := os.Open("./storages/products.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	idx := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logrus.Error(err)
		}
		if idx != 0 {
			re := regexp.MustCompile(`\((.*?)\)`)
			sizeAndColor := re.FindStringSubmatch(line[1])[1]
			split := strings.Split(sizeAndColor, ",")
			name := re.ReplaceAllString(line[1], "")
			product := &models.Product{
				SKU:   line[0],
				Name:  name,
				Size:  split[0],
				Color: split[1],
			}
			if err := s.productDAO.Create(rs, product); err != nil {
				return nil, err
			}
		}
		idx++
	}
	return s.productDAO.List(rs)
}

// ImportIncomingGoods import ./storages/incoming-goods.csv
func (s *ImportCsvService) ImportIncomingGoods(rs app.RequestScope) ([]models.IncomingGoods, error) {
	csvFile, err := os.Open("./storages/incoming-goods.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	idx := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logrus.Error(err)
		}
		if idx != 0 {
			amountOrders, err := strconv.Atoi(line[3])
			if err != nil {
				return nil, err
			}
			amountReceived, err := strconv.Atoi(line[4])
			if err != nil {
				return nil, err
			}
			purchasePrice, err := strconv.Atoi(line[5])
			if err != nil {
				return nil, err
			}
			incomingGoods := &models.IncomingGoods{
				SKU:            line[1],
				AmountOrders:   int64(amountOrders),
				AmountReceived: int64(amountReceived),
				PurchasePrice:  float64(purchasePrice),
				ReceiptNumber:  line[7],
			}
			if err := s.incomingGoodsDAO.Create(rs, incomingGoods); err != nil {
				return nil, err
			}
		}
		idx++
	}
	return s.incomingGoodsDAO.List(rs)
}
