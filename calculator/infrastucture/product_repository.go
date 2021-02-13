package infrastucture

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/arangodb/go-driver"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
)

var GetProductDocumentError = errors.New("impossible read a product")

type ProductArangoDB struct {
	collection driver.Collection
}

func NewProductRepository(coll driver.Collection) *ProductArangoDB {
	return &ProductArangoDB{collection: coll}
}

func (p *ProductArangoDB) GetProductById(id string) (*entity.Product, error) {
	var product entity.Product
	_, err := p.collection.ReadDocument(nil, id, &product)
	if err != nil {
		log.WithError(err).Warning(GetProductDocumentError)
		return nil, GetProductDocumentError
	}
	return &product, nil
}
