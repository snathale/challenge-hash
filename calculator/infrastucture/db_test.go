package infrastucture

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/test_helper"
)

func TestDb(t *testing.T) {
	g := NewGomegaWithT(t)
	config := mockDBConfig()
	defer test_helper.DeleteDatabase(g, config.Database)
	repository, err := NewRepositories(config)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(repository.db).ShouldNot(BeNil())
	g.Expect(repository.ProductRepository).ShouldNot(BeNil())
	g.Expect(repository.UserRepository).ShouldNot(BeNil())
	client := test_helper.GetArangoClient(g)
	g.Expect(client.DatabaseExists(nil, config.Database)).Should(BeTrue())
	db, err := client.Database(nil, config.Database)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(db.CollectionExists(nil, config.UserCollectionName)).Should(BeTrue())
	g.Expect(db.CollectionExists(nil, config.ProductCollectionName)).Should(BeTrue())
}

func mockDBConfig() Config {
	return Config{
		Port:                  8529,
		Password:              "dummy_passowrd",
		Database:              "dummy_test",
		Host:                  "http://localhost",
		User:                  "root",
		UserCollectionName:    "user",
		ProductCollectionName: "product",
	}
}
