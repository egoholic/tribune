package repository

// repository package implements through interfaces use-cases of repository pattern.
import (
	"github.com/egoholic/tribune/framework/entity"
)

type QueryForOneOwner interface {
	One(string) entity.Entity
}

type QueryForManyOwner interface {
	Many(string) []entity.Entity
}

type QueryForCountOwner interface {
	Count(string) int
}

type QueryForExistenceOwner interface {
	IsExsisting(string) bool
}

type Persister interface {
	Persist(p *entity.Entity) error
}
