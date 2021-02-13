package repository

import "github.com/snathale/challenge-hash/calculator/domain/entity"

type UserRepository interface {
	GetUserById(id string) (*entity.User, error)
}
