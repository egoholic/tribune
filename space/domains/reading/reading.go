package reading

import (
	"github.com/egoholic/tribune/space/particle/publication"
	"github.com/egoholic/tribune/space/particle/publication/query"
)

func ReadLatestOne() publication.Publication {
	return query.LatestOne()
}

func ReadArchive() []publication.Publication {
	return query.AllStartingFromLatest()
}
