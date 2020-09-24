package src

import (
	"github.com/gin-gonic/gin"
	"product/src/handlers"
)

// CreateRoutes router
func CreateRoutes(r *gin.Engine, container handlers.Container) {
	r.GET("/products", container.Product.ReadAll)
}
