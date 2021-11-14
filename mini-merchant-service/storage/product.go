package storage

import (
	"mini-merchant-service/query"
	"mini-merchant-service/entity"
	"time"

	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(product entity.Products) (entity.Products, error)
	CreatedDisplayImage(displayImage entity.ImageProducts) (entity.ImageProducts, error)
	ShowAllProduct() ([]entity.Products, error)
	FindProductByID(ID string) (entity.Products, error)
	FindProductWithImageByID(ID string) (entity.Products, error)
	UpdateProductByID(ID string, input entity.UpdateProductInputs) (entity.Products, error)
	DeleteProductByID(ID string) (string, error)
	FindOutletProductByID(ID string) (entity.Outlets, error)
}

func NewProductDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) CreateProduct(product entity.Products) (entity.Products, error) {

	qry := query.QueryCreateProduct

	err := r.db.Raw(qry,
		product.ProductID,
		product.ProductName,
		product.Price,
		product.Sku,
		product.Picture,
		product.CreatedAt,
		product.UpdatedAt,
		product.OutletID).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) CreatedDisplayImage(displayImage entity.ImageProducts) (entity.ImageProducts, error) {
	qry := query.QueryCreateImage

	err := r.db.Raw(qry,
		displayImage.ImageProductID,
		displayImage.DisplayImage,
		displayImage.ProductID).Scan(&displayImage).Error

	if err != nil {
		return displayImage, err
	}

	return displayImage, nil
}

func (r *dao) ShowAllProduct() ([]entity.Products, error) {
	var product []entity.Products

	qry := query.QueryFindAllProduct

	err := r.db.Raw(qry).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) FindProductByID(ID string) (entity.Products, error) {
	var product entity.Products

	qry := query.QueryFindProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) FindProductWithImageByID(ID string) (entity.Products, error) {
	var product entity.Products

	err := r.db.Where("id = ?", ID).Preload("ImageProduct").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) UpdateProductByID(ID string, input entity.UpdateProductInputs) (entity.Products, error) {

	var product entity.Products

	input.UpdatedAt = time.Now()

	qry := query.QueryUpdateProductByID
	err := r.db.Raw(qry,
		input.ProductName,
		input.Price,
		input.Sku,
		input.Picture,
		input.OutletID,
		input.UpdatedAt,
		ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) DeleteProductByID(ID string) (string, error) {
	product := &entity.Products{}
	qry := query.QueryDeleteProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) FindOutletProductByID(ID string) (entity.Outlets, error) {
	var outlet entity.Outlets

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
