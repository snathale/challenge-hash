package entity

import (
	"github.com/snathale/challenge-hash/calculator/src/utc_time"
)

type Product struct {
	Id           string `json:"_key,omitempty"`
	PriceInCents int    `json:"price_of_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

func NewProduct(title, description string, price int) Product {
	return Product{
		Id:           "",
		PriceInCents: price,
		Title:        title,
		Description:  description,
	}
}

func isBlackFriday() bool {
	today := utc_time.Now()
	return today.Day() == 25 && today.Month() == 11
}

func (p *Product) getDiscount(user User) float32 {
	if isBlackFriday() {
		return 0.1
	}
	if user.isBirthday() {
		return 0.05
	}
	return 0
}
