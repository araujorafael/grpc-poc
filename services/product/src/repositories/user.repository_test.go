package repositories_test

import (
	"os"
	"product/src/libs/databases"
	"product/src/repositories"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product > Repository > UserRepository")
}

var _ = Describe("Find", func() {
	databaseURL := os.Getenv("DATABASE_URL")
	db := databases.NewPostgres(databaseURL)
	repo := repositories.NewUserRepository(db)

	Context("Success", func() {
		It("Should return user", func() {
			result, err := repo.Find("8822888d-2713-4911-851b-c94b8aa60490")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.ID).Should(Equal("8822888d-2713-4911-851b-c94b8aa60490"))
		})
	})

	Context("Error", func() {
		It("Should return error when nothing is found", func() {
			_, err := repo.Find("33ea12ca-3970-49e0-8fa3-5cce468a1829")
			Expect(err).Should(HaveOccurred())
		})
	})
})
