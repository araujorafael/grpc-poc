package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"os"
	"product/src"
	"product/src/handlers"
	"product/src/libs/databases"
	"product/src/repositories"
	"product/src/rpc"
)

func getRpcConn(url string) *grpc.ClientConn {
	client := rpc.NewClient(url)
	return client.Connect()
}

func main() {
	r := gin.Default()

	discountConn := getRpcConn(os.Getenv("DISCOUNT_ADDR"))
	discountClient := rpc.NewDiscountRPCContainer(discountConn)

	databaseURL := os.Getenv("DATABASE_URL")
	db := databases.NewPostgres(databaseURL)

	repos := repositories.Container{
		Product: repositories.NewProductRepository(db),
		User:    repositories.NewUserRepository(db),
	}

	handlerContainer := handlers.Container{
		Product: handlers.NewProductHandler(discountClient, repos),
	}

	src.CreateRoutes(r, handlerContainer)

	r.Run()
}
