package html

import (
	"strings"

	"github.com/egoholic/tribune/world/particles/publication/entity"
)

type Publication struct {
	entity *entity.Publication
}

// TODO: implement
func (p *Publication) AddFullPresentationTo(sb strings.Builder) error {
	return nil
}

// TODO: implement
func (p *Publication) AddPreviewPresentationTo(sb strings.Builder) error {
	return nil
}

// TODO: implement
func (p *Publication) AddLinkPresentationTo(sb strings.Builder) error {
	return nil
}
