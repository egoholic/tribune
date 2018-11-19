package persistence

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Publication struct {
	Title       string `bson: "title"`
	Preview     string `bson: "preview"`
	Content     string `bson: "content"`
	RawContent  string `bson: "rawContent"`
	PublishedAt string `bson: "publishedAt"`
	CreatedAt   string `bson: "createdAt"`
}

func (p *Publication Save() {

}
