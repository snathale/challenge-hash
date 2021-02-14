package entity

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/test_helper"
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
		test_helper.ReplaceDate(time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC))
		g.Expect(isBlackFriday()).Should(BeTrue())
	})
	t.Run("validate receive false when is not blackFriday", func(t *testing.T) {
		test_helper.ReplaceDate(time.Date(2021, 12, 25, 0, 0, 0, 0, time.UTC))
		g.Expect(isBlackFriday()).Should(BeFalse())
	})
	t.Run("validate receive discount equal 0.05 when is user birthday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		birthday := time.Date(2021, 01, 01, 01, 01, 0, 0, time.UTC)
		test_helper.ReplaceDate(birthday)
		user := NewUser("User", "Test", birthday)

		g.Expect(product.GetDiscount(*user)).To(Equal(float32(0.05)))
	})
	t.Run("validate receive discount equal 0.1 when is blackFriday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		test_helper.ReplaceDate(time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC))
		user := NewUser("User", "Test", time.Now())

		g.Expect(product.GetDiscount(*user)).To(Equal(float32(0.1)))
	})
	t.Run("validate not receive discount when not is user birthday and blackFriday", func(t *testing.T) {
		product := NewProduct("iphone", "smartphone", 3200)
		date := time.Now()
		test_helper.ReplaceDate(time.Now().AddDate(0, 0, 2))
		user := NewUser("User", "Test", date)

		g.Expect(product.GetDiscount(*user)).To(Equal(float32(0)))
	})
	t.Run("validate receive blackFriday discount even today is user bithday", func(t *testing.T) {
		date := time.Date(2021, 11, 25, 0, 0, 0, 0, time.UTC)
		test_helper.ReplaceDate(date)
		product := NewProduct("iphone", "smartphone", 3200)
		user := NewUser("User", "Test", date)

		g.Expect(product.GetDiscount(*user)).To(Equal(float32(0.1)))
	})
}
