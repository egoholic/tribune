package persistence

import "gopkg.in/mgo.v2/bson"

type Publication struct {
	ID         bson.ObjectId `bson:"_id`
	Slug       string        `bson:"slug"`
	Title      string        `bson:"title"`
	Preview    string        `bson:"preview"`
	Content    string        `bson:"content"`
	RawContent string        `bson:"rawContent"`
}

func New() {

}

func (pub *Publication) Assign() {

}

func (pub *Publication) Save() {

}
