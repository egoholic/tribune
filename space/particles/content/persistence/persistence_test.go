package persistence_test

import (
	. "github.com/egoholic/tribune/space/particles/publication/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Particles.Publication.Persistence", func() {
	slug := "slug"
	title := "title"
	preview := "preview"
	content := "content"
	rawContent := "rawContent"
	publishedAt := "publishedAt"
	createdAt := "createdAt"

	Describe("Build()", func() {
		It("creates a new Publication", func() {
			pp := Build(&slug, &title, &preview, &content, &rawContent, &publishedAt, &createdAt)
			Expect(pp.Slug()).To(Equal(slug))
			Expect(pp.Title()).To(Equal(title))
			Expect(pp.Preview()).To(Equal(preview))
			Expect(pp.Content()).To(Equal(content))
			Expect(pp.RawContent()).To(Equal(rawContent))
			Expect(pp.PublishedAt()).To(Equal(publishedAt))
			Expect(pp.CreatedAt()).To(Equal(createdAt))
		})
	})
})
