package queries_test

import (
	"github.com/egoholic/tribune/world/particles/publication/persistence"
	. "github.com/egoholic/tribune/world/particles/publication/queries"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var dataStoreMock = []*persistence.Publication{
	persistence.Build("slug-1", "Title-1", "Preview-1", "Content-1", "RawContent-1", "published-at-1", "created-at-1"),
	persistence.Build("slug-2", "Title-2", "Preview-2", "Content-2", "RawContent-2", "published-at-2", "created-at-2"),
	persistence.Build("slug-3", "Title-3", "Preview-3", "Content-3", "RawContent-3", "published-at-3", "created-at-3"),
	persistence.Build("slug-4", "Title-4", "Preview-4", "Content-4", "RawContent-4", "published-at-4", "created-at-4"),
	persistence.Build("slug-5", "Title-5", "Preview-5", "Content-5", "RawContent-5", "published-at-5", "created-at-5"),
	persistence.Build("slug-6", "Title-6", "Preview-6", "Content-6", "RawContent-6", "published-at-6", "created-at-6"),
	persistence.Build("slug-7", "Title-7", "Preview-7", "Content-7", "RawContent-7", "published-at-7", "created-at-7"),
	persistence.Build("slug-8", "Title-8", "Preview-8", "Content-8", "RawContent-8", "published-at-8", "created-at-8"),
	persistence.Build("slug-9", "Title-9", "Preview-9", "Content-9", "RawContent-9", "published-at-9", "created-at-9"),
	persistence.Build("slug-10", "Title-10", "Preview-10", "Content-10", "RawContent-10", "published-at-10", "created-at-10"),
}

var dataStore = &DataStore{dataStoreMock}

var _ = Describe("Particles.Publication.Queries", func() {
	Describe(".Latest10()", func() {
		It("Returns latest 10 publications", func() {
			result := dataStore.Latest10()
			expected := [10]*persistence.Publication{
				persistence.Build("slug-10", "Title-10", "Preview-10", "Content-10", "RawContent-10", "published-at-10", "created-at-10"),
				persistence.Build("slug-9", "Title-9", "Preview-9", "Content-9", "RawContent-9", "published-at-9", "created-at-9"),
				persistence.Build("slug-8", "Title-8", "Preview-8", "Content-8", "RawContent-8", "published-at-8", "created-at-8"),
				persistence.Build("slug-7", "Title-7", "Preview-7", "Content-7", "RawContent-7", "published-at-7", "created-at-7"),
				persistence.Build("slug-6", "Title-6", "Preview-6", "Content-6", "RawContent-6", "published-at-6", "created-at-6"),
				persistence.Build("slug-5", "Title-5", "Preview-5", "Content-5", "RawContent-5", "published-at-5", "created-at-5"),
				persistence.Build("slug-4", "Title-4", "Preview-4", "Content-4", "RawContent-4", "published-at-4", "created-at-4"),
				persistence.Build("slug-3", "Title-3", "Preview-3", "Content-3", "RawContent-3", "published-at-3", "created-at-3"),
				persistence.Build("slug-2", "Title-2", "Preview-2", "Content-2", "RawContent-2", "published-at-2", "created-at-2"),
				persistence.Build("slug-1", "Title-1", "Preview-1", "Content-1", "RawContent-1", "published-at-1", "created-at-1"),
			}
			Expect(result).To(Equal(expected))
		})
	})

	Describe("BySlug()", func() {
		Context("when publication with given slug exists", func() {
			It("returns publication with given slug", func() {
				result, err := dataStore.BySlug("slug-7")
				Expect(err).To(BeNil())
				Expect(result.Slug).To(Equal("slug-7"))
			})
		})
	})
})
