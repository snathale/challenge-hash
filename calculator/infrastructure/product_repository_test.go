package infrastructure

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
	"github.com/snathale/challenge-hash/calculator/test_helper"
)

func TestProductRepository(t *testing.T) {
	g := NewGomegaWithT(t)
	defer test_helper.DeleteDatabase(g, "dummy_test")
	db := test_helper.CreateDatabase(g, "dummy_test")
	t.Run("validate retrieve a existing product", func(t *testing.T) {
		coll := test_helper.GetCollection(g, "product", db)
		doc := entity.Product{
			PriceInCents: 3000,
			Title:        "iphone 11",
			Description:  "smartphone apple",
		}
		meta := test_helper.CreateDocument(g, coll, doc)
		rep := NewProductRepository(coll)
		product, err := rep.GetProductById(meta.Key)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(product.PriceInCents).Should(BeEquivalentTo(3000))
		g.Expect(product.Title).Should(BeEquivalentTo("iphone 11"))
		g.Expect(product.Description).Should(BeEquivalentTo("smartphone apple"))
		g.Expect(product.Id).ShouldNot(BeEmpty())
	})
	t.Run("validade receive error when product not exists", func(t *testing.T) {
		coll := test_helper.GetCollection(g, "product", db)
		rep := NewProductRepository(coll)
		product, err := rep.GetProductById("14")
		g.Expect(product).Should(BeNil())
		g.Expect(err).Should(MatchError(GetProductDocumentError))
	})
}
