package adapters

import (
	"fmt"
	pb "product/proto"
	"product/src/structs"
)

// ProductToProto converts domain product to proto
func ProductToProto(product structs.Product) *pb.Product {
	return &pb.Product{
		ProductName: product.Name,
		Price:       product.Price,
	}
}

// UserToProto converts domain product to proto
func UserToProto(user structs.User) *pb.User {
	return &pb.User{
		DateOfBirth: uint64(user.Birthday.Unix()),
		UserId:      user.ID,
	}
}

// DiscountSearchToDomain converts proto discount to domain stuct
func DiscountSearchToDomain(discount *pb.SearchResponse) structs.Discount {
	return structs.Discount{
		DiscountRate:  fmt.Sprintf("%d%%", discount.GetDiscountRate()),
		DiscountPrice: discount.GetNewPrice(),
	}
}
