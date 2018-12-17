package query

import (
	"github.com/egoholic/tribune/space/particle/publication"
)

func LatestOne() publication.Publication {
	latestId = len(data) - 1
	return data[latestId]
}

func AllStartingFromLatest() []publication.Publication {
	return data
}

func Persist(p *publication.Publication) bool {
	data = append(data, *p)
	return true
}
