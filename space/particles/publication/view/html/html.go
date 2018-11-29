package html

import (
	"fmt"
	"strings"

	"github.com/egoholic/tribune/space/particles/publication/entity"
)

type Publication struct {
	entity *entity.Publication
}

func Build(p *entity.Publication) *Publication {
	return &Publication{p}
}

// DO NOT USE entity's methods directly!
func (p *Publication) Title() string {
	return p.entity.Title()
}

func (p *Publication) Content() string {
	return p.entity.Content()
}

func (p *Publication) Preview() string {
	return p.entity.Preview()
}

func (p *Publication) AddFullPresentationTo(sb *strings.Builder) error {
	content := fmt.Sprintf(`<article class="article-full"><h1 class="article-full__title">%s<h1><div class="article-full__content">%s</div></article>`, p.Title(), p.Content())
	_, err := sb.WriteString(content)

	return err
}

func (p *Publication) AddPreviewPresentationTo(sb *strings.Builder) error {
	content := fmt.Sprintf(`<article class="article-preview"><h1 class="article-preview__title">%s<h1><div class="article-preview__preview">%s</div></article>`, p.Title(), p.Preview())
	_, err := sb.WriteString(content)

	return err
}

func (p *Publication) AddLinkPresentationTo(sb *strings.Builder) error {
	content := fmt.Sprintf(`<article class="article-link"><a><span class="article-link__title">%s</span></a>`, p.Title())
	_, err := sb.WriteString(content)

	return err
}
