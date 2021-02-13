package infrastucture

import (
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/domain/repository"
)

var (
	ArangoDBConnectionError       = errors.New("impossible connect to arangoDb")
	ArangoDBClientConnectionError = errors.New("impossible client connect to arangoDb")
	ArangoDBDatabaseExistError    = errors.New("impossible check database exists")
	ArangoDBDatabaseCreateError   = errors.New("impossible create database")
	ArangoDBDatabaseGetError      = errors.New("impossible retreive database")
	ArangoDBCollectionExistError  = errors.New("impossible check collection exist")
	ArangoDBCollectionCreateError = errors.New("impossible create collection")
	ArangoDBCollectionGetError    = errors.New("impossible retrieve collection")
)

type Repository struct {
	userRepository    repository.UserRepository
	productRepository repository.ProductRespository
	db                driver.Database
}

func NewRepositories(config Config) (*Repository, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{fmt.Sprintf("%s:%d", config.Host, config.Port)},
	})
	if err != nil {
		log.WithError(err).Warning(ArangoDBConnectionError)
		return nil, ArangoDBConnectionError
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.User, config.Password),
	})
	if err != nil {
		log.WithError(err).Warning(ArangoDBClientConnectionError)
		return nil, ArangoDBClientConnectionError
	}
	var db driver.Database
	if db, err = getDatabase(client, config.Database); err != nil {
		return nil, err
	}
	var productColl driver.Collection
	if productColl, err = getCollection(db, config.ProductCollectionName); err != nil {
		return nil, err
	}
	var userColl driver.Collection
	if userColl, err = getCollection(db, config.UserCollectionName); err != nil {
		return nil, err
	}
	return &Repository{
		UserRepository:    NewUserRepository(userColl),
		ProductRepository: NewProductRepository(productColl),
		db:                db,
	}, nil
}

func getDatabase(c driver.Client, name string) (driver.Database, error) {
	dbExist, err := c.DatabaseExists(nil, name)
	if err != nil {
		log.WithError(err).Warning(ArangoDBDatabaseExistError)
		return nil, ArangoDBDatabaseExistError
	}
	var db driver.Database
	if !dbExist {
		if db, err = c.CreateDatabase(nil, name, nil); err != nil {
			log.WithError(err).Warning(ArangoDBDatabaseCreateError)
			return nil, ArangoDBDatabaseCreateError
		}
	}
	if db, err = c.Database(nil, name); err != nil {
		log.WithError(err).Warning(ArangoDBDatabaseGetError)
		return nil, ArangoDBDatabaseGetError
	}
	return db, nil
}

func getCollection(db driver.Database, name string) (driver.Collection, error) {
	collExist, err := db.CollectionExists(nil, name)
	if err != nil {
		log.WithError(err).Warning(ArangoDBCollectionExistError)
		return nil, ArangoDBCollectionExistError
	}
	var coll driver.Collection
	if !collExist {
		if coll, err = db.CreateCollection(nil, name, nil); err != nil {
			log.WithError(err).Warning(ArangoDBCollectionCreateError)
			return nil, ArangoDBCollectionCreateError
		}
		return coll, nil
	}
	if coll, err = db.Collection(nil, name); err != nil {
		log.WithError(err).Warning(ArangoDBCollectionGetError)
		return nil, ArangoDBCollectionGetError
	}
	return coll, nil
}
