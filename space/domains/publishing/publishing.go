package publishing

import (
	"github.com/egoholic/tribune/space/particle/content"
	"github.com/egoholic/tribune/space/particle/publication"
	pq "github.com/egoholic/tribune/space/particle/publication/query"
)

func PublishNow(c *content.Content) publication.Publication {
	p := publication.Make(c)
	pq.Persist(&p)
	return p
}

// func PublishLater(c ce.Content, pd time.Time) {

// }
