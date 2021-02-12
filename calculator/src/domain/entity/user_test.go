package entity

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/src/utc_time"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)
	date := time.Date(2021, 01, 01, 01, 01, 0, 0, time.UTC)
	nowFunc := func() time.Time {
		return date
	}
	utc_time.Reset(nowFunc)
	t.Run("validate can create user", func(t *testing.T) {
		birthday := time.Now()
		user := NewUser("User", "Test", birthday)

		g.Expect(user.Id).To(Equal(""))
		g.Expect(user.FirstName).To(Equal("User"))
		g.Expect(user.LastName).To(Equal("Test"))
		g.Expect(user.Birthday).To(Equal(birthday))
	})
	t.Run("validate receve true when today is user bithday", func(t *testing.T) {
		user := NewUser("User", "Test", date)

		g.Expect(user.isBirthday()).To(BeTrue())
	})
	t.Run("validate receve false when today is not user bithday", func(t *testing.T) {
		user := NewUser("User", "Test", date.AddDate(0, 1, 0))

		g.Expect(user.isBirthday()).To(BeFalse())
	})
}
