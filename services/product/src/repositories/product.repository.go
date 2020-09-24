package repositories

import (
	"log"
	"product/src/libs/databases"
	"product/src/structs"
)

// ProductRepositoryImpl implementation
type ProductRepositoryImpl struct {
	querybuilder ProductQueryBuilder
	db           *databases.Postgres
}

// NewProductRepository return implementation of Product repository
func NewProductRepository(db *databases.Postgres) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

// ReadAll return a list od products
func (u *ProductRepositoryImpl) ReadAll() ([]structs.Product, error) {
	query := u.querybuilder.ReadAll()
	var found []structs.Product

	err := u.db.GetConnection().Select(&found, query)
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	return found, err
}
