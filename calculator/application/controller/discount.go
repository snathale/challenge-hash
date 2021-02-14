package controller

import "github.com/snathale/challenge-hash/calculator/domain/entity"

type Discount struct {
	Percentage   float32
	ValueInCents int
}

func NewProductDiscount(product entity.Product, percentage float32, value_in_cents int) *Discount {
	return &Discount{
		Percentage:   percentage,
		ValueInCents: value_in_cents,
	}
}
