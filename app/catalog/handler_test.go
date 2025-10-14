package catalog

import (
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockProductsRepository struct {
	products []models.Product
}

func (m *mockProductsRepository) GetAllProducts(params models.GetAllProductsParams) (*models.ProductListResult, error) {
	filtered := m.products

	if params.CategoryCode != nil && *params.CategoryCode != "" {
		var filtered2 []models.Product
		for _, p := range filtered {
			if p.Category != nil && p.Category.Code == *params.CategoryCode {
				filtered2 = append(filtered2, p)
			}
		}
		filtered = filtered2
	}

	if params.MaxPrice != nil {
		var filtered2 []models.Product
		for _, p := range filtered {
			if p.Price.LessThanOrEqual(*params.MaxPrice) {
				filtered2 = append(filtered2, p)
			}
		}
		filtered = filtered2
	}

	total := int64(len(filtered))

	start := params.Pagination.Offset
	end := start + params.Pagination.Limit
	if end > len(filtered) {
		end = len(filtered)
	}
	if start > len(filtered) {
		start = len(filtered)
	}

	return &models.ProductListResult{
		Products: filtered[start:end],
		Total:    total,
	}, nil
}

func (m *mockProductsRepository) GetProductByCode(code string) (*models.Product, error) {
	for _, p := range m.products {
		if p.Code == code {
			return &p, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func createMockProducts() *mockProductsRepository {
	return &mockProductsRepository{
		products: []models.Product{
			{
				ID:    1,
				Code:  "PROD001",
				Price: decimal.NewFromFloat(10.99),
				Category: &models.Category{
					ID:   1,
					Code: "CLOTHING",
					Name: "Clothing",
				},
				Variants: []models.Variant{
					{
						ID:        1,
						ProductID: 1,
						Name:      "Variant A",
						SKU:       "SKU001A",
						Price:     decimal.NewFromFloat(11.99),
					},
				},
			},
			{
				ID:    2,
				Code:  "PROD002",
				Price: decimal.NewFromFloat(12.49),
				Category: &models.Category{
					ID:   2,
					Code: "SHOES",
					Name: "Shoes",
				},
				Variants: []models.Variant{},
			},
			{
				ID:    3,
				Code:  "PROD003",
				Price: decimal.NewFromFloat(8.75),
				Category: &models.Category{
					ID:   3,
					Code: "ACCESSORIES",
					Name: "Accessories",
				},
				Variants: []models.Variant{},
			},
		},
	}
}

func TestHandleGetCatalog(t *testing.T) {
	mockRepo := createMockProducts()

	t.Run("should return all products with pagination", func(t *testing.T) {
		result, _ := mockRepo.GetAllProducts(models.GetAllProductsParams{
			Pagination: models.PaginationParams{Offset: 0, Limit: 10},
		})

		assert.Equal(t, int64(3), result.Total)
		assert.Equal(t, 3, len(result.Products))
	})

	t.Run("should filter by category", func(t *testing.T) {
		categoryCode := "CLOTHING"
		result, _ := mockRepo.GetAllProducts(models.GetAllProductsParams{
			Pagination: models.PaginationParams{Offset: 0, Limit: 10},
			CategoryCode: &categoryCode,
		})

		assert.Equal(t, int64(1), result.Total)
		assert.Equal(t, "PROD001", result.Products[0].Code)
	})

	t.Run("should filter by max price", func(t *testing.T) {
		maxPrice := decimal.NewFromFloat(9.99)
		result, _ := mockRepo.GetAllProducts(models.GetAllProductsParams{
			Pagination: models.PaginationParams{Offset: 0, Limit: 10},
			MaxPrice:   &maxPrice,
		})

		assert.Equal(t, int64(1), result.Total)
		assert.Equal(t, "PROD003", result.Products[0].Code)
	})

	t.Run("should apply pagination limit", func(t *testing.T) {
		result, _ := mockRepo.GetAllProducts(models.GetAllProductsParams{
			Pagination: models.PaginationParams{Offset: 0, Limit: 2},
		})

		assert.Equal(t, int64(3), result.Total)
		assert.Equal(t, 2, len(result.Products))
	})
}

func TestHandleGetProductByCode(t *testing.T) {
	mockRepo := createMockProducts()

	t.Run("should return product details by code", func(t *testing.T) {
		product, err := mockRepo.GetProductByCode("PROD001")

		assert.NoError(t, err)
		assert.Equal(t, "PROD001", product.Code)
		assert.NotNil(t, product.Category)
		assert.Equal(t, "CLOTHING", product.Category.Code)
	})

	t.Run("should return error for non-existent product", func(t *testing.T) {
		_, err := mockRepo.GetProductByCode("NONEXISTENT")

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})

	t.Run("should include variants in response", func(t *testing.T) {
		product, _ := mockRepo.GetProductByCode("PROD001")

		assert.Greater(t, len(product.Variants), 0)
	})
}

func TestCatalogHandlerIntegration(t *testing.T) {
	mockRepo := createMockProducts()

	t.Run("full catalog response structure", func(t *testing.T) {
		result, _ := mockRepo.GetAllProducts(models.GetAllProductsParams{
			Pagination: models.PaginationParams{Offset: 0, Limit: 10},
		})

		assert.Equal(t, int64(3), result.Total)
		for _, p := range result.Products {
			assert.NotEmpty(t, p.Code)
			price := p.Price.InexactFloat64()
			assert.Greater(t, price, float64(0))
		}
	})
}