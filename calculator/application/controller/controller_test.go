package controller

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
	"github.com/snathale/challenge-hash/calculator/infrastucture"
	"github.com/snathale/challenge-hash/calculator/test_helper"
)

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	defer test_helper.DeleteDatabase(g, "dummy_test_controller")
	db := test_helper.CreateDatabase(g, "dummy_test_controller")
	productColl := test_helper.GetCollection(g, "product", db)
	userColl := test_helper.GetCollection(g, "user", db)
	doc := entity.Product{
		PriceInCents: 3000,
		Title:        "iphone 11",
		Description:  "smartphone apple",
	}
	productMeta := test_helper.CreateDocument(g, productColl, doc)
	rep, err := infrastucture.NewRepositories(mockDBConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	ctrl := NewController(rep)
	t.Run("validate receive a discount 0.1 when is blackFriday", func(t *testing.T) {
		doc := entity.User{
			FirstName: "Joe",
			LastName:  "Lee",
			Birthday:  time.Now(),
		}
		userMeta := test_helper.CreateDocument(g, userColl, doc)
		date := time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC)
		test_helper.ReplaceDate(date)
		discount, err := ctrl.CalculateDiscount(userMeta.Key, productMeta.Key)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0.1)))
		g.Expect(discount.ValueInCents).To(Equal(2700))
	})
	t.Run("validate receive a discount 0.5 when is user bithday", func(t *testing.T) {
		date := time.Date(2021, 01, 25, 0, 0, 0, 0, time.UTC)
		doc := entity.User{
			FirstName: "Joe",
			LastName:  "Lee",
			Birthday:  date,
		}
		userMeta := test_helper.CreateDocument(g, userColl, doc)
		test_helper.ReplaceDate(date)
		discount, err := ctrl.CalculateDiscount(userMeta.Key, productMeta.Key)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0.05)))
		g.Expect(discount.ValueInCents).To(Equal(2850))
	})
	t.Run("validate receive a product without discount", func(t *testing.T) {
		doc := entity.User{
			FirstName: "Joe",
			LastName:  "Lee",
			Birthday:  time.Now(),
		}
		date := time.Date(2021, 01, 25, 0, 0, 0, 0, time.UTC)
		userMeta := test_helper.CreateDocument(g, userColl, doc)
		test_helper.ReplaceDate(date)
		discount, err := ctrl.CalculateDiscount(userMeta.Key, productMeta.Key)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0)))
		g.Expect(discount.ValueInCents).To(Equal(3000))
	})
}

func mockDBConfig() infrastucture.Config {
	return infrastucture.Config{
		Port:                  8529,
		Password:              "dummy_passowrd",
		Database:              "dummy_test_controller",
		Host:                  "http://localhost",
		User:                  "root",
		UserCollectionName:    "user",
		ProductCollectionName: "product",
	}
}
