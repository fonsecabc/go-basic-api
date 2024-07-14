package database

import (
	"testing"

	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductDB_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	user, _ := entities.NewUser("test@gmail.com", "password")

	product, _ := entities.NewProduct("test", 10.0, user.ID)
	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	loadedProduct, err := productDB.LoadById(product.ID)

	assert.Nil(t, err)
	assert.Equal(t, product.ID, loadedProduct.ID)
	assert.Equal(t, product.Name, loadedProduct.Name)
	assert.Equal(t, product.Price, loadedProduct.Price)
	assert.Equal(t, product.UserID, loadedProduct.UserID)
}

func TestProductDB_List(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	user, _ := entities.NewUser("test@gmail.com", "password")

	tests := []struct {
		name   string
		price  float32
		userID value_objects.ID
	}{
		{"test", 10.2, user.ID},
		{"test2", 20.2, user.ID},
		{"test3", 30.2, user.ID},
	}

	productDB := NewProductDB(db)
	var products []entities.Product

	for _, testP := range tests {
		p, _ := entities.NewProduct(testP.name, testP.price, testP.userID)

		err = productDB.Create(p)
		assert.Nil(t, err)

		products = append(products, *p)
	}

	pagination := value_objects.PaginationParams[ListFilterBy]{PerPage: 10, Page: 1, Filter: ListFilterBy{}}
	listedProducts, err := productDB.List(user.ID, pagination)
	assert.Nil(t, err)

	assert.Equal(t, listedProducts.CurrentPage, 1)
	assert.Equal(t, listedProducts.PerPage, 10)
	assert.Equal(t, listedProducts.Total, int64(len(products)))
	assert.Equal(t, listedProducts.Total, int64(len(listedProducts.Data)))

	for i, p := range listedProducts.Data {
		assert.Equal(t, products[i].ID, p.ID)
		assert.Equal(t, products[i].Name, p.Name)
		assert.Equal(t, products[i].Price, p.Price)
		assert.Equal(t, products[i].UserID, p.UserID)
	}

}

func TestProductDB_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	user, _ := entities.NewUser("test@gmail.com", "password")

	product, _ := entities.NewProduct("test", 10.0, user.ID)
	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	assert.Equal(t, "test", product.Name)
	assert.Equal(t, float32(10.0), product.Price)
	assert.Equal(t, user.ID, product.UserID)

	product.Name = "new name"
	product.Price = 20.0

	err = productDB.Update(product)
	assert.Nil(t, err)

	loadedProduct, err := productDB.LoadById(product.ID)
	assert.Nil(t, err)

	assert.Equal(t, "new name", loadedProduct.Name)
	assert.Equal(t, float32(20.0), loadedProduct.Price)
}

func TestProductDB_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	user, _ := entities.NewUser("test@gmail.com", "password")

	product, _ := entities.NewProduct("test", 10.0, user.ID)
	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	assert.Equal(t, "test", product.Name)
	assert.Equal(t, float32(10.0), product.Price)
	assert.Equal(t, user.ID, product.UserID)

	err = productDB.Delete(product.ID)
	assert.Nil(t, err)

	_, err = productDB.LoadById(product.ID)
	assert.NotNil(t, err)
}
