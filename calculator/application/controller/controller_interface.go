package controller

type Controller interface {
	CalculateDiscount(userId, producId string) (*Discount, error)
}
