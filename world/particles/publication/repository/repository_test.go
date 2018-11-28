package repository_test

import (
	"github.com/egoholic/tribune/world/particles/publication/persistence"
	. "github.com/egoholic/tribune/world/particles/publication/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func strPtr(s string) *string {
	return &s
}

var dataStoreMock = []*persistence.Publication{
	persistence.Build(strPtr("slug-1"), strPtr("Title-1"), strPtr("Preview-1"), strPtr("Content-1"), strPtr("RawContent-1"), strPtr("published-at-1"), strPtr("created-at-1")),
	persistence.Build(strPtr("slug-2"), strPtr("Title-2"), strPtr("Preview-2"), strPtr("Content-2"), strPtr("RawContent-2"), strPtr("published-at-2"), strPtr("created-at-2")),
	persistence.Build(strPtr("slug-3"), strPtr("Title-3"), strPtr("Preview-3"), strPtr("Content-3"), strPtr("RawContent-3"), strPtr("published-at-3"), strPtr("created-at-3")),
	persistence.Build(strPtr("slug-4"), strPtr("Title-4"), strPtr("Preview-4"), strPtr("Content-4"), strPtr("RawContent-4"), strPtr("published-at-4"), strPtr("created-at-4")),
	persistence.Build(strPtr("slug-5"), strPtr("Title-5"), strPtr("Preview-5"), strPtr("Content-5"), strPtr("RawContent-5"), strPtr("published-at-5"), strPtr("created-at-5")),
	persistence.Build(strPtr("slug-6"), strPtr("Title-6"), strPtr("Preview-6"), strPtr("Content-6"), strPtr("RawContent-6"), strPtr("published-at-6"), strPtr("created-at-6")),
	persistence.Build(strPtr("slug-7"), strPtr("Title-7"), strPtr("Preview-7"), strPtr("Content-7"), strPtr("RawContent-7"), strPtr("published-at-7"), strPtr("created-at-7")),
	persistence.Build(strPtr("slug-8"), strPtr("Title-8"), strPtr("Preview-8"), strPtr("Content-8"), strPtr("RawContent-8"), strPtr("published-at-8"), strPtr("created-at-8")),
	persistence.Build(strPtr("slug-9"), strPtr("Title-9"), strPtr("Preview-9"), strPtr("Content-9"), strPtr("RawContent-9"), strPtr("published-at-9"), strPtr("created-at-9")),
	persistence.Build(strPtr("slug-10"), strPtr("Title-10"), strPtr("Preview-10"), strPtr("Content-10"), strPtr("RawContent-10"), strPtr("published-at-10"), strPtr("created-at-10")),
}

var conn = CollectionFrom(MongoDBSession())
var repo = Make(conn)

var _ = Describe("Particles.Publication.Queries", func() {
	Describe(".Latest10()", func() {
		It("Returns latest 10 publications", func() {
			result := repo.LatestPublished10()
			expected := [10]*persistence.Publication{
				persistence.Build(strPtr("slug-10"), strPtr("Title-10"), strPtr("Preview-10"), strPtr("Content-10"), strPtr("RawContent-10"), strPtr("published-at-10"), strPtr("created-at-10")),
				persistence.Build(strPtr("slug-9"), strPtr("Title-9"), strPtr("Preview-9"), strPtr("Content-9"), strPtr("RawContent-9"), strPtr("published-at-9"), strPtr("created-at-9")),
				persistence.Build(strPtr("slug-8"), strPtr("Title-8"), strPtr("Preview-8"), strPtr("Content-8"), strPtr("RawContent-8"), strPtr("published-at-8"), strPtr("created-at-8")),
				persistence.Build(strPtr("slug-7"), strPtr("Title-7"), strPtr("Preview-7"), strPtr("Content-7"), strPtr("RawContent-7"), strPtr("published-at-7"), strPtr("created-at-7")),
				persistence.Build(strPtr("slug-6"), strPtr("Title-6"), strPtr("Preview-6"), strPtr("Content-6"), strPtr("RawContent-6"), strPtr("published-at-6"), strPtr("created-at-6")),
				persistence.Build(strPtr("slug-5"), strPtr("Title-5"), strPtr("Preview-5"), strPtr("Content-5"), strPtr("RawContent-5"), strPtr("published-at-5"), strPtr("created-at-5")),
				persistence.Build(strPtr("slug-4"), strPtr("Title-4"), strPtr("Preview-4"), strPtr("Content-4"), strPtr("RawContent-4"), strPtr("published-at-4"), strPtr("created-at-4")),
				persistence.Build(strPtr("slug-3"), strPtr("Title-3"), strPtr("Preview-3"), strPtr("Content-3"), strPtr("RawContent-3"), strPtr("published-at-3"), strPtr("created-at-3")),
				persistence.Build(strPtr("slug-2"), strPtr("Title-2"), strPtr("Preview-2"), strPtr("Content-2"), strPtr("RawContent-2"), strPtr("published-at-2"), strPtr("created-at-2")),
				persistence.Build(strPtr("slug-1"), strPtr("Title-1"), strPtr("Preview-1"), strPtr("Content-1"), strPtr("RawContent-1"), strPtr("published-at-1"), strPtr("created-at-1")),
			}
			Expect(result).To(Equal(expected))
		})
	})

	Describe("BySlug()", func() {
		Context("when publication with given slug exists", func() {
			It("returns publication with given slug", func() {
				result, err := repo.BySlug("slug-7")
				Expect(err).To(BeNil())
				Expect(result.Slug).To(Equal("slug-7"))
			})
		})
	})
})
