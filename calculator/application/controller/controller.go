package controller

import (
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/infrastucture"
)

var (
	CalculateDiscountControllerUserError    = errors.New("impossible restore user")
	CalculateDiscountControllerProductError = errors.New("impossible restore product")
)

type Controller struct {
	repository *infrastucture.Repository
}

func NewController(repository *infrastucture.Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) CalculateDiscount(userId, producId string) (*Discount, error) {
	user, err := c.repository.UserRepository.GetUserById(userId)
	if err != nil {
		logrus.WithError(err).Warning(CalculateDiscountControllerUserError)
		return nil, err
	}
	product, err := c.repository.ProductRepository.GetProductById(producId)
	if err != nil {
		logrus.WithError(err).Warning(CalculateDiscountControllerProductError)
		return nil, err
	}
	percentage := product.GetDiscount(*user)
	if percentage != 0 {
		value := (product.PriceInCents) - int(float32(product.PriceInCents)*percentage)
		return NewProductDiscount(*product, percentage, value), nil
	}
	return NewProductDiscount(*product, percentage, product.PriceInCents), nil
}
