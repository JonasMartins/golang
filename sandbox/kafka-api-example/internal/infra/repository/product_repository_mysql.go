package repository

import (
	"database/sql"
	"project/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(p *entity.Product) error {
	_, err := r.DB.Exec(
		"insert into products (id, name, price) values (?,?.?)",
		p.ID,
		p.Name,
		p.Price,
	)

	if err != nil {
		return err
	}

	return nil

}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil

}
