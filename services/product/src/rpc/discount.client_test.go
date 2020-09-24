package rpc_test

import (
	"product/mocks"
	pb "product/proto"
	"product/src/rpc"
	"product/src/structs"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product > RPC > DiscountServiceClient")
}

var _ = Describe("Search", func() {
	var (
		stub    *mocks.DiscountClient
		ctx     context.Context
		product structs.Product
		user    structs.User
		resp    structs.Discount
		err     error
	)

	BeforeEach(func() {
		product = structs.Product{
			ID:    1001,
			Name:  "Maquininha de cartão",
			Sku:   "a12f8bc6-b17e-4cae-a98f-85aaaaaaaaa",
			Price: 1000,
		}

		user = structs.User{
			ID:       "a12f8bc6-b17e-4cae-a98f-85aac3b4d821",
			Name:     "Roberta Sanches",
			Birthday: time.Date(1989, time.November, 10, 23, 0, 0, 0, time.UTC),
		}
	})

	JustBeforeEach(func() {
		discountServiceClient := rpc.NewDiscountServiceClient(ctx, stub)
		resp, err = discountServiceClient.Search(user, product)
	})

	Context("Success", func() {
		BeforeEach(func() {
			discount := &pb.SearchResponse{
				DiscountRate: 10,
				NewPrice:     900,
			}
			stub = new(mocks.DiscountClient)
			stub.On("Calculate", mock.Anything, mock.Anything).Return(discount, nil)
		})

		It("Should return parsed product with discount", func() {
			Expect(resp.DiscountPrice).Should(Equal(int32(900)))

			expectedRequest := &pb.SearchRequest{
				Product: &pb.Product{
					ProductName: product.Name,
					Price:       product.Price,
				},
				User: &pb.User{
					DateOfBirth: uint64(user.Birthday.Unix()),
					UserId:      user.ID,
				},
			}
			stub.AssertCalled(GinkgoT(), "Calculate", ctx, expectedRequest)
		})
	})

	Context("Error", func() {
		BeforeEach(func() {
			err := status.Error(codes.NotFound, "Discount Not found")
			stub = new(mocks.DiscountClient)
			stub.On("Calculate", mock.Anything, mock.Anything).Return(&pb.SearchResponse{}, err)
		})

		It("Should propagate error returned by gRPC", func() {
			Expect(err).Should(HaveOccurred())

			st, _ := status.FromError(err)
			Expect(st.Code()).Should(Equal(codes.NotFound))
		})
	})
})
