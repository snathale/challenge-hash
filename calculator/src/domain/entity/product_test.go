package entity

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/src/utc_time"
)

func TestProduct(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("validate create a product", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)

		g.Expect(product.Id).To(Equal(""))
		g.Expect(product.Title).To(Equal("iphone"))
		g.Expect(product.Description).To(Equal("smartphone"))
		g.Expect(product.PriceInCents).To(Equal(3200))
	})
	t.Run("validate receive true when is blackFriday", func(t *testing.T) {
		replaceDate(time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC))
		g.Expect(isBlackFriday()).Should(BeTrue())
	})
	t.Run("validate receive false when is not blackFriday", func(t *testing.T) {
		replaceDate(time.Date(2021, 12, 25, 0, 0, 0, 0, time.UTC))
		g.Expect(isBlackFriday()).Should(BeFalse())
	})
	t.Run("validate receive discount equal 0.05 when is user birthday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		birthday := time.Date(2021, 01, 01, 01, 01, 0, 0, time.UTC)
		replaceDate(birthday)
		user := NewUser("User", "Test", birthday)

		g.Expect(product.getDiscount(user)).To(Equal(float32(0.05)))
	})
	t.Run("validate receive discount equal 0.1 when is blackFriday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		replaceDate(time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC))
		user := NewUser("User", "Test", time.Now())

		g.Expect(product.getDiscount(user)).To(Equal(float32(0.1)))
	})
	t.Run("validate not receive discount when not is user birthday and blackFriday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		date := time.Now()
		user := NewUser("User", "Test", date)
		replaceDate(time.Now().AddDate(0, 0, -1))
		g.Expect(product.getDiscount(user)).To(Equal(float32(0)))
	})
}

func replaceDate(date time.Time) {
	utc_time.Reset(func() time.Time {
		return date
	})
}
