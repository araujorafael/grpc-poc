package repositories_test

import (
	"os"
	"product/src/libs/databases"
	"product/src/repositories"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProductRespository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product > Repository > ProductRepository")
}

var _ = Describe("ReadAll", func() {
	databaseURL := os.Getenv("DATABASE_URL")
	db := databases.NewPostgres(databaseURL)
	repo := repositories.NewProductRepository(db)

	Context("Success", func() {
		It("Should return a list of products", func() {
			result, err := repo.ReadAll()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(HaveLen(10))
		})
	})
})
