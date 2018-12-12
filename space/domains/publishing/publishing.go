package publishing

import (
	"time"

	ce "github.com/egoholic/tribune/space/particles/content/entity"

	pe "github.com/egoholic/tribune/space/particles/publication/entity"
	pr "github.com/egoholic/tribune/space/particles/publication/repository"
)

func PublishNow(c ce.Content) pe.Publication {
	publication := pe.Make(&c)
	pr.Persist(&publication)
	return publication
}

func PublishLater(c ce.Content, pd time.Time)
