package repository

import "github.com/snathale/challenge-hash/calculator/src/domain/entity"

type ProductRespository interface {
	GetProductById(id string) (*entity.Product, error)
}
