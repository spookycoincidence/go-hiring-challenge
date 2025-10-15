package catalog

import (
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	Code     string                   `json:"code"`
	Name     string                   `json:"name"`
	Price    float64                  `json:"price"`
	Category *CategoryResponse        `json:"category,omitempty"`
	Variants []ProductVariantResponse `json:"variants,omitempty"`
}

type CategoryResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductVariantResponse struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
}

type CatalogResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
}

type CatalogHandler struct {
	repo *models.ProductsRepository
}

func NewCatalogHandler(r *models.ProductsRepository) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	offset := 0
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if val, err := strconv.Atoi(offsetStr); err == nil {
			offset = val
		}
	}

	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if val, err := strconv.Atoi(limitStr); err == nil {
			limit = val
		}
	}

	var categoryCode *string
	if cat := r.URL.Query().Get("category"); cat != "" {
		categoryCode = &cat
	}

	var maxPrice *decimal.Decimal
	if priceStr := r.URL.Query().Get("maxPrice"); priceStr != "" {
		if price, err := decimal.NewFromString(priceStr); err == nil {
			maxPrice = &price
		}
	}

	params := models.GetAllProductsParams{
		Pagination: models.PaginationParams{
			Offset: offset,
			Limit:  limit,
		},
		CategoryCode: categoryCode,
		MaxPrice:     maxPrice,
	}

	result, err := h.repo.GetAllProducts(params)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	products := make([]ProductResponse, len(result.Products))
	for i, p := range result.Products {
		products[i] = mapProductToResponse(p)
	}

	response := CatalogResponse{
		Products: products,
		Total:    result.Total,
	}

	api.OKResponse(w, response)
}

func (h *CatalogHandler) HandleGetByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	product, err := h.repo.GetProductByCode(code)
	if err != nil {
		api.ErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	response := mapProductToResponse(*product)
	api.OKResponse(w, response)
}

func mapProductToResponse(p models.Product) ProductResponse {
	resp := ProductResponse{
		Code:  p.Code,
		Name:  p.Name,
		Price: p.Price.InexactFloat64(),
	}

	if p.Category != nil {
		resp.Category = &CategoryResponse{
			Code: p.Category.Code,
			Name: p.Category.Name,
		}
	}

	variants := make([]ProductVariantResponse, len(p.Variants))
	for i, v := range p.Variants {
		price := v.Price
		if price.IsZero() {
			price = p.Price
		}

		variants[i] = ProductVariantResponse{
			Name:  v.Name,
			SKU:   v.SKU,
			Price: price.InexactFloat64(),
		}
	}
	resp.Variants = variants

	return resp
}
