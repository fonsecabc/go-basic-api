package database

import (
	"github.com/fonsecabc/go-basic-api/internal/entities"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

type ListFilterBy struct {
	Name string
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDB) LoadById(id value_objects.ID) (*entities.Product, error) {
	var product entities.Product
	err := p.DB.First(&product, id).Error
	return &product, err
}

func (p *ProductDB) List(
	userId value_objects.ID,
	pagination value_objects.PaginationParams[ListFilterBy],
) (value_objects.PaginationResponse[entities.Product], error) {
	var total int64
	query := p.DB.Table("products").Where("user_id = ?", userId)

	if pagination.Filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+pagination.Filter.Name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return value_objects.PaginationResponse[entities.Product]{}, err
	}

	var products []entities.Product
	err = query.
		Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&products).Error

	if err != nil {
		return value_objects.PaginationResponse[entities.Product]{}, err
	}

	return value_objects.PaginationResponse[entities.Product]{
		Data:        products,
		Total:       total,
		CurrentPage: pagination.Page,
		PerPage:     pagination.PerPage,
	}, nil
}

func (p *ProductDB) Update(product *entities.Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductDB) Delete(id value_objects.ID) error {
	return p.DB.Delete(&entities.Product{}, id).Error
}
