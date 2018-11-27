package persistence

import "gopkg.in/mgo.v2/bson"

type Publication struct {
	id          bson.ObjectId `bson:"_id"`
	slug        *string       `bson:"slug"`
	title       *string       `bson:"title"`
	preview     *string       `bson:"preview"`
	content     *string       `bson:"content"`
	rawContent  *string       `bson:"rawContent"`
	publishedAt *string       `bson:"publishedAt"`
	createdAt   *string       `bson:"createdAt"`
}

func Init(id bson.ObjectId, slug, title, preview, content, rawContent, publishedAt, createdAt *string) *Publication {
	return &Publication{id, slug, title, preview, content, rawContent, publishedAt, createdAt}
}

func Build(id bson.ObjectId, slug, title, preview, content, rawContent, publishedAt, createdAt string) *Publication {
	return &Publication{id, &slug, &title, &preview, &content, &rawContent, &publishedAt, &createdAt}
}

func (p *Publication) isPersisted() bool {
	return p.ID != nil
}

func (p *Publication) ID() string {
	return string(p.id)
}

func (p *Publication) Slug() string {
	return *p.slug
}

func (p *Publication) Title() string {
	return *p.title
}

func (p *Publication) Preview() string {
	return *p.preview
}

func (p *Publication) Content() string {
	return *p.content
}

func (p *Publication) RawContent() string {
	return *p.rawContent
}

func (p *Publication) PublishedAt() string {
	return *p.publishedAt
}

func (p *Publication) CreatedAt() string {
	return *p.createdAt
}
