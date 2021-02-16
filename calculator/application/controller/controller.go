package controller

import "github.com/snathale/challenge-hash/calculator/infrastructure"

type controller struct {
	repository *infrastructure.Repository
}

func NewController(repository *infrastructure.Repository) Controller {
	return &controller{
		repository: repository,
	}
}

func (c *controller) CalculateDiscount(userId, producId string) (*Discount, error) {
	user, err := c.repository.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	product, err := c.repository.ProductRepository.GetProductById(producId)
	if err != nil {
		return nil, err
	}
	percentage := product.GetDiscount(*user)
	if percentage != 0 {
		value := (product.PriceInCents) - int(float32(product.PriceInCents)*percentage)
		return NewProductDiscount(*product, percentage, value), nil
	}
	return NewProductDiscount(*product, percentage, product.PriceInCents), nil
}
