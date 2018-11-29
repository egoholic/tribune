package entity_test

import (
	. "github.com/egoholic/tribune/space/particles/publication/entity"
	"github.com/egoholic/tribune/space/particles/publication/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	slug        = "slug"
	title       = "title"
	preview     = "preview"
	content     = "content"
	rawContent  = "rawContent"
	publishedAt = "publishedAt"
	createdAt   = "createdAt"
)

var _ = Describe("Particles.Publication.Entity", func() {
	Describe("Build()", func() {
		It("returns a new Publication", func() {
			p := Build(&slug, &title, &preview, &content, &rawContent, &publishedAt, &createdAt)
			Expect(p.Slug()).To(Equal("slug"))
			Expect(p.Title()).To(Equal("title"))
			Expect(p.Preview()).To(Equal("preview"))
			Expect(p.Content()).To(Equal("content"))
			Expect(p.RawContent()).To(Equal("rawContent"))
			Expect(p.PublishedAt()).To(Equal("publishedAt"))
			Expect(p.CreatedAt()).To(Equal("createdAt"))
		})
	})

	Describe("Wrap()", func() {
		It("returns a new Publication", func() {
			pp := persistence.Build(&slug, &title, &preview, &content, &rawContent, &publishedAt, &createdAt)
			p := Wrap(pp)
			Expect(p.Slug()).To(Equal(pp.Slug()))
			Expect(p.Title()).To(Equal(pp.Title()))
			Expect(p.Preview()).To(Equal(pp.Preview()))
			Expect(p.Content()).To(Equal(pp.Content()))
			Expect(p.RawContent()).To(Equal(pp.RawContent()))
			Expect(p.PublishedAt()).To(Equal(pp.PublishedAt()))
			Expect(p.CreatedAt()).To(Equal(pp.CreatedAt()))
		})
	})
})
