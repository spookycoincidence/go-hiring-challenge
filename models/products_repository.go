package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

type PaginationParams struct {
	Offset int
	Limit  int
}

type GetAllProductsParams struct {
	Pagination PaginationParams
	CategoryCode *string
	MaxPrice   *decimal.Decimal
}

type ProductListResult struct {
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
}

func (r *ProductsRepository) GetAllProducts(params GetAllProductsParams) (*ProductListResult, error) {
	if params.Pagination.Limit < 1 {
		params.Pagination.Limit = 1
	}
	if params.Pagination.Limit > 100 {
		params.Pagination.Limit = 100
	}
	if params.Pagination.Offset < 0 {
		params.Pagination.Offset = 0
	}

	query := r.db

	if params.CategoryCode != nil && *params.CategoryCode != "" {
		query = query.Joins("JOIN categories ON products.category_id = categories.id").
			Where("categories.code = ?", *params.CategoryCode)
	}

	if params.MaxPrice != nil {
		query = query.Where("price <= ?", params.MaxPrice)
	}

	var total int64
	if err := query.Model(&Product{}).Count(&total).Error; err != nil {
		return nil, err
	}

	var products []Product
	if err := query.
		Preload("Category").
		Preload("Variants").
		Offset(params.Pagination.Offset).
		Limit(params.Pagination.Limit).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return &ProductListResult{
		Products: products,
		Total:    total,
	}, nil
}

func (r *ProductsRepository) GetProductByCode(code string) (*Product, error) {
	var product Product
	if err := r.db.
		Preload("Category").
		Preload("Variants").
		Where("code = ?", code).
		First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}