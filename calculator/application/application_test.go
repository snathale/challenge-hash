package application

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
	"github.com/snathale/challenge-hash/calculator/infrastructure"
	"github.com/snathale/challenge-hash/calculator/interface/proto"
	"github.com/snathale/challenge-hash/calculator/interface/server"
	"github.com/snathale/challenge-hash/calculator/test_helper"
	"google.golang.org/grpc"
)

func TestApplication(t *testing.T) {
	g := NewGomegaWithT(t)
	defer test_helper.DeleteDatabase(g, "dummy_test_app")
	db := test_helper.CreateDatabase(g, "dummy_test_app")
	productColl := test_helper.GetCollection(g, "product", db)
	userColl := test_helper.GetCollection(g, "user", db)
	doc := entity.Product{
		PriceInCents: 6999,
		Title:        "macbook pro",
		Description:  "notebook apple",
	}
	productMeta := test_helper.CreateDocument(g, productColl, doc)
	config := mockConfig()
	app, err := NewApp(config)
	app.Run()
	defer app.Close()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port), grpc.WithInsecure(), grpc.WithBlock())
	g.Expect(err).ShouldNot(HaveOccurred())
	client := proto.NewCalculatorClient(conn)
	defer conn.Close()
	t.Run("validate receive a discount 0.1 when is blackFriday", func(t *testing.T) {
		doc := entity.User{
			FirstName: "Joe",
			LastName:  "Lee",
			Birthday:  time.Now(),
		}
		userMeta := test_helper.CreateDocument(g, userColl, doc)
		date := time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC)
		test_helper.ReplaceDate(date)
		discount, err := client.GetProductDiscount(context.Background(), &proto.Request{
			UserId:    userMeta.Key,
			ProductId: productMeta.Key,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0.1)))
		g.Expect(discount.ValueInCents).To(Equal(int32(6300)))
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
		discount, err := client.GetProductDiscount(context.Background(), &proto.Request{
			UserId:    userMeta.Key,
			ProductId: productMeta.Key,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0.05)))
		g.Expect(discount.ValueInCents).To(Equal(int32(6650)))
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
		discount, err := client.GetProductDiscount(context.Background(), &proto.Request{
			UserId:    userMeta.Key,
			ProductId: productMeta.Key,
		})
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(discount.Percentage).To(Equal(float32(0)))
		g.Expect(discount.ValueInCents).To(Equal(int32(0)))
	})
}

func mockConfig() *Config {
	return &Config{
		Db: infrastructure.Config{
			Port:                  8529,
			Password:              "dummy_passowrd",
			Database:              "dummy_test_app",
			Host:                  "http://localhost",
			User:                  "root",
			UserCollectionName:    "user",
			ProductCollectionName: "product",
		},
		Server: server.Config{
			Port: 5001,
			Host: "localhost",
		},
		LogLevel: 6,
	}
}
