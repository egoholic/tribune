package Persister

// repository package implements through interfaces use-cases of repository pattern.
import (
	"github.com/egoholic/tribune/framework/entity"
)

type Persister interface {
	Persist(p entity.Entity) error
}
