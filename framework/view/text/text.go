package text

import "strings"

type Concatenator interface {
	WriteStringTo(*strings.Builder)
}

type Wrapper interface {
	Wrap(strings.Builder)
}
