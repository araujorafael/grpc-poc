package handlers

import (
	"log"
	"net/http"
	"product/src/repositories"
	"product/src/rpc"
	"product/src/structs"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// ProductHandler handler implementation
type ProductHandler struct {
	discountService   rpc.DiscountServiceClient
	productRepository *repositories.ProductRepositoryImpl
	userRepository    *repositories.UserRepositoryImpl
}

// NewProductHandler create a implementation
func NewProductHandler(discountService rpc.DiscountRPCContainer, repos repositories.Container) *ProductHandler {
	return &ProductHandler{
		discountService:   discountService.Discount,
		productRepository: repos.Product,
		userRepository:    repos.User,
	}
}

// ReadAll return a list of products
func (ph ProductHandler) ReadAll(c *gin.Context) {
	products, _ := ph.productRepository.ReadAll()

	userID := c.Request.Header.Get("X-USER-ID")
	if len(userID) > 0 {
		user, err := ph.userRepository.Find(userID)
		if err == nil {
			products = ph.getDiscounts(user, products)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func (ph ProductHandler) getDiscounts(user structs.User, products []structs.Product) []structs.Product {
	for i := range products {
		la, err := ph.discountService.Search(user, products[i])
		if err != nil {
			grpc, _ := status.FromError(err)
			log.Println("[INFO] ", grpc.Message())
			break
		}

		products[i].Discount = la
	}

	return products
}
