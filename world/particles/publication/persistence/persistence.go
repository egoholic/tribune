package persistence

type Publication struct {
	Slug        string `bson:"slug"`
	Title       string `bson:"title"`
	Preview     string `bson:"preview"`
	Content     string `bson:"content"`
	RawContent  string `bson:"rawContent"`
	PublishedAt string `bson:"publishedAt"`
	CreatedAt   string `bson:"createdAt"`
}

func (p *Publication) Save() {

}

func Build(slug, title, preview, content, rawContent, publishedAt, createdAt string) *Publication {
	return &Publication{slug, title, preview, content, rawContent, publishedAt, createdAt}
}
