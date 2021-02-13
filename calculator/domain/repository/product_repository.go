package repository

import "github.com/snathale/challenge-hash/calculator/domain/entity"

type ProductRespository interface {
	GetProductById(id string) (*entity.Product, error)
}
