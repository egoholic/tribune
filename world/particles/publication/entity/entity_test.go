package entity_test

import (
	. "github.com/egoholic/tribune/world/particles/publication/entity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Particles.Publication.Entity", func() {
	Describe("Build()", func() {
		It("creates a new Publication", func() {
			p := Build("slug", "title", "preview", "content", "rawContent", "publishedAt", "createdAt")
			Expect(p.Slug()).To(Equal("slug"))
			Expect(p.Title()).To(Equal("title"))
			Expect(p.Preview()).To(Equal("preview"))
			Expect(p.Content()).To(Equal("content"))
			Expect(p.RawContent()).To(Equal("rawContent"))
			Expect(p.PublishedAt()).To(Equal("publishedAt"))
			Expect(p.CreatedAt()).To(Equal("createdAt"))
		})
	})
})
