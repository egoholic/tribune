package repository

import (
	"github.com/egoholic/tribune/framework/entity"
	. "github.com/egoholic/tribune/framework/repository"
)

type Repository struct {
	queries map[string]interface{}
}

func One(qn string) entity.Entity {

}

func Many(qn string) entity.Entity {

}
