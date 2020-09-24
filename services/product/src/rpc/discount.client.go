package rpc

import (
	"context"
	"log"
	pb "product/proto"
	"product/src/adapters"
	"product/src/structs"
)

// DiscountServiceClientImpl implementarion of Discount client
type DiscountServiceClientImpl struct {
	ctx          context.Context
	discountGrpc pb.DiscountClient
}

// DiscountServiceClient interface
type DiscountServiceClient interface {
	Search(user structs.User, product structs.Product) (structs.Discount, error)
}

// NewDiscountServiceClient Create new client implementation
func NewDiscountServiceClient(ctx context.Context, service pb.DiscountClient) DiscountServiceClient {
	return &DiscountServiceClientImpl{ctx, service}
}

// Search function search for a discount
func (ds DiscountServiceClientImpl) Search(user structs.User, product structs.Product) (structs.Discount, error) {
	search := &pb.SearchRequest{
		Product: adapters.ProductToProto(product),
		User:    adapters.UserToProto(user),
	}

	result, err := ds.discountGrpc.Calculate(ds.ctx, search)
	if err != nil {
		log.Println("Error: ", err)
		return structs.Discount{}, err
	}

	return adapters.DiscountSearchToDomain(result), nil
}
