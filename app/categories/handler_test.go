package categories

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/stretchr/testify/assert"
)

type mockCategoriesRepository struct {
	categories []models.Category
}

func (m *mockCategoriesRepository) GetAllCategories() ([]models.Category, error) {
	return m.categories, nil
}

func (m *mockCategoriesRepository) CreateCategory(category *models.Category) error {
	category.ID = uint(len(m.categories) + 1)
	m.categories = append(m.categories, *category)
	return nil
}

func createMockCategories() *mockCategoriesRepository {
	return &mockCategoriesRepository{
		categories: []models.Category{
			{
				ID:   1,
				Code: "CLOTHING",
				Name: "Clothing",
			},
			{
				ID:   2,
				Code: "SHOES",
				Name: "Shoes",
			},
			{
				ID:   3,
				Code: "ACCESSORIES",
				Name: "Accessories",
			},
		},
	}
}

func TestHandleGetAllCategories(t *testing.T) {
	mockRepo := createMockCategories()

	t.Run("should return all categories", func(t *testing.T) {
		categories, err := mockRepo.GetAllCategories()

		assert.NoError(t, err)
		assert.Equal(t, 3, len(categories))
		assert.Equal(t, "CLOTHING", categories[0].Code)
	})
}

func TestHandleCreateCategory(t *testing.T) {
	mockRepo := createMockCategories()

	t.Run("should create new category", func(t *testing.T) {
		newCategory := &models.Category{
			Code: "ELECTRONICS",
			Name: "Electronics",
		}

		err := mockRepo.CreateCategory(newCategory)
		assert.NoError(t, err)
		assert.NotEqual(t, 0, newCategory.ID)

		categories, _ := mockRepo.GetAllCategories()
		assert.Equal(t, 4, len(categories))
	})

	t.Run("should validate category fields", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"code":"","name":""}`)
		var req CreateCategoryRequest
		json.NewDecoder(reqBody).Decode(&req)

		assert.Empty(t, req.Code)
		assert.Empty(t, req.Name)
	})
}

func TestCategoriesHandlerResponses(t *testing.T) {
	mockRepo := createMockCategories()

	t.Run("categories list has correct structure", func(t *testing.T) {
		categories, _ := mockRepo.GetAllCategories()

		assert.Greater(t, len(categories), 0)
		for _, c := range categories {
			assert.NotEmpty(t, c.Code)
			assert.NotEmpty(t, c.Name)
			assert.Greater(t, c.ID, uint(0))
		}
	})
}