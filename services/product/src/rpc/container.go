package rpc

import (
	"context"
	"google.golang.org/grpc"
	pb "product/proto"
)

// DiscountRPCContainer of discount services intances
type DiscountRPCContainer struct {
	Discount DiscountServiceClient
}

// NewDiscountRPCContainer Creates Discount clints instances
func NewDiscountRPCContainer(conn *grpc.ClientConn) DiscountRPCContainer {
	ctx := context.Background()

	return DiscountRPCContainer{
		Discount: NewDiscountServiceClient(ctx, pb.NewDiscountClient(conn)),
	}
}
