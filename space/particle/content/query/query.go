package query

import (
	"github.com/egoholic/tribune/framework/entity"
)

type Repository struct {
	queries map[string]interface{}
}

// func One(qn string) entity.Entity {

// }

// func Many(qn string) entity.Entity {

// }

func Persist(c entity.Entity) bool {
	return true
}
