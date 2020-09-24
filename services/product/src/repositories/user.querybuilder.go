package repositories

// UserQueryBuilder implementation
type UserQueryBuilder struct{}

// NewUserQueryBuilder create new instance of User querybuilder
func NewUserQueryBuilder() *UserQueryBuilder {
	return &UserQueryBuilder{}
}

// Find return query to find one user
func (uq *UserQueryBuilder) Find() string {
	return "SELECT * FROM users WHERE id = $1"
}
