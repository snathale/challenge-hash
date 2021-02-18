package test_helper

import (
	"fmt"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/utc_time"
)

const DBHOST = "http://localhost"
const DBPORT = 8529
const DBUSER = "root"
const DBPASSWORD = "dummy_passowrd"

func DeleteDatabase(g *GomegaWithT, name string) {
	c := GetArangoClient(g)
	db, err := c.Database(nil, name)
	g.Expect(err).ShouldNot(HaveOccurred())
	err = db.Remove(nil)
	g.Expect(err).ShouldNot(HaveOccurred())
}

func GetArangoClient(g *GomegaWithT) driver.Client {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{fmt.Sprintf("%s:%d", DBHOST, DBPORT)},
	})
	g.Expect(err).ShouldNot(HaveOccurred())
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUSER, DBPASSWORD),
	})
	g.Expect(err).ShouldNot(HaveOccurred())
	return c
}

func GetCollection(g *GomegaWithT, name string, db driver.Database) driver.Collection {
	exist, err := db.CollectionExists(nil, name)
	g.Expect(err).ShouldNot(HaveOccurred())
	if exist {
		coll, err := db.Collection(nil, name)
		g.Expect(err).ShouldNot(HaveOccurred())
		return coll
	}
	coll, err := db.CreateCollection(nil, name, nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	return coll
}

func CreateDatabase(g *GomegaWithT, name string) driver.Database {
	c := GetArangoClient(g)
	exist, err := c.DatabaseExists(nil, name)
	g.Expect(err).ShouldNot(HaveOccurred())
	if exist {
		db, err := c.Database(nil, name)
		g.Expect(err).ShouldNot(HaveOccurred())
		return db
	}
	db, err := c.CreateDatabase(nil, name, nil)
	g.Expect(err).ShouldNot(HaveOccurred())
	return db
}

func GetDatabase(g *GomegaWithT, name string) driver.Database {
	c := GetArangoClient(g)
	db, err := c.Database(nil, name)
	g.Expect(err).ShouldNot(HaveOccurred())
	return db
}

func CreateDocument(g *GomegaWithT, coll driver.Collection, doc interface{}) driver.DocumentMeta {
	meta, err := coll.CreateDocument(nil, doc)
	g.Expect(err).ShouldNot(HaveOccurred())
	return meta
}

func ReplaceDate(date time.Time) {
	utc_time.Reset(func() time.Time {
		return date
	})
}
