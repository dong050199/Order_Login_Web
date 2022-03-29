package repository

import (
	models "SQLite_JWT/model/product"
)

type ProductRepo interface {
	SelectId(id int) (models.Product, error)
	Select() ([]models.Product, error)
	Insert(u models.Product) error
	Update(u models.Product, id int) error
	Delete(u int) error
}
