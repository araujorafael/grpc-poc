package repositories

// Container aggragetes repositories
type Container struct {
	Product *ProductRepositoryImpl
	User    *UserRepositoryImpl
}
