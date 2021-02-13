package infrastucture

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/snathale/challenge-hash/calculator/domain/entity"
	"github.com/snathale/challenge-hash/calculator/utc_time"
)

func TestUserRepository(t *testing.T) {
	g := NewGomegaWithT(t)
	defer DeleteDatabase(g, "dummy_test")
	db := CreateDatabase(g, "dummy_test")
	date := time.Date(2021, 01, 01, 01, 01, 0, 0, time.UTC)
	nowFunc := func() time.Time {
		return date
	}
	utc_time.Reset(nowFunc)
	t.Run("validate retrieve a existing user", func(t *testing.T) {
		coll := GetCollection(g, "user", db)
		birthday := utc_time.Now()
		doc := entity.User{
			FirstName: "Joe",
			LastName:  "Lee",
			Birthday:  birthday,
		}
		meta := CreateDocument(g, coll, doc)
		rep := NewUserRepository(coll)
		user, err := rep.GetUserById(meta.Key)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(user.FirstName).Should(Equal("Joe"))
		g.Expect(user.LastName).Should(Equal("Lee"))
		g.Expect(user.Birthday).Should(Equal(birthday))
		g.Expect(user.Id).ShouldNot(BeEmpty())
	})
	t.Run("validade receive error when user not exists", func(t *testing.T) {
		coll := GetCollection(g, "user", db)
		rep := NewUserRepository(coll)
		product, err := rep.GetUserById("14")
		g.Expect(product).Should(BeNil())
		g.Expect(err).Should(MatchError(GetUserDocumentError))
	})
}
