package product

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	productDomain "myclothing/entities/product"
	"testing"
)

func TestProductPostgresRepository_SelectProductCategoryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "Product 1", "Description 1")

	mock.ExpectQuery(`(?i)--\s*name:\s*GetProductCategoryById\s*:one\s*select\s*id,\s*name,\s*description\s*from\s*product_category\s*where\s*id\s*=\s*\$1`).
		WillReturnRows(rows)

	repo := NewProductPostgresRepository(db)

	category, err := repo.SelectProductCategoryById(1)
	if err != nil {
		t.Errorf("error was not expected when calling SelectProducts: %s", err)
	}

	expectedCategory := productDomain.Category{
		Id:          1,
		Name:        "Product 1",
		Description: "Description 1",
	}

	assert.Equal(t, &expectedCategory, category)
}
