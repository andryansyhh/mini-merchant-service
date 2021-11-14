package formatter

import (
	"mini-merchant-service/entity"
	"time"
)

type ProductFormat struct {
	ID          string    `json:"id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	Sku         string    `json:"sku"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OutletID    string    `json:"outlet_id"`
}

func FormatProduct(product entity.Products) ProductFormat {
	var formatProduct = ProductFormat{
		ID:          product.ProductID,
		ProductName: product.ProductName,
		Price:       product.Price,
		Sku:         product.Sku,
		Picture:     product.Picture,
		OutletID:    product.OutletID,
	}

	return formatProduct
}

type ProductDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDeleteProduct(msg string) ProductDeleteFormat {
	var deleteFormat = ProductDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
