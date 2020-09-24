package repositories

import (
	"log"
	"product/src/libs/databases"
	"product/src/structs"
)

// UserRepositoryImpl implementation
type UserRepositoryImpl struct {
	querybuilder UserQueryBuilder
	db           *databases.Postgres
}

// NewUserRepository return implementation of user repository
func NewUserRepository(db *databases.Postgres) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// Find search for user using user ID
func (u *UserRepositoryImpl) Find(userID string) (structs.User, error) {
	query := u.querybuilder.Find()
	var found structs.User

	err := u.db.GetConnection().Get(&found, query, userID)
	if err != nil {
		log.Println("[ERROR] ", "User not found")
	}

	return found, err
}
