package blog_test

import (
	"strings"

	. "github.com/egoholic/tribune/viewPoints/blog/view/layouts/blog"
	"github.com/egoholic/tribune/world/particles/publication/entity"
	"github.com/egoholic/tribune/world/particles/publication/view/html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var publication = entity.Build("slug", "title", "preview", "content", "rawContent", "publishedAt", "createdAt")
var view = html.Build(publication)

var titleApplicator = func(sb *strings.Builder) error {
	sb.WriteString("Some Awesome Page Title")
	return nil
}
var contentApplicator = view.AddFullPresentationTo

var _ = Describe("ViewPoints.Blog.View.Layouts.Blog", func() {
	Describe("Build()", func() {
		It("returns layout", func() {
			l := Build(&titleApplicator, &contentApplicator)
			Expect(l).NotTo(BeNil())
		})
	})
	Describe(".Render()", func() {
		It("returns string", func() {
			l := Build(&titleApplicator, &contentApplicator)
			expected := `<!DOCTYPE html><html><head><title>Some Awesome Page Title</title></head><body><body><article class="article-full"><h1 class="article-full__title">title<h1><div class="article-full__content">content</div></article></body></html>`
			Expect(l.Render()).To(Equal(expected))
		})
	})
})
