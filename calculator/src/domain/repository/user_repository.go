package repository

import "github.com/snathale/challenge-hash/calculator/src/domain/entity"

type UserRepository interface {
	GetUserById(id string) (*entity.User, error)
}
