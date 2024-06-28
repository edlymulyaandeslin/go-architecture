package repository

import (
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"database/sql"
	"math"
)

// buat interface
type ProductRepository interface {
	GetAll(page int, size int) ([]model.Product, dto.Paginate, error)
	GetById(id string) (model.Product, error)
}

// struct
type productRepository struct {
	db *sql.DB
}

func (p *productRepository) GetAll(page int, size int) ([]model.Product, dto.Paginate, error) {
	var listData []model.Product

	// rumus pagination
	offset := (page - 1) * size

	rows, err := p.db.Query("SELECT id, name, price, type, created_at, updated_at FROM products LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		return nil, dto.Paginate{}, err
	}

	totalRows := 0
	err = p.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&totalRows)
	if err != nil {
		return nil, dto.Paginate{}, err
	}

	for rows.Next() {
		var product model.Product

		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Type, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, dto.Paginate{}, err
		}

		listData = append(listData, product)
	}

	paginate := dto.Paginate{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return listData, paginate, nil

}

func (p *productRepository) GetById(id string) (model.Product, error) {
	var product model.Product

	err := p.db.QueryRow("SELECT id, name, price, type, created_at, updated_at FROM products WHERE id =$1", id).Scan(&product.Id, &product.Name, &product.Price, &product.Type, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func NewProductRepository(database *sql.DB) ProductRepository {
	return &productRepository{db: database}
}

// constructor
