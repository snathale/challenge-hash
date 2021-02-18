package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
	"github.com/snathale/challenge-hash/calculator/test_helper"
)

func TestIntegration(t *testing.T) {
	g := NewGomegaWithT(t)
	db := test_helper.GetDatabase(g, "dummy_discount_db")
	productColl := test_helper.GetCollection(g, "product", db)
	err := productColl.Truncate(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	userColl := test_helper.GetCollection(g, "user", db)
	err = userColl.Truncate(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	product := entity.Product{
		PriceInCents: 6999,
		Title:        "macbook pro",
		Description:  "notebook apple",
	}
	productMeta := test_helper.CreateDocument(g, productColl, product)
	user := entity.User{
		FirstName: "Joe",
		LastName:  "Lee",
		Birthday:  time.Now(),
	}
	userMeta := test_helper.CreateDocument(g, userColl, user)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3333/product", nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	req.Header.Set("X-USER-ID", userMeta.Key)
	t.Run("validate receive a product with birthday discount on node client", func(t *testing.T) {
		resp, err := client.Do(req)
		g.Expect(err).ShouldNot(HaveOccurred())
		body, err := ioutil.ReadAll(resp.Body)
		var products Response
		err = json.Unmarshal(body, &products)
		fmt.Println(products)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(products).Should(Equal(Response{
			Status: "success",
			Data: []FakeProductDiscount{
				{
					Id:           productMeta.Key,
					PriceInCents: 6999,
					Title:        "macbook pro",
					Description:  "notebook apple",
					Discount: FakeDiscount{
						Percentage:   float32(0.05000000074505806),
						ValueInCents: 6650,
					},
				},
			},
		}))

	})
}

type Response struct {
	Status string                `json:"status"`
	Data   []FakeProductDiscount `json:"data"`
}

type FakeDiscount struct {
	Percentage   float32 `json:"percentage"`
	ValueInCents int     `json:"value_in_cents"`
}

type FakeProductDiscount struct {
	Id           string `json:"id"`
	PriceInCents int    `json:"price_in_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Discount     FakeDiscount
}
