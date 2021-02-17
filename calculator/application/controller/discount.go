package controller

type Discount struct {
	Percentage   float32
	ValueInCents int
}

func NewProductDiscount(percentage float32, value_in_cents int) *Discount {
	return &Discount{
		Percentage:   percentage,
		ValueInCents: value_in_cents,
	}
}
