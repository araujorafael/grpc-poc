package repositories

// ProductQueryBuilder implementation
type ProductQueryBuilder struct{}

// NewProductQueryBuilder create new instance of Product querybuilder
func NewProductQueryBuilder() *ProductQueryBuilder {
	return &ProductQueryBuilder{}
}

// ReadAll return query to find a list of Products
func (uq *ProductQueryBuilder) ReadAll() string {
	return "SELECT * FROM products"
}
