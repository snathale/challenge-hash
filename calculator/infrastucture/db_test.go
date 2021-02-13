package infrastucture

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDb(t *testing.T) {
	g := NewGomegaWithT(t)
	config := mockDBConfig()
	defer DeleteDatabase(g, config.Database)
	repository, err := NewRepositories(config)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(repository.db).ShouldNot(BeNil())
	g.Expect(repository.productRepository).ShouldNot(BeNil())
	g.Expect(repository.userRepository).ShouldNot(BeNil())
	client := GetArangoClient(g)
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
