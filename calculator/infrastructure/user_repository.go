package infrastructure

import (
	"github.com/arangodb/go-driver"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
)

var GetUserDocumentError = errors.New("impossible read a user")

type UserArangoDB struct {
	collection driver.Collection
}

func NewUserRepository(coll driver.Collection) *UserArangoDB {
	return &UserArangoDB{collection: coll}
}

func (p *UserArangoDB) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	_, err := p.collection.ReadDocument(nil, id, &user)
	if err != nil {
		log.WithError(err).Warning(GetUserDocumentError)
		return nil, GetUserDocumentError
	}
	return &user, nil
}
