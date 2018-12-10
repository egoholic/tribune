package html_test

import (
	"strings"

	"github.com/egoholic/tribune/space/particles/publication/entity"
	. "github.com/egoholic/tribune/space/particles/publication/view/html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var publication = entity.Build("slug", "title", "preview", "content", "rawContent", "publishedAt", "createdAt")

var _ = Describe("Particles.Publication.View.Html", func() {
	Describe("Build()", func() {
		It("return new publication view", func() {
			pv := Build(publication)
			Expect(pv.Title()).To(Equal("title"))
		})
	})
	Describe(".Preview()", func() {
		It("returns preview", func() {
			pv := Build(publication)
			Expect(pv.Preview()).To(Equal("preview"))
		})
	})
	Describe(".Title()", func() {
		It("returns title", func() {
			pv := Build(publication)
			Expect(pv.Title()).To(Equal("title"))
		})
	})
	Describe(".Content()", func() {
		It("returns content", func() {
			pv := Build(publication)
			Expect(pv.Content()).To(Equal("content"))
		})
	})
	Describe(".AddFullPresentationTo()", func() {
		It("adds new content", func() {
			pv := Build(publication)
			sb := &strings.Builder{}
			pv.AddFullPresentationTo(sb)
			Expect(sb.String()).To(Equal(`<article class="article-full"><h1 class="article-full__title">title<h1><div class="article-full__content">content</div></article>`))
		})
	})
	Describe(".AddPreviewPresentationTo()", func() {
		It("adds new content", func() {
			pv := Build(publication)
			sb := &strings.Builder{}
			pv.AddPreviewPresentationTo(sb)
			Expect(sb.String()).To(Equal(`<article class="article-preview"><h1 class="article-preview__title">title<h1><div class="article-preview__preview">preview</div></article>`))
		})
	})
	Describe(".AddLinkPresentationTo()", func() {
		It("adds new content", func() {
			pv := Build(publication)
			sb := &strings.Builder{}
			pv.AddLinkPresentationTo(sb)
			Expect(sb.String()).To(Equal(`<article class="article-link"><a><span class="article-link__title">title</span></a>`))
		})
	})
})
