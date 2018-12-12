package repository

import (
	"github.com/egoholic/tribune/framework/entity"
)

type Repository struct {
	queries map[string]interface{}
}

func Persist(c *entity.Entity) bool {
	return true
}
